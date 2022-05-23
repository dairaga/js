//go:build js && wasm

package js

import (
	"fmt"
	"syscall/js"
)

func Alert(a ...any) {
	js.Global().Call("alert", fmt.Sprint(a...))
}

// -----------------------------------------------------------------------------

func Alertf(format string, a ...any) {
	js.Global().Call("alert", fmt.Sprintf(format, a...))
}
