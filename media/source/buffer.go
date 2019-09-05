package source

import "github.com/dairaga/js"

// Buffer represents javascript SourceBuffer from MediaSource. https://developer.mozilla.org/en-US/docs/Web/API/SourceBuffer
type Buffer struct {
	ref js.Value
}

// BufferOf returns a buffer.
func BufferOf(v js.Value) *Buffer {
	return &Buffer{ref: v}
}

// JSValue ...
func (buf *Buffer) JSValue() js.Value {
	return buf.ref
}

// Append adds bytes to buffer.
func (buf *Buffer) Append(raw []byte) *Buffer {
	//typedArr := js.TypedArrayOf(raw)
	typedArr := js.ToJSUint8Array(raw)
	buf.ref.Call("appendBuffer", typedArr)
	//typedArr.Release()

	return buf
}
