//go:build js && wasm

// Package builtin includes common Javascript class.
package builtin

import "syscall/js"

// Constructor represents a Javascript class constructor.
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

func (c Constructor) Is(v js.Value) bool {
	return v.InstanceOf(c.JSValue())
}

// -----------------------------------------------------------------------------

const (
	Error          = Constructor("Error")          // Javascript Error class.
	EvalError      = Constructor("EvalError")      // Javascript EvalError class.
	RangeError     = Constructor("RangeError")     // Javascript RangeError class.
	ReferenceError = Constructor("ReferenceError") // Javascript ReferenceError class.
	SyntaxError    = Constructor("SyntaxError")    // Javascript SyntaxError class.
	TypeError      = Constructor("TypeError")      // Javascript TypeError class.
	URIError       = Constructor("URIError")       // Javascript URIError class.

	Object      = Constructor("Object")      // Javascript Object class.
	Promise     = Constructor("Promise")     // Javascript Promise class.
	MessagePort = Constructor("MessagePort") // Javascript MessagePort class.

	EventTarget = Constructor("EventTarget") // Javascript EventTarget class.
	Event       = Constructor("Event")       // Javascript Event class.

	DataView          = Constructor("DataView")          // Javascript DataView class.
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

	URL             = Constructor("URL")             // Javascript URL class.
	URLSearchParams = Constructor("URLSearchParams") // Javascript URLSearchParams class.
	XMLHttpRequest  = Constructor("XMLHttpRequest")  // Javascript XMLHttpRequest class.

	HTMLFormElement = Constructor("HTMLFormElement") // Javascript HTMLFormElement class.
	FormData        = Constructor("FormData")        // Javascript FormData class.

	WebSocket = Constructor("WebSocket") // Javascript WebSocket class.

	Element = Constructor("Element") // Javascript Element class.

	HTMLInputElement    = Constructor("HTMLInputElement")    // Javascript HTMLInputElement class.
	HTMLSelectElement   = Constructor("HTMLSelectElement")   // Javascript HTMLSelectElement class.
	HTMLTextAreaElement = Constructor("HTMLTextAreaElement") // Javascript HTMLTextAreaElement class.
	HTMLTemplateElement = Constructor("HTMLTemplateElement") // Javascript HTMLTemplateElement class.

	ReadableStream = Constructor("ReadableStream") // Javascript ReadableStream class.
	Blob           = Constructor("Blob")           // Javascript Blob class.
	File           = Constructor("File")           // Javascript File class.

	RegExp = Constructor("RegExp") // Javascript RegExp class.

	Array = Constructor("Array") // Javascript Array class.

	MediaSource                = Constructor("MediaSource")                // Javascript MediaSource class.
	MediaStream                = Constructor("MediaStream")                // Javascript MediaStream class.
	MediaDeviceInfo            = Constructor("MediaDeviceInfo")            // Javascript MediaDeviceInfo class.
	MediaStreamAudioSourceNode = Constructor("MediaStreamAudioSourceNode") // Javascript MediaStreamAudioSourceNode class.

	AudioContext          = Constructor("AudioContext")          // Javascript AudioContext class.
	AnalyserNode          = Constructor("AnalyserNode")          // Javascript AnalyserNode class.
	AudioDestinationNode  = Constructor("AudioDestinationNode")  // Javascript AudioDestinationNode class.
	AudioWorklet          = Constructor("AudioWorklet")          // Javascript AudioWorklet class.
	AudioWorkletNode      = Constructor("AudioWorkletNode")      // Javascript AudioWorkletNode class.
	AudioWorkletProcessor = Constructor("AudioWorkletProcessor") // Javascript AudioWorkletProcessor class.

	Headers  = Constructor("Headers")  // Javascript Headers class.
	Request  = Constructor("Request")  // Javascript Request class.
	Response = Constructor("Response") // Javascript Response class.
)

// -----------------------------------------------------------------------------

func Is(v js.Value, c Constructor) bool {
	return c.Is(v)
}

// -----------------------------------------------------------------------------

func In(v js.Value, first Constructor, others ...Constructor) bool {
	if Is(v, first) {
		return true
	}
	for _, c := range others {
		if Is(v, c) {
			return true
		}
	}
	return false
}

// -----------------------------------------------------------------------------

// IsTypedArray returns true if given instance v is an typed array.
func IsTypedArray(v js.Value) bool {
	return v.InstanceOf(Int8Array.JSValue()) || v.InstanceOf(Uint8Array.JSValue()) || v.InstanceOf(Uint8ClampedArray.JSValue()) ||
		v.InstanceOf(Int16Array.JSValue()) || v.InstanceOf(Uint16Array.JSValue()) ||
		v.InstanceOf(Int32Array.JSValue()) || v.InstanceOf(Uint32Array.JSValue()) ||
		v.InstanceOf(Float32Array.JSValue()) || v.InstanceOf(Float64Array.JSValue()) ||
		v.InstanceOf(BigInt64Array.JSValue()) || v.InstanceOf(BigUint64Array.JSValue())
}

// -----------------------------------------------------------------------------

// HasValueProperty returns true if given instance v has a `value` property.
func HasValueProperty(v js.Value) bool {
	return v.InstanceOf(HTMLInputElement.JSValue()) ||
		v.InstanceOf(HTMLSelectElement.JSValue()) ||
		v.InstanceOf(HTMLTextAreaElement.JSValue())
}

/*
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

// -----------------------------------------------------------------------------

// IsMediaStream returns true if given instance v is a MediaStream.
func IsMediaStream(v js.Value) bool {
	return v.InstanceOf(MediaStream.JSValue())
}

// -----------------------------------------------------------------------------

// IsMediaDeviceInfo returns true if given instance v is a MediaDeviceInfo.
func IsMediaDeviceInfo(v js.Value) bool {
	return v.InstanceOf(MediaDeviceInfo.JSValue())
}

// -----------------------------------------------------------------------------

// IsPromise returns true if given instance v is a Promise.
func IsPromise(v js.Value) bool {
	return v.InstanceOf(Promise.JSValue())
}
*/
