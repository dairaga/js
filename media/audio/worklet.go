//go:build js && wasm

package audio

import "github.com/dairaga/js/v2"

type Worklet js.Value

// -----------------------------------------------------------------------------

func (w Worklet) JSValue() js.Value {
	return js.Value(w)
}

// -----------------------------------------------------------------------------

func (w Worklet) AddModule(url string, c ...js.Credential) js.Promise {
	if len(c) > 0 {
		return js.PromiseOf(js.Value(w).Call("addModule", url, js.Obj{"credentials": c[0]}))
	}
	return js.PromiseOf(js.Value(w).Call("addModule", url))
}

// -----------------------------------------------------------------------------

type WorkletNode interface {
	Node
	Parameters() ParamMap
}

// -----------------------------------------------------------------------------

type workletNode struct {
	node
}

var _ WorkletNode = &workletNode{}

// -----------------------------------------------------------------------------

func (n workletNode) Parameters() ParamMap {
	return ParamMap(n.JSValue().Get("parameters"))
}

// -----------------------------------------------------------------------------

func (n workletNode) Port() js.MessagePort {
	return js.MessagePort(n.JSValue().Get("port"))
}
