//go:build js && wasm

// Package app is a WASM application to handle web browser history and MVVM.
package app

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/url"
)

// -----------------------------------------------------------------------------

// app represents a application to handle web browser history and MVVM.
type app struct {
	window  js.Value // Javascript window.
	history js.Value // Javascript history of window.
	//models   map[string]reflect.Value
	triggers map[string][]func(string) // Callback functions when value changed.
	//hander   StateHandler
	ch    chan int
	ended bool
}

// an internal global app.
var _app = &app{
	//models:   make(map[string]reflect.Value),
	triggers: make(map[string][]func(string)),
	ch:       make(chan int, 1),
	ended:    false,
}

// -----------------------------------------------------------------------------

// init intializes the app.
func (a *app) init() {
	a.window = js.Window()
	a.history = a.window.Get("history")
}

// -----------------------------------------------------------------------------

// url returns current url of window.
func (a *app) url() url.URL {
	return url.New(a.window.Get("location").Get("href").String())
}

// -----------------------------------------------------------------------------

// state returns current state in history and unmarshal it to given x.
func (a *app) state(x any) (err error) {
	state := a.history.Get("state")
	err = js.Unmarshal(state, x)
	return
}

// -----------------------------------------------------------------------------

// change is a helper function to handle history state. Method is one of "pushState" or "replaceState".
func (a *app) change(method string, x any, newURL ...string) error {
	state, err := js.Marshal(x)

	if err != nil {
		return err
	}

	if len(newURL) > 0 {
		a.history.Call(method, state, "", newURL[0])
	} else {
		a.history.Call(method, state, "")
	}

	return nil
}

// -----------------------------------------------------------------------------

// push is to call javascript window.history.pushState.
func (a *app) push(x any, newURL ...string) error {
	return a.change("pushState", x, newURL...)
}

// -----------------------------------------------------------------------------

// replace is to call javascript window.history.replaceState.
func (a *app) replace(x any, newURL ...string) error {
	return a.change("replaceState", x, newURL...)
}

// -----------------------------------------------------------------------------

// go is to load page from window history.
func (a *app) _go(delta int) {
	if delta != 0 {
		a.history.Call("go", delta)
	}
}

// -----------------------------------------------------------------------------

// changeHash changes hash of window url. It will trigger hashchange event.
func (a *app) changeHash(new string) {
	cur := a.url()
	cur.SetHash(new)
	js.Redirect(cur.String())
}

// -----------------------------------------------------------------------------

// watch is to add a trigger function fn for some value.
func (a *app) watch(name string, fn func(string)) {
	//val := reflect.ValueOf(fn)
	//if reflect.Func != val.Kind() {
	//	panic(fmt.Sprintf("x must be a function, but %v", val.Kind()))
	//}
	a.triggers[name] = append(a.triggers[name], fn)
}

// -----------------------------------------------------------------------------

//func (a *app) _var(name string, fn func(string)) {
//	//v := reflect.ValueOf(val)
//	//if reflect.Ptr != v.Kind() {
//	//	panic(fmt.Sprintf("x must be ptr, but %v", v.Kind()))
//	//}
//
//	//old, ok := a.models[name]
//	//if ok {
//	//	oldv := reflect.ValueOf(old)
//	//	if oldv != v {
//	//		panic(fmt.Sprintf("%s existed", name))
//	//	}
//	//	return
//	//}
//	//a.models[name] = v
//	a.watch(name, fn)
//}

// -----------------------------------------------------------------------------

// remove removes all triggers for some value. The given `name` is from bindElement or watch.
func (a *app) remove(name string) {
	//delete(a.models, name)
	delete(a.triggers, name)
}

// -----------------------------------------------------------------------------

// trigger is to fire all functions binded the value. Give a `sender` to mark who fires.
func (a *app) trigger(sender, name string) {
	//val, ok := a.models[name]
	//if !ok {
	//	return
	//}

	//callbacks := a.triggers[name]
	//size := len(callbacks)
	//if size > 0 {

	//args := []reflect.Value{reflect.ValueOf(sender), val.Elem()}
	//for i := 0; i < size; i++ {
	//	callbacks[i].Call(args)
	//}
	//}

	for _, cb := range a.triggers[name] {
		cb(sender)
	}
}

// -----------------------------------------------------------------------------

// bindElement binds a value to a element. Given value is changed when element value changed.
func (a *app) bindElement(x any, val *string, name string, fn func(string)) js.Element {
	elm := js.ElementOf(x)
	a.watch(name, fn)
	//a._var(val, name, fn)
	return elm.OnChange(func(e js.Element, _ js.Event) {
		e.Var(val)
		a.trigger(e.Tattoo(), name)
	})
}

// -----------------------------------------------------------------------------

// init package initialization.
func init() {
	_app.init()
}

// -----------------------------------------------------------------------------

// Start starts a WASM application. all given functions fn will be called sequentially before blocking main process.
func Start(fn ...func()) int {
	for i := range fn {
		if !_app.ended {
			fn[i]()
		}
	}
	exitCode := <-_app.ch
	close(_app.ch)
	return exitCode
}

// -----------------------------------------------------------------------------

func Exit(code int) {
	_app.ended = true
	_app.ch <- code
}

// -----------------------------------------------------------------------------

// URL returns current url.
func URL() url.URL {
	return _app.url()
}

// -----------------------------------------------------------------------------

// ChangeHash changes hash of window url. It will trigger hashchange event.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/hashchange_event.
func ChangeHash(new string) {
	_app.changeHash(new)
}

// -----------------------------------------------------------------------------

// PushState changes history state and push it to history.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/History/pushState.
func Push(x any, newURL ...string) error {
	return _app.push(x, newURL...)
}

// -----------------------------------------------------------------------------

// ReplaceState changes history state and replace it to history.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/History/replaceState.
func Replace(x any, newURL ...string) error {
	return _app.replace(x, newURL...)
}

// -----------------------------------------------------------------------------

// State returns current state in history and unmarshal it to given x.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/History/state.
func State(x any) error {
	return _app.state(x)
}

// -----------------------------------------------------------------------------

// Go is javascript window.history.go(detla).
//
// See https://developer.mozilla.org/en-US/docs/Web/API/History/go.
func Go(delta int) {
	_app._go(delta)
}

// -----------------------------------------------------------------------------

// Forward is javascript window.history.forward.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/History/forward
func Forward() {
	_app._go(1)
}

// -----------------------------------------------------------------------------

// Back is javascript window.history.back.
//
// https://developer.mozilla.org/en-US/docs/Web/API/History/back
func Back() {
	_app._go(-1)
}

// -----------------------------------------------------------------------------

//func Var(val any, name string, fn any) {
//	_app._var(val, name, fn)
//}

// -----------------------------------------------------------------------------

// Watch watches a value and call fn when value changed.
func Watch(name string, fn func(string)) {
	_app.watch(name, fn)
}

// -----------------------------------------------------------------------------

// Remove removes all triggers for some value. The given `name` is from BindElement or Watch.
func Remove(name string) {
	_app.remove(name)
}

// -----------------------------------------------------------------------------

// Trigger is to fire all functions binded the value. Give a `sender` to notify others who fires.
func Trigger(sender, name string) {
	_app.trigger(sender, name)
}

// -----------------------------------------------------------------------------

// BindElement binds a value to a element. Given value is changed when element value changed.
func BindElement(x any, val *string, name string, fn func(string)) js.Element {
	return _app.bindElement(x, val, name, fn)
}
