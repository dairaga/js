//go:build js && wasm

package builtin

import "syscall/js"

type Constructor string

func (c Constructor) JSValue() js.Value {
	return js.Global().Get(string(c))
}

// -----------------------------------------------------------------------------

func (c Constructor) New(args ...any) js.Value {
	return c.JSValue().New(args...)
}

// -----------------------------------------------------------------------------

const (
	EventTarget = Constructor("EventTarget")
	Event       = Constructor("Event")

	ArrayBuffer       = Constructor("ArrayBuffer")
	Int8Array         = Constructor("Int8Array")
	Uint8Array        = Constructor("Uint8Array")
	Uint8ClampedArray = Constructor("Uint8ClampedArray")
	Int16Array        = Constructor("Int16Array")
	Uint16Array       = Constructor("Uint16Array")
	Int32Array        = Constructor("Int32Array")
	Uint32Array       = Constructor("Uint32Array")
	Float32Array      = Constructor("Float32Array")
	Float64Array      = Constructor("Float64Array")
	BigInt64Array     = Constructor("BigInt64Array")
	BigUint64Array    = Constructor("BigUint64Array")

	URL            = Constructor("URL")
	XMLHttpRequest = Constructor("XMLHttpRequest")

	FormData = Constructor("FormData")
	Blob     = Constructor("Blob")

	WebSocket = Constructor("WebSocket")

	Element = Constructor("Element")

	HTMLInputElement = Constructor("HTMLInputElement")
)

// -----------------------------------------------------------------------------

func IsEventTarget(v js.Value) bool {
	return v.InstanceOf(EventTarget.JSValue())
}

// -----------------------------------------------------------------------------

func IsEvent(v js.Value) bool {
	return v.InstanceOf(Event.JSValue())
}

// -----------------------------------------------------------------------------

func IsArrayBuffer(v js.Value) bool {
	return v.InstanceOf(ArrayBuffer.JSValue())
}

// -----------------------------------------------------------------------------

func IsArrayBufferView(v js.Value) bool {
	return v.InstanceOf(Int8Array.JSValue()) || v.InstanceOf(Uint8Array.JSValue()) || v.InstanceOf(Uint8ClampedArray.JSValue()) ||
		v.InstanceOf(Int16Array.JSValue()) || v.InstanceOf(Uint16Array.JSValue()) ||
		v.InstanceOf(Int32Array.JSValue()) || v.InstanceOf(Uint32Array.JSValue()) ||
		v.InstanceOf(Float32Array.JSValue()) || v.InstanceOf(Float64Array.JSValue()) ||
		v.InstanceOf(BigInt64Array.JSValue()) || v.InstanceOf(BigUint64Array.JSValue())
}

// -----------------------------------------------------------------------------

func IsUint8Array(v js.Value) bool {
	return v.InstanceOf(Uint8Array.JSValue()) || v.InstanceOf(Uint8ClampedArray.JSValue())
}

// -----------------------------------------------------------------------------

func IsElement(v js.Value) bool {
	return v.InstanceOf(Element.JSValue())
}

// -----------------------------------------------------------------------------

func IsInputElement(v js.Value) bool {
	return v.InstanceOf(HTMLInputElement.JSValue())
}
