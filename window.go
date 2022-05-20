//go:build js && wasm

package js

import "fmt"

func Alert(a ...any) {
	global.Call("alert", fmt.Sprint(a...))
}

// -----------------------------------------------------------------------------

func Alertf(format string, a ...any) {
	global.Call("alert", fmt.Sprintf(format, a...))
}

// -----------------------------------------------------------------------------

func Redirect(url string) {
	global.Get("location").Set("href", url)
}
