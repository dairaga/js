//go:build js && wasm

package audio

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

type WorkletNode interface {
	Node
	Parameters() ParamMap
	Port() js.MessagePort
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

// -----------------------------------------------------------------------------

func NewWorklet(ctx Context, name string) WorkletNode {
	val := builtin.AudioWorkletNode.New(ctx.JSValue(), name)
	return &workletNode{node: node(val)}
}

// -----------------------------------------------------------------------------

func WorkleNodetOf(v js.Value) WorkletNode {
	if !builtin.AudioWorkletNode.Is(v) {
		panic(js.ValueError{
			Method: "WorkleNodetOf",
			Type:   v.Type(),
		})
	}
	return &workletNode{node: node(v)}
}
