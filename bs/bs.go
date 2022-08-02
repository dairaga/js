//go:build js && wasm

// Package bs is a helper toolkit for Bootstrap 5.0.
package bs

import "github.com/dairaga/js/v3"

var global = js.Window().Get("bootstrap") // an bootstrap object.

// Global returns the global Bootstrap object.
func Global() js.Value {
	return global
}
