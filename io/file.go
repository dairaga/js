//go:build js && wasm

package io

import (
	"fmt"
	"time"

	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

type File js.Value

func (f File) JSValue() js.Value {
	return js.Value(f)
}

// -----------------------------------------------------------------------------

func (f File) Name() string {
	return js.Value(f).Get("name").String()
}

// -----------------------------------------------------------------------------

func (f File) Type() string {
	return js.Value(f).Get("type").String()
}

// -----------------------------------------------------------------------------

func (f File) LastModified() time.Time {
	m := int64(js.Value(f).Get("lastModified").Int())
	return time.Unix(0, m*int64(time.Millisecond))
}

// -----------------------------------------------------------------------------

func (f File) WebkitRelativePath() string {
	return js.Value(f).Get("webkitRelativePath").String()
}

// -----------------------------------------------------------------------------

func fileSupported(v js.Value) bool {
	return v.Truthy() && (builtin.IsArrayBuffer(v) ||
		builtin.IsArrayBufferView(v) ||
		builtin.IsBlob(v))
}

// -----------------------------------------------------------------------------

func FileOf(x any, opts ...string) File {
	name := ""
	opt := map[string]any{
		"type": "",
	}

	switch len(opts) {
	case 1:
		name = opts[0]
	case 2:
		name = opts[0]
		opt["type"] = opts[1]
	}

	val := js.Null()
	switch v := x.(type) {
	case string:
		return File(builtin.File.New([]any{v}, name, opt))
	case []byte:
		return File(builtin.File.New([]any{js.Uint8Array(v)}, name, opt))
	case js.Wrapper:
		val = v.JSValue()
	case js.Value:
		val = v
	}

	if val.Truthy() {
		if builtin.IsFile(val) {
			return File(val)
		}

		if fileSupported(val) {
			return File(builtin.File.New([]any{val}, name, opt))
		}

	}

	panic(fmt.Sprintf("unsupported type %T", x))
}
