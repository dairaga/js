//go:build js && wasm

package app

import (
	"fmt"
	"reflect"

	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/url"
)

// -----------------------------------------------------------------------------

type app struct {
	window   js.Value
	history  js.Value
	models   map[string]reflect.Value
	triggers map[string][]reflect.Value
	//hander   StateHandler
}

var _app = &app{
	models:   make(map[string]reflect.Value),
	triggers: make(map[string][]reflect.Value),
}

// -----------------------------------------------------------------------------

func (a *app) init() {
	a.window = js.Window()
	a.history = a.window.Get("history")
}

// -----------------------------------------------------------------------------

func (a *app) url() url.URL {
	return url.New(a.window.Get("location").Get("href").String())
}

// -----------------------------------------------------------------------------

func (a *app) state(x any) (err error) {
	state := a.history.Get("state")
	err = js.Unmarshal(state, x)
	return
}

// -----------------------------------------------------------------------------

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

func (a *app) push(x any, newURL ...string) error {
	return a.change("pushState", x, newURL...)
}

// -----------------------------------------------------------------------------

func (a *app) replace(x any, newURL ...string) error {
	return a.change("replaceState", x, newURL...)
}

// -----------------------------------------------------------------------------

func (a *app) watch(name string, fn any) {
	val := reflect.ValueOf(fn)
	if reflect.Func != val.Kind() {
		panic(fmt.Sprintf("x must be a function, but %v", val.Kind()))
	}
	a.triggers[name] = append(a.triggers[name], val)
}

// -----------------------------------------------------------------------------

func (a *app) _go(delta int) {
	if delta != 0 {
		a.history.Call("go", delta)
	}
}

// -----------------------------------------------------------------------------

func (a *app) changeHash(new string) {
	cur := a.url()
	cur.SetHash(new)
	js.Redirect(cur.String())
}

// -----------------------------------------------------------------------------

func (a *app) _var(val any, name string, fn any) {
	v := reflect.ValueOf(val)
	if reflect.Ptr != v.Kind() {
		panic(fmt.Sprintf("x must be ptr, but %v", v.Kind()))
	}

	old, ok := a.models[name]
	if ok {
		oldv := reflect.ValueOf(old)
		if oldv != v {
			panic(fmt.Sprintf("%s existed", name))
		}
		return
	}
	a.models[name] = v
	a.watch(name, fn)
}

// -----------------------------------------------------------------------------

func (a *app) remove(name string) {
	delete(a.models, name)
	delete(a.triggers, name)
}

// -----------------------------------------------------------------------------

func (a *app) trigger(sender, name string) {
	val, ok := a.models[name]
	if !ok {
		return
	}

	callbacks := a.triggers[name]
	size := len(callbacks)
	if size > 0 {
		args := []reflect.Value{reflect.ValueOf(sender), val.Elem()}
		for i := 0; i < size; i++ {
			callbacks[i].Call(args)
		}
	}
}

// -----------------------------------------------------------------------------

func (a *app) bindElement(x any, val *string, name string, fn func(string, string)) js.Element {
	elm := js.ElementOf(x)
	a._var(val, name, fn)
	return elm.OnChange(func(e js.Element, _ js.Event) {
		e.Var(val)
		a.trigger(e.Tattoo(), name)
	})
}

// -----------------------------------------------------------------------------

func init() {
	_app.init()
}

// -----------------------------------------------------------------------------

func Start(fn ...func()) {
	for i := range fn {
		fn[i]()
	}
	select {}
}

// -----------------------------------------------------------------------------

// URL returns current url.
func URL() url.URL {
	return _app.url()
}

// -----------------------------------------------------------------------------

func ChangeHash(new string) {
	_app.changeHash(new)
}

// -----------------------------------------------------------------------------

func Push(x any, newURL ...string) error {
	return _app.push(x, newURL...)
}

// -----------------------------------------------------------------------------

func Replace(x any, newURL ...string) error {
	return _app.replace(x, newURL...)
}

// -----------------------------------------------------------------------------

func State(x any) error {
	return _app.state(x)
}

// -----------------------------------------------------------------------------

// Go is javascript window.history.go(detla).
// See https://developer.mozilla.org/en-US/docs/Web/API/History/go
func Go(delta int) {
	_app._go(delta)
}

// -----------------------------------------------------------------------------

// Forward is javascript window.history.forward.
// See https://developer.mozilla.org/en-US/docs/Web/API/History/forward
func Forward() {
	_app._go(1)
}

// -----------------------------------------------------------------------------

// Back is javascript window.history.back.
// https://developer.mozilla.org/en-US/docs/Web/API/History/back
func Back() {
	_app._go(-1)
}

// -----------------------------------------------------------------------------

func Var(val any, name string, fn any) {
	_app._var(val, name, fn)
}

// -----------------------------------------------------------------------------

func Watch(name string, fn any) {
	_app.watch(name, fn)
}

// -----------------------------------------------------------------------------

func Remove(name string) {
	_app.remove(name)
}

// -----------------------------------------------------------------------------

func Trigger(sender, name string) {
	_app.trigger(sender, name)
}

// -----------------------------------------------------------------------------

func BindElement(x any, val *string, name string, fn func(string, string)) js.Element {
	return _app.bindElement(x, val, name, fn)
}
