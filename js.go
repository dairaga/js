//go:build js && wasm

package js

import "syscall/js"

type Value = js.Value
type Func = js.Func
type Type = js.Func
type Error = js.Error

type JSFunc = func(Value, []Value) any

var global = js.Global()

// -----------------------------------------------------------------------------

type Wrapper interface {
	JSValue() Value
}

// -----------------------------------------------------------------------------

func ValueOf(x any) Value {
	switch v := x.(type) {
	case Wrapper:
		return v.JSValue()
	case Value:
		return v
	default:
		return js.ValueOf(x)
	}
}

// -----------------------------------------------------------------------------

func FuncOf(fn JSFunc) Func {
	return js.FuncOf(fn)
}

// -----------------------------------------------------------------------------

func ParseInt(val string, radix int) (int, bool) {
	x := global.Call("parseInt", val, radix)
	if x.IsNaN() {
		return 0, false
	}

	return x.Int(), true
}

// -----------------------------------------------------------------------------

func ParseFloat(val string, radix int) (float64, bool) {
	x := global.Call("parseFloat", val, radix)
	if x.IsNaN() {
		return 0.0, false
	}

	return x.Float(), true
}
