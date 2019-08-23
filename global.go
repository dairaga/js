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

// ----------------------------------------------------------------------------

// IsNaN ...
func IsNaN(v js.Value) bool {
	return global.Call("isNaN", v).Bool()
}

// ParseInt ...
func ParseInt(val string, radix int) (int, bool) {
	x := global.Call("parseInt", val, radix)
	if IsNaN(x) {
		return 0, false
	}
	return x.Int(), true
}

// ParseFloat ...
func ParseFloat(val string) (float64, bool) {
	x := global.Call("parseFloat", val)
	if IsNaN(x) {
		return 0.0, false
	}

	return x.Float(), true
}
