//go:build js && wasm

package audio

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

// Worklet is Javascript AudioWorklet.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/AudioWorklet.
type Worklet js.Value

// -----------------------------------------------------------------------------

// JSValue returns the underlying Javascript value.
func (w Worklet) JSValue() js.Value {
	return js.Value(w)
}

// -----------------------------------------------------------------------------

// Addmodule loads the module from given url and name.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/Worklet/addModule.
// See https://developer.mozilla.org/en-US/docs/Web/API/AudioWorkletGlobalScope/registerProcessor.
func (w Worklet) AddModule(url string, c ...js.Credential) js.Promise {
	if len(c) > 0 {
		return js.PromiseOf(js.Value(w).Call("addModule", url, js.Obj{"credentials": c[0].String()}))
	}
	return js.PromiseOf(js.Value(w).Call("addModule", url))
}

// -----------------------------------------------------------------------------

// WorkletOf converts to an AudioWorklet from given Javascript value v.
func WorkletOf(v js.Value) Worklet {
	if !builtin.AudioWorklet.Is(v) {
		panic(js.ValueError{
			Method: "WorkletOf",
			Type:   v.Type(),
		})
	}
	return Worklet(v)
}
