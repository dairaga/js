package io

import (
	"github.com/dairaga/js"
)

// FileReader represents javascript FileReader.
type FileReader struct {
	js.EventTarget
}

// NewFileReader returns a file reader.
func NewFileReader() FileReader {
	return FileReader{js.EventTargetOf(js.New("FileReader"))}
}

// Read https://developer.mozilla.org/en-US/docs/Web/API/FileReader/readAsArrayBuffer
func (r FileReader) Read(file File) {
	r.JSValue().Call("readAsArrayBuffer", file.ref)
}

// Done https://developer.mozilla.org/en-US/docs/Web/API/FileReader/onload
func (r FileReader) Done(h func([]byte)) {
	cb := js.FuncOf(func(_this js.Value, _ []js.Value) interface{} {
		result := js.ValueOf(_this.Get("result"))
		h(js.Bytes(result))
		return nil
	})

	r.On("load", cb)
}

// Fail https://developer.mozilla.org/en-US/docs/Web/API/FileReader/error_event
func (r FileReader) Fail(h func(js.Event)) {
	cb := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		h(js.EventOf(args[0]))
		return nil
	})

	r.On("error", cb)
}

// Always https://developer.mozilla.org/en-US/docs/Web/API/FileReader/loadend_event
func (r FileReader) Always(h func(js.Event)) {
	cb := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		h(js.EventOf(args[0]))
		return nil
	})

	r.On("loadend", cb)
}

// Progress https://developer.mozilla.org/en-US/docs/Web/API/FileReader/progress_event
func (r FileReader) Progress(h func(int, int, bool)) {
	cb := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		h(args[0].Get("loaded").Int(), args[0].Get("total").Int(), args[0].Get("lengthComputable").Bool())
		return nil
	})

	r.On("progress", cb)
}
