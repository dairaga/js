//go:build js && wasm

package builtin

import "syscall/js"

type Wrapper interface {
	JSValue() js.Value
}

// -----------------------------------------------------------------------------

type Constructor string

func (c Constructor) New(args ...any) js.Value {
	return js.Global().Get(string(c)).New(args...)
}

func (c Constructor) InstanceOf(v js.Value) bool {
	return v.InstanceOf(js.Global().Get(string(c)))
}

// -----------------------------------------------------------------------------

const (
	ArrayBuffer       = Constructor("ArrayBuffer")
	Int8Array         = Constructor("Int8Array")
	Uint8Array        = Constructor("Uint8Array")
	Uint8ClampedArray = Constructor("Uint8ClampedArray")
	Int16Array        = Constructor("Int16Array")
	Uint16Array       = Constructor("Uint16Array")
	Int32Array        = Constructor("Int32Array")
	Uint32Array       = Constructor("Uint32Array")
	Float32Array      = Constructor("Float32Array")
	Float64Array      = Constructor("Float64Array")
	BigInt64Array     = Constructor("BigInt64Array")
	BigUint64Array    = Constructor("BigUint64Array")
)
