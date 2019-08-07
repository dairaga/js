package io

import (
	"syscall/js"
)

// FileReader ...
type FileReader struct {
	ref js.Value
}

// JSValue ...
func (r FileReader) JSValue() js.Value {
	return r.ref
}

// NewFileReader ...
func NewFileReader() FileReader {
	return FileReader{js.Global().Get("FileReader").New("FileReader")}
}

// ReadAsArrayBuffer ...
func (r FileReader) ReadAsArrayBuffer(file File) {
	r.ref.Call("readAsArrayBuffer", file.ref)
}

// OnBufferLoaded ...
func (r FileReader) OnBufferLoaded(h func([]byte)) {

	cb := js.FuncOf(func(_this js.Value, _ []js.Value) interface{} {
		result := js.ValueOf(_this.Get("result"))
		h(toBytes(result))
		return nil
	})

	r.ref.Call("addEventListener", "load", cb)
}

func toBytes(v js.Value) []byte {
	size := v.Get("byteLength")
	if !size.Truthy() {
		return nil
	}

	ret := make([]uint8, size.Int())
	destArray := js.TypedArrayOf(ret)

	srcArray := js.Global().Get("Uint8Array").New(v)

	destArray.Call("set", srcArray, 0)

	destArray.Release()

	return []byte(ret)
}
