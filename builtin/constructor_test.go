//go:build js && wasm

package builtin

import (
	"fmt"
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {

	target := EventTarget.New()

	assert.True(t, target.InstanceOf(js.Global().Get("EventTarget")))
	assert.True(t, EventTarget.Is(target))
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
		EventTarget.Is,
		true,
	},
	{
		Event.New("abc"),
		EventTarget.Is,
		false,
	},
	{
		js.Global().Get("CustomEvent").New("abc"),
		Event.Is,
		true,
	},
	{
		EventTarget.New(),
		Event.Is,
		false,
	},
	{
		Int8Array.New(1),
		IsTypedArray,
		true,
	},
	{
		Uint8Array.New(1),
		IsTypedArray,
		true,
	},
	{
		Uint8ClampedArray.New(1),
		IsTypedArray,
		true,
	},
	{
		Int16Array.New(1),
		IsTypedArray,
		true,
	},
	{
		Uint16Array.New(1),
		IsTypedArray,
		true,
	},
	{
		Int32Array.New(1),
		IsTypedArray,
		true,
	},
	{
		Uint32Array.New(1),
		IsTypedArray,
		true,
	},
	{
		Float32Array.New(1),
		IsTypedArray,
		true,
	},
	{
		Float64Array.New(1),
		IsTypedArray,
		true,
	},
	{
		BigInt64Array.New(1),
		IsTypedArray,
		true,
	},
	{
		BigUint64Array.New(1),
		IsTypedArray,
		true,
	},
	{
		EventTarget.New(),
		IsTypedArray,
		false,
	},
	{
		Uint8Array.New(1),
		Uint8Array.Is,
		true,
	},
	{
		Uint8ClampedArray.New(1),
		Uint8Array.Is,
		false,
	},
	{
		Int16Array.New(1),
		Uint8Array.Is,
		false,
	},
	{
		ArrayBuffer.New(1),
		ArrayBuffer.Is,
		true,
	},
	{
		Int16Array.New(1),
		ArrayBuffer.Is,
		false,
	},
}

func TestIs(t *testing.T) {
	for i, data := range testData {
		assert.Equal(t, data.ans, data.fn(data.value), fmt.Sprintf("at %d", i))
	}
}
