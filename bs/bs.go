//go:build js && wasm

package bs

import "github.com/dairaga/js/v2"

var global = js.Window().Get("bootstrap")

func Global() js.Value {
	return global
}
