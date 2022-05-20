//go:build js && wasm

package builtin

import (
	"fmt"
	"syscall/js"
)

// -----------------------------------------------------------------------------

func Panic(format string, a ...any) {
	panic(fmt.Sprintf(format, a...))
}

// -----------------------------------------------------------------------------

func ToGoBytes(arr js.Value) []byte {
	size := arr.Get("length").Int()
	dst := make([]byte, size)
	js.CopyBytesToGo(dst, arr)
	return dst
}

// -----------------------------------------------------------------------------

func ToUint8Array(src []byte) js.Value {
	dst := Uint8Array.New(len(src))
	js.CopyBytesToJS(dst, src)
	return dst
}

// -----------------------------------------------------------------------------

func IsArrayBufferView(v js.Value) bool {
	return Int8Array.InstanceOf(v) || Uint8Array.InstanceOf(v) || Uint8ClampedArray.InstanceOf(v) ||
		Int16Array.InstanceOf(v) || Uint16Array.InstanceOf(v) ||
		Int32Array.InstanceOf(v) || Uint32Array.InstanceOf(v) ||
		Float32Array.InstanceOf(v) || Float64Array.InstanceOf(v) ||
		BigInt64Array.InstanceOf(v) || BigUint64Array.InstanceOf(v)
}
