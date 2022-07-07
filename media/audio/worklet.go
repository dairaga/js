//go:build js && wasm

package audio

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

type Worklet js.Value

// -----------------------------------------------------------------------------

func (w Worklet) JSValue() js.Value {
	return js.Value(w)
}

// -----------------------------------------------------------------------------

func (w Worklet) AddModule(url string, c ...js.Credential) js.Promise {
	if len(c) > 0 {
		return js.PromiseOf(js.Value(w).Call("addModule", url, js.Obj{"credentials": c[0].String()}))
	}
	return js.PromiseOf(js.Value(w).Call("addModule", url))
}

// -----------------------------------------------------------------------------

func WorkletOf(v js.Value) Worklet {
	if !builtin.AudioWorklet.Is(v) {
		panic(js.ValueError{
			Method: "WorkletOf",
			Type:   v.Type(),
		})
	}
	return Worklet(v)
}
