package js

import "syscall/js"

var (
	global = js.Global()
)

// Global returns javascript global.
func Global() Value {
	return global
}

// ----------------------------------------------------------------------------

// New returns a javascript object.
func New(constructor string, args ...interface{}) Value {
	return global.Get(constructor).New(args...)
}

// Call invoke a global function.
func Call(fn string, args ...interface{}) Value {
	return global.Call(fn, args...)
}

// RegisterFunc ...
func RegisterFunc(name string, fn Func) {
	global.Set(name, fn)
}
