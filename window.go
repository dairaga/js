//go:build js && wasm

package js

import (
	"fmt"
	"syscall/js"
)

var global = js.Global()

func Window() js.Value {
	return global
}

// -----------------------------------------------------------------------------

func Alert(a ...any) {
	global.Call("alert", fmt.Sprint(a...))
}

// -----------------------------------------------------------------------------

func Alertf(format string, a ...any) {
	global.Call("alert", fmt.Sprintf(format, a...))
}

// -----------------------------------------------------------------------------

func Redirect(newURL string) {
	global.Get("location").Set("href", newURL)
}
