//go:build js && wasm

package js

import (
	"fmt"
	"time"

	"github.com/dairaga/js/v2/builtin"
)

type File Value

func (f File) JSValue() Value {
	return Value(f)
}

// -----------------------------------------------------------------------------

func (f File) Name() string {
	return Value(f).Get("name").String()
}

// -----------------------------------------------------------------------------

func (f File) Type() string {
	return Value(f).Get("type").String()
}

// -----------------------------------------------------------------------------

func (f File) LastModified() time.Time {
	m := int64(Value(f).Get("lastModified").Int())
	return time.Unix(0, m*int64(time.Millisecond))
}

// -----------------------------------------------------------------------------

func (f File) WebkitRelativePath() string {
	return Value(f).Get("webkitRelativePath").String()
}

// -----------------------------------------------------------------------------

func fileSupported(v Value) bool {
	return v.Truthy() && (builtin.ArrayBuffer.Is(v) ||
		builtin.IsTypedArray(v) ||
		builtin.Blob.Is(v))
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

	val := Null()
	switch v := x.(type) {
	case string:
		return File(builtin.File.New([]any{v}, name, opt))
	case []byte:
		return File(builtin.File.New([]any{Uint8Array(v)}, name, opt))
	case Wrapper:
		val = v.JSValue()
	case Value:
		val = v
	}

	if val.Truthy() {
		if builtin.File.Is(val) {
			return File(val)
		}

		if fileSupported(val) {
			return File(builtin.File.New([]any{val}, name, opt))
		}

	}

	panic(fmt.Sprintf("unsupported type %T", x))
}
