//go:build js && wasm

// Package is a WASM application. Invoking app.Init first with given handler called when URL's hash changed,
// and than app.Start with given handler called when application starting.
// Handler in app.Init will be called when application starting, if there is no handler in app.Start.
package app

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/url"
)

// -----------------------------------------------------------------------------

// Handler responds to event when URL's hash changed.
type Handler interface {
	Serve(oldURL, curURL url.URL, state any)
}

// -----------------------------------------------------------------------------

// The HandlerFunc type is an adapter to allow the use of ordinary functions as app Handler.
// If f is a function with appropriate signature, HandlerFunc(f) is a Handler that calls f.
type HandlerFunc func(url.URL, url.URL, any)

func (f HandlerFunc) Serve(oldURL, newURL url.URL, state any) {
	f(oldURL, newURL, state)
}

// -----------------------------------------------------------------------------

// The app type represents a WASM application.
type app struct {
	window     js.Value       // javascript Window.
	history    js.Value       // javascript Window.history.
	currentURL url.URL        // current web page URL.
	handler    Handler        // app Handler called when url's hash changed.
	hashFunc   js.Func        // listener function to handle event that url's hash changed.
	states     map[string]any // store history.state
}

var _app *app

// -----------------------------------------------------------------------------

func (a *app) init() {
	cb := js.FuncOf(func(_ js.Value, args []js.Value) any {
		if a.handler != nil {
			old := url.New(args[0].Get("oldURL").String())
			new := url.New(args[0].Get("newURL").String())
			state := a.state()
			a.handler.Serve(old, new, state)
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

func (a *app) state() any {
	tmp := a.history.Get("state")
	if tmp.Truthy() && tmp.Type() == js.TypeString {
		return a.states[tmp.String()]
	}
	return nil
}

// -----------------------------------------------------------------------------

func (a *app) pushState(newURL, state string, x any) {
	oldURL := a.currentURL
	if state != "" {
		a.states[state] = x
	}
	a.history.Call("pushState", state, "", newURL)
	a.currentURL = a.url()

	if a.handler != nil {
		a.handler.Serve(oldURL, a.currentURL, x)
	}
}

// -----------------------------------------------------------------------------

func (a *app) _go(delta int) {
	oldURL := a.currentURL
	if delta != 0 {
		a.history.Call("go", delta)
		a.currentURL = a.url()
	}

	state := a.state()

	if a.handler != nil {
		a.handler.Serve(oldURL, a.currentURL, state)
	}
}

// -----------------------------------------------------------------------------

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
		states:     make(map[string]any),
	}
	_app.init()
}

// -----------------------------------------------------------------------------

func URL() url.URL {
	return _app.currentURL
}

// -----------------------------------------------------------------------------

func ChangeHash(new string) {
	_app.changeHash(new)
}

// -----------------------------------------------------------------------------

func PushState(newURL, state string, data any) {
	_app.pushState(newURL, state, data)
}

// -----------------------------------------------------------------------------

func Push(newURL string) {
	_app.pushState(newURL, "", nil)
}

// -----------------------------------------------------------------------------

func State() any {
	return _app.state()
}

// -----------------------------------------------------------------------------

func Go(delta int) {
	_app._go(delta)
}

// -----------------------------------------------------------------------------

func Forward() {
	_app._go(1)
}

// -----------------------------------------------------------------------------

func Back() {
	_app._go(-1)
}

// -----------------------------------------------------------------------------

func Start(h ...Handler) {
	if len(h) > 0 {
		h[0].Serve(_app.currentURL, _app.currentURL, "")
	} else if _app.handler != nil {
		_app.handler.Serve(_app.currentURL, _app.currentURL, "")
	}

	select {}
}
