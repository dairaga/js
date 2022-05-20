//go:build js && wasm

package io

import (
	"encoding/json"

	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

type Blob interface {
	builtin.Wrapper
	Type() string
	Size() int
}

type blob js.Value

func (b blob) JSValue() js.Value {
	return js.Value(b)
}

// -----------------------------------------------------------------------------

func (b blob) Type() string {
	return b.JSValue().Get("type").String()
}

// -----------------------------------------------------------------------------

func (b blob) Size() int {
	return b.JSValue().Get("size").Int()
}

// -----------------------------------------------------------------------------

func BlobOf(x any, mine ...string) Blob {
	switch v := x.(type) {
	case string:
		return blob(builtin.Blob.New(v))
	case []byte:
		// convert go bytes to js Uint8Array
		arr := builtin.ToUint8Array(v)
		return blob(builtin.Blob.New(arr))
	case builtin.Wrapper:
		return BlobOf(v.JSValue())
	case js.Value:
		switch v.Type() {
		case js.TypeNumber:
			// is a number and make a array buffer
			return blob(builtin.Blob.New(builtin.ArrayBuffer.New(v.Int())))
		case js.TypeString:
			// is a string
			return blob(builtin.Blob.New(v))
		}

		// is an ArrayBuffer
		if builtin.ArrayBuffer.InstanceOf(v) {
			return blob(builtin.Blob.New(v))
		}

		// is an ArrayBufferView
		if builtin.IsArrayBufferView(v) {
			return blob(builtin.Blob.New(v))
		}

		// is an Blob
		if builtin.Blob.InstanceOf(v) {
			return blob(v)
		}
	}

	xbytes, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	return BlobOf(xbytes)
}
