// +build js,wasm

package js

import "fmt"

// window

var window = global.Get("window")

// Window ...
func Window() Value {
	return window
}

// Alert invokes window.alert function.
func Alert(a ...interface{}) {
	window.Call("alert", fmt.Sprint(a...))
}

// Alertf invokes window.alert function.
func Alertf(format string, a ...interface{}) {
	window.Call("alert", fmt.Sprintf(format, a...))
}

// Confirm invokes window.confirm function.
func Confirm(a ...interface{}) bool {
	return window.Call("confirm", fmt.Sprint(a...)).Bool()
}

// Confirmf invokes window.confirm function.
func Confirmf(format string, a ...interface{}) bool {
	return window.Call("confirm", fmt.Sprintf(format, a...)).Bool()
}
