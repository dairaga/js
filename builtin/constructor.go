//go:build js && wasm

// Package builtin includes common Javascript class.
package builtin

import "syscall/js"

// Constructor represents a javascript class constructor.
type Constructor string

func (c Constructor) JSValue() js.Value {
	return js.Global().Get(string(c))
}

// -----------------------------------------------------------------------------

// New returns a new corresponding instance.
func (c Constructor) New(args ...any) js.Value {
	return c.JSValue().New(args...)
}

// -----------------------------------------------------------------------------

const (
	EventTarget = Constructor("EventTarget") // Javascript EventTarget class.
	Event       = Constructor("Event")       // Javascript Event class.

	ArrayBuffer       = Constructor("ArrayBuffer")       // Javascript ArrayBuffer class.
	Int8Array         = Constructor("Int8Array")         // Javascript Int8Array class.
	Uint8Array        = Constructor("Uint8Array")        // Javascript Uint8Array class.
	Uint8ClampedArray = Constructor("Uint8ClampedArray") // Javascript Uint8ClampedArray class.
	Int16Array        = Constructor("Int16Array")        // Javascript Int16Array class.
	Uint16Array       = Constructor("Uint16Array")       // Javascript Uint16Array class.
	Int32Array        = Constructor("Int32Array")        // Javascript Int32Array class.
	Uint32Array       = Constructor("Uint32Array")       // Javascript Uint32Array class.
	Float32Array      = Constructor("Float32Array")      // Javascript Float32Array class.
	Float64Array      = Constructor("Float64Array")      // Javascript Float64Array class.
	BigInt64Array     = Constructor("BigInt64Array")     // Javascript BigInt64Array class.
	BigUint64Array    = Constructor("BigUint64Array")    // Javascript BigUint64Array class.

	URL            = Constructor("URL")            // Javascript URL class.
	XMLHttpRequest = Constructor("XMLHttpRequest") // Javascript XMLHttpRequest class.

	HTMLFormElement = Constructor("HTMLFormElement") // Javascript HTMLFormElement class.
	FormData        = Constructor("FormData")        // Javascript FormData class.

	WebSocket = Constructor("WebSocket") // Javascript WebSocket class.

	Element = Constructor("Element") // Javascript Element class.

	HTMLInputElement    = Constructor("HTMLInputElement")    // Javascript HTMLInputElement class.
	HTMLSelectElement   = Constructor("HTMLSelectElement")   // Javascript HTMLSelectElement class.
	HTMLTextAreaElement = Constructor("HTMLTextAreaElement") // Javascript HTMLTextAreaElement class.
	HTMLTemplateElement = Constructor("HTMLTemplateElement") // Javascript HTMLTemplateElement class.

	Blob = Constructor("Blob") // Javascript Blob class.
	File = Constructor("File") // Javascript File class.

	RegExp = Constructor("RegExp") // Javascript RegExp class.

	Array = Constructor("Array") // Javascript Array class.

	MediaSource = Constructor("MediaSource") // Javascript MediaSource class.
)

// -----------------------------------------------------------------------------

// IsEventTarget returns true if given instance v is an EventTarget.
func IsEventTarget(v js.Value) bool {
	return v.InstanceOf(EventTarget.JSValue())
}

// -----------------------------------------------------------------------------

// IsEvent returns true if given instance v is an Event.
func IsEvent(v js.Value) bool {
	return v.InstanceOf(Event.JSValue())
}

// -----------------------------------------------------------------------------

// IsArrayBuffer returns true if given instance v is an ArrayBuffer.
func IsArrayBuffer(v js.Value) bool {
	return v.InstanceOf(ArrayBuffer.JSValue())
}

// -----------------------------------------------------------------------------

// IsArrayBufferView returns true if given instance v is an ArrayBufferView.
func IsArrayBufferView(v js.Value) bool {
	return v.InstanceOf(Int8Array.JSValue()) || v.InstanceOf(Uint8Array.JSValue()) || v.InstanceOf(Uint8ClampedArray.JSValue()) ||
		v.InstanceOf(Int16Array.JSValue()) || v.InstanceOf(Uint16Array.JSValue()) ||
		v.InstanceOf(Int32Array.JSValue()) || v.InstanceOf(Uint32Array.JSValue()) ||
		v.InstanceOf(Float32Array.JSValue()) || v.InstanceOf(Float64Array.JSValue()) ||
		v.InstanceOf(BigInt64Array.JSValue()) || v.InstanceOf(BigUint64Array.JSValue())
}

// -----------------------------------------------------------------------------

// IsUint8Array returns true if given instance v is a Uint8Array.
func IsUint8Array(v js.Value) bool {
	return v.InstanceOf(Uint8Array.JSValue()) || v.InstanceOf(Uint8ClampedArray.JSValue())
}

// -----------------------------------------------------------------------------

// IsElement returns true if given instance v is an Element.
func IsElement(v js.Value) bool {
	return v.InstanceOf(Element.JSValue())
}

// -----------------------------------------------------------------------------

// IsInputElement returns true if given instance v is a HTMLInputElement.
func IsInputElement(v js.Value) bool {
	return v.InstanceOf(HTMLInputElement.JSValue())
}

// -----------------------------------------------------------------------------

// HasValueProperty returns true if given instance v has a `value` property.
func HasValueProperty(v js.Value) bool {
	return v.InstanceOf(HTMLInputElement.JSValue()) ||
		v.InstanceOf(HTMLSelectElement.JSValue()) ||
		v.InstanceOf(HTMLTextAreaElement.JSValue())
}

// -----------------------------------------------------------------------------

// IsFile returns true if given instance v is a File.
func IsFile(v js.Value) bool {
	return v.InstanceOf(File.JSValue())
}

// -----------------------------------------------------------------------------

// IsBlob returns true if given instance v is a Blob.
func IsBlob(v js.Value) bool {
	return v.InstanceOf(Blob.JSValue())
}

// -----------------------------------------------------------------------------

// IsForm returns true if given instance v is a HTMLFormElement.
func IsForm(v js.Value) bool {
	return v.InstanceOf(HTMLFormElement.JSValue())
}

// -----------------------------------------------------------------------------

// IsArray return s true if given instance v is an Array.
func IsArray(v js.Value) bool {
	return v.InstanceOf(Array.JSValue())
}

// -----------------------------------------------------------------------------

// isTemplate returns true if given instance v is a HTMLTemplateElement.
func IsTemplate(v js.Value) bool {
	return v.InstanceOf(HTMLTemplateElement.JSValue())
}

// -----------------------------------------------------------------------------

// IsMediaSource returns true if given instance v is a MediaSource.
func IsMediaSource(v js.Value) bool {
	return v.InstanceOf(MediaSource.JSValue())
}
