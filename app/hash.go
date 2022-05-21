//go:build js && wasm

package app

import (
	"syscall/js"

	"github.com/dairaga/js/v2/url"
)

type HashHandler interface {
	ServeHash(url url.URL, oldHash, newHash string)
}

// -----------------------------------------------------------------------------

func ServHash(h HashHandler) {

	window := js.Global()

	cb := js.FuncOf(func(_this js.Value, args []js.Value) any {
		oldURL := url.New(args[0].Get("oldURL").String())
		newURL := url.New(args[0].Get("newURL").String())
		h.ServeHash(newURL, oldURL.Hash(), newURL.Hash())
		return nil
	})

	window.Call("addEventListener", "hashchange", cb)
	curURL := url.New(window.Get("location").Get("href").String())

	h.ServeHash(curURL, "", curURL.Hash())
}
