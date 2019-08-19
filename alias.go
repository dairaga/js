package js

import "syscall/js"

// Value alias syscall/js.Value
type Value = js.Value

// Func alias syscall/js.Func
type Func = js.Func

// ValueOf ...
func ValueOf(x interface{}) Value {
	return js.ValueOf(x)
}

// FuncOf ...
func FuncOf(fn func(js.Value, []js.Value) interface{}) Func {
	return js.FuncOf(fn)
}
