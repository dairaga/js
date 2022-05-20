//go:build js && wasm

package app

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/url"
)

var currentURL = url.New(js.Global().Get("location").Get("href").String())

func URL() url.URL {
	return currentURL
}

// -----------------------------------------------------------------------------

func SetHash(hash string) {
	currentURL.SetHash(hash)
	js.Global().Get("location").Set("href", currentURL.String())
}
