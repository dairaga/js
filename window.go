package js

import "fmt"

// window

var window = global.Get("window")

// Window ...
func Window() Value {
	return window
}

// Alert ...
func Alert(a ...interface{}) {
	window.Call("alert", fmt.Sprint(a...))
}

// Alertf ...
func Alertf(format string, a ...interface{}) {
	window.Call("alert", fmt.Sprintf(format, a...))
}

// Confirm ...
func Confirm(a ...interface{}) bool {
	return window.Call("confirm", fmt.Sprint(a...)).Bool()
}

// Confirmf ...
func Confirmf(format string, a ...interface{}) bool {
	return window.Call("confirm", fmt.Sprintf(format, a...)).Bool()
}
