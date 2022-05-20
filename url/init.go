//go:build js && wasm

package url

import "syscall/js"

var constructor = js.Value{}

func init() {
	if x := js.Global().Get("URL"); x.Truthy() {
		constructor = x
	} else if x := js.Global().Get("webkitURL"); x.Truthy() {
		constructor = x
	} else {
		panic("Window.URL not supported")
	}
}
