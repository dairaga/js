//go:build js && wasm

package js

import (
	"syscall/js"

	"github.com/dairaga/js/v2/builtin"
)

type Value = js.Value
type Type = js.Type
type Func = js.Func
type JSFunc = func(Value, []Value) any

type Wrapper interface {
	JSValue() Value
}

// -----------------------------------------------------------------------------

func ValueOf(x any) Value {
	switch v := x.(type) {
	case Wrapper:
		return v.JSValue()
	case Appendable:
		return v.Ref()
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

func GoBytes(src Value) []byte {
	if !builtin.IsUint8Array(src) {
		panic("src is not an Uint8Array")
	}

	size := src.Get("byteLength").Int()
	dst := make([]byte, size)
	js.CopyBytesToGo(dst, src)
	return dst
}

// -----------------------------------------------------------------------------

func Uint8Array(src []byte) Value {
	dst := builtin.Uint8Array.New(len(src))
	js.CopyBytesToJS(dst, src)
	return dst
}

// -----------------------------------------------------------------------------

func ArrayBufferToBytes(src Value) []byte {
	if !builtin.IsArrayBuffer(src) {
		panic("src is not an ArrayBuffer")
	}

	return GoBytes(builtin.Uint8Array.New(src))
}
