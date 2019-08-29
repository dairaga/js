// +build js,wasm

package js

import (
	"syscall/js"
)

// Bytes convert javascript array buffer to Go byte slice.
func Bytes(arrayBuffer Value) []byte {
	size := arrayBuffer.Get("byteLength")
	if !size.Truthy() {
		return nil
	}

	ret := make([]uint8, size.Int())
	destArray := TypedArrayOf(ret)

	srcArray := New("Uint8Array", arrayBuffer)

	destArray.Call("set", srcArray, 0)

	destArray.Release()

	return []byte(ret)
}

// TypedArrayOf return a javascript typed array.
// x must be must be []int8, []int16, []int32, []uint8, []uint16, []uint32, []float32 and []float64.
func TypedArrayOf(x interface{}) js.TypedArray {
	switch v := x.(type) {
	case []int8, []int16, []int32, []uint8, []uint16, []uint32, []float32, []float64:
		return js.TypedArrayOf(v)
	default:
		panic("data type not supported")
	}
}
