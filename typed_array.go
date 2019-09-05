// +build js,wasm

package js

import (
	"syscall/js"
)

var (
	int8Array    = Global().Get("Int8Array")
	int16Array   = Global().Get("Int16Array")
	int32Array   = Global().Get("Int32Array")
	uint8Array   = Global().Get("Uint8Array")
	uint16Array  = Global().Get("Uint16Array")
	uint32Array  = Global().Get("Uint32Array")
	float32Array = Global().Get("Float32Array")
	float64Array = Global().Get("Float64Array")

	// not all browsers support.
	int64Array  = Global().Get("BigInt64Array")  // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/BigInt64Array
	uint64Array = Global().Get("BigUint64Array") // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/BigUint64Array
)

// ToGoBytes copy javascript uint8Array content to go bytes.
func ToGoBytes(uint8Array Value) []byte {
	size := uint8Array.Length()

	dest := make([]byte, size)

	js.CopyBytesToGo(dest, uint8Array)

	return dest
}

// ToJSUint8Array copy go bytes to javascript uint8Array
func ToJSUint8Array(src []byte) Value {
	dest := New("Uint8Array", len(src))

	js.CopyBytesToJS(dest, src)

	return dest
}

// TypedArrayOf return a javascript typed array.
// x must be must be []int8, []int16, []int32, []uint8, []uint16, []uint32, []float32 and []float64.
func TypedArrayOf(x interface{}) js.Value {
	switch v := x.(type) {
	case []int8:
		len := len(v)
		if len <= 0 {
			return int8Array.New(0)
		}
		arr := make([]interface{}, len)
		for i, z := range v {
			arr[i] = z
		}
		return int8Array.New(arr)
	case []int16:
		len := len(v)
		if len <= 0 {
			return int16Array.New(0)
		}
		arr := make([]interface{}, len)
		for i, z := range v {
			arr[i] = z
		}
		return int16Array.New(arr)
	case []int32:
		len := len(v)
		if len <= 0 {
			return int32Array.New(0)
		}
		arr := make([]interface{}, len)
		for i, z := range v {
			arr[i] = z
		}
		return int32Array.New(arr)
	/*case []int64:
	len := len(v)
	if len <= 0 {
		return int64Array.New(0)
	}
	arr := make([]interface{}, len)
	for i, z := range v {
		arr[i] = z
	}
	return int64Array.New(arr)*/
	case []uint8:
		/*len := len(v)
		if len <= 0 {
			return uint8Array.New(0)
		}
		arr := make([]interface{}, len)
		for i, z := range v {
			arr[i] = z
		}
		return uint8Array.New(arr)*/
		return ToJSUint8Array(v)
	case []uint16:
		len := len(v)
		if len <= 0 {
			return uint16Array.New(0)
		}
		arr := make([]interface{}, len)
		for i, z := range v {
			arr[i] = z
		}
		return uint16Array.New(arr)
	case []uint32:
		len := len(v)
		if len <= 0 {
			return uint32Array.New(0)
		}
		arr := make([]interface{}, len)
		for i, z := range v {
			arr[i] = z
		}
		return uint32Array.New(arr)
	/*case []uint64:
	len := len(v)
	if len <= 0 {
		return uint32Array.New(0)
	}
	arr := make([]interface{}, len)
	for i, z := range v {
		arr[i] = z
	}
	return uint32Array.New(arr)*/
	case []float32:
		len := len(v)
		if len <= 0 {
			return float32Array.New(0)
		}
		arr := make([]interface{}, len)
		for i, z := range v {
			arr[i] = z
		}
		return float32Array.New(arr)
	case []float64:
		len := len(v)
		if len <= 0 {
			return float64Array.New(0)
		}
		arr := make([]interface{}, len)
		for i, z := range v {
			arr[i] = z
		}
		return float64Array.New(arr)
	default:
		panic("data type not supported")
	}
}
