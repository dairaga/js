//go:build js && wasm

package js

import (
	"fmt"
	"syscall/js"
)

type IntervalID = js.Value

// -----------------------------------------------------------------------------

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

// -----------------------------------------------------------------------------

func SetInterval(fn js.Func, ms int, args ...any) IntervalID {
	if ms < 4 {
		ms = 0
	}

	newArgs := append([]any{fn, ms}, args...)

	return global.Call("setInterval", newArgs...)
}

// -----------------------------------------------------------------------------

func ClearInterval(id IntervalID) {
	global.Call("clearInterval", id)
}
