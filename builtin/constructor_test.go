//go:build js && wasm

package builtin

import (
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {

	target := EventTarget.New()

	assert.True(t, target.InstanceOf(js.Global().Get("EventTarget")))
	assert.True(t, IsEventTarget(target))

}

// -----------------------------------------------------------------------------

type TestData struct {
	value js.Value
	fn    func(js.Value) bool
	ans   bool
}

var testData = []TestData{
	{
		EventTarget.New(),
		IsEventTarget,
		true,
	},
	{
		Event.New("abc"),
		IsEventTarget,
		false,
	},
	{
		js.Global().Get("CustomEvent").New("abc"),
		IsEvent,
		true,
	},
	{
		EventTarget.New(),
		IsEvent,
		false,
	},
	{
		Int8Array.New(1),
		IsArrayBufferView,
		true,
	},
	{
		Uint8Array.New(1),
		IsArrayBufferView,
		true,
	},
	{
		Uint8ClampedArray.New(1),
		IsArrayBufferView,
		true,
	},
	{
		Int16Array.New(1),
		IsArrayBufferView,
		true,
	},
	{
		Uint16Array.New(1),
		IsArrayBufferView,
		true,
	},
	{
		Int32Array.New(1),
		IsArrayBufferView,
		true,
	},
	{
		Uint32Array.New(1),
		IsArrayBufferView,
		true,
	},
	{
		Float32Array.New(1),
		IsArrayBufferView,
		true,
	},
	{
		Float64Array.New(1),
		IsArrayBufferView,
		true,
	},
	{
		BigInt64Array.New(1),
		IsArrayBufferView,
		true,
	},
	{
		BigUint64Array.New(1),
		IsArrayBufferView,
		true,
	},
	{
		EventTarget.New(),
		IsArrayBufferView,
		false,
	},
	{
		Uint8Array.New(1),
		IsUint8Array,
		true,
	},
	{
		Uint8ClampedArray.New(1),
		IsUint8Array,
		true,
	},
	{
		Int16Array.New(1),
		IsUint8Array,
		false,
	},
	{
		ArrayBuffer.New(1),
		IsArrayBuffer,
		true,
	},
	{
		Int16Array.New(1),
		IsArrayBuffer,
		false,
	},
}

func TestIs(t *testing.T) {
	for _, data := range testData {
		assert.Equal(t, data.ans, data.fn(data.value))
	}
}
