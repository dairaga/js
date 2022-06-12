//go:build js && wasm

package app

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/url"
)

// -----------------------------------------------------------------------------

type Handler interface {
	Serve(curURL url.URL, curHash, oldHash string)
}

// -----------------------------------------------------------------------------

type HandlerFunc func(url.URL, string, string)

func (f HandlerFunc) Serve(curURL url.URL, curHash, oldHash string) {
	f(curURL, curHash, oldHash)
}

// -----------------------------------------------------------------------------

type app struct {
	window     js.Value
	currentURL url.URL
	handler    Handler
	hashFunc   js.Func
}

var _app *app

// -----------------------------------------------------------------------------

func (a *app) init() {
	cb := js.FuncOf(func(_ js.Value, args []js.Value) any {
		if a.handler != nil {
			old := url.New(args[0].Get("oldURL").String())
			new := url.New(args[0].Get("newURL").String())
			a.handler.Serve(a.currentURL, new.Hash(), old.Hash())
		}
		return nil
	})
	a.hashFunc = cb

	a.window.Call("addEventListener", "hashchange", cb)

	a.currentURL = url.New(a.window.Get("location").Get("href").String())
}

// -----------------------------------------------------------------------------

func (a *app) changeHash(new string) {
	a.currentURL.SetHash(new)
	js.Redirect(a.currentURL.String())
}

// -----------------------------------------------------------------------------

func Init(h ...Handler) {
	var handler Handler = nil

	if len(h) > 0 {
		handler = h[0]
	}

	_app = &app{
		window:     js.Window(),
		currentURL: url.New(js.Window().Get("location").Get("href").String()),
		handler:    handler,
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

func Start(h ...Handler) {
	if len(h) > 0 {
		h[0].Serve(_app.currentURL, _app.currentURL.Hash(), "")
	} else if _app.handler != nil {
		_app.handler.Serve(_app.currentURL, _app.currentURL.Hash(), "")
	}

	select {}
}
