//go:build js && wasm

// Package is a WASM application. Invoking app.Init first with given handler responses to URL changed event.
//
// And than invoking app.Start with given handler called when application starting.
//
// Handler in app.Init will be called when application starting, if there is no handler in app.Start.
//
// Example Usage
//
//		func main() {
//			app.Init(app.HandlerFunc(func(old, new url.URL) {
//
//			}))
//
//			app.Start(app.HandlerFunc(func(old, new url.URL) {
//
//			}))
//
//		}
//
package app

import (
	"fmt"

	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/url"
)

// -----------------------------------------------------------------------------

// Handler responds to event when URL's hash changed.
type Handler interface {
	Serve(oldURL, curURL url.URL)
}

// -----------------------------------------------------------------------------

// The HandlerFunc type is an adapter to allow the use of ordinary functions as app Handler.
// If f is a function with appropriate signature, HandlerFunc(f) is a Handler that calls f.
type HandlerFunc func(url.URL, url.URL)

func (f HandlerFunc) Serve(oldURL, newURL url.URL) {
	f(oldURL, newURL)
}

// -----------------------------------------------------------------------------

// The app type represents a WASM application.
type app struct {
	window     js.Value // javascript Window.
	history    js.Value // javascript Window.history.
	currentURL url.URL  // current web page URL.
	handler    Handler  // app Handler called when url's hash changed.
	hashFunc   js.Func  // listener function to handle event that url's hash changed.
}

var _app *app

// -----------------------------------------------------------------------------

func (a *app) init() {
	cb := js.FuncOf(func(_ js.Value, args []js.Value) any {
		if a.handler != nil {
			old := url.New(args[0].Get("oldURL").String())
			new := url.New(args[0].Get("newURL").String())
			a.handler.Serve(old, new)
		}
		return nil
	})
	a.hashFunc = cb

	a.window.Call("addEventListener", "hashchange", cb)

	a.currentURL = a.url()
}

// -----------------------------------------------------------------------------

func (a *app) changeHash(new string) {
	a.currentURL.SetHash(new)
	js.Redirect(a.currentURL.String())
}

// -----------------------------------------------------------------------------

func (a *app) url() url.URL {
	return url.New(a.window.Get("location").Get("href").String())
}

// -----------------------------------------------------------------------------

func (a *app) push(newURL string, x ...any) {
	state := js.Null()
	size := len(x)
	var err error
	if size == 1 {
		state, err = js.Marshal(x[0])
	} else if size > 1 {
		state, err = js.Marshal(x)
	}
	if err != nil {
		// console.log error, do not panic to interrupt application.
		fmt.Println("warn: push state:", err)
	}

	oldURL := a.currentURL
	a.history.Call("pushState", state, "", newURL)
	a.currentURL = a.url()

	if a.handler != nil {
		a.handler.Serve(oldURL, a.currentURL)
	}
}

// -----------------------------------------------------------------------------

// state returns current state in the
func (a *app) state(x any) (err error) {
	state := a.history.Get("state")
	err = js.Unmarshal(state, x)
	return
}

// -----------------------------------------------------------------------------

func (a *app) _go(delta int) {
	oldURL := a.currentURL
	if delta != 0 {
		a.history.Call("go", delta)
		a.currentURL = a.url()
	}

	if a.handler != nil {
		a.handler.Serve(oldURL, a.currentURL)
	}
}

// -----------------------------------------------------------------------------

// Init initialize WASM application and add a handler reponses to URL changed event.
func Init(h ...Handler) {
	var handler Handler = nil

	if len(h) > 0 {
		handler = h[0]
	}

	_app = &app{
		window:     js.Window(),
		history:    js.Window().Get("history"),
		currentURL: url.New(js.Window().Get("location").Get("href").String()),
		handler:    handler,
	}
	_app.init()
}

// -----------------------------------------------------------------------------

// URL returns current url.
func URL() url.URL {
	return _app.currentURL
}

// -----------------------------------------------------------------------------

// ChangeHash changes current url hash. It will invoke Handler given to app.Init.
func ChangeHash(new string) {
	_app.changeHash(new)
}

// -----------------------------------------------------------------------------

// Push changes current url. It will invoke Handler given to app.Init.
// It makes current url changed, but browser does not load this url.
func Push(newURL string, x ...any) {
	_app.push(newURL, x...)
}

// -----------------------------------------------------------------------------

// State returns current state added by Push.
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

// Start starts a WASM application with a given Hanlder.
// The given Handler will be invoked when appication started.
// Handler given to app.Init will be invoked if no hanlder is for Start.
func Start(h ...Handler) {
	if len(h) > 0 {
		h[0].Serve(_app.currentURL, _app.currentURL)
	} else if _app.handler != nil {
		_app.handler.Serve(_app.currentURL, _app.currentURL)
	}

	select {} // block main
}
