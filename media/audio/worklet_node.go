//go:build js && wasm

package audio

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

// WorkletNode is Javascript AudioWorkletNode.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/AudioWorkletNode.
type WorkletNode interface {
	Node // inherits from AudioNode.

	// Parameters returns the parameters of the worklet.
	Parameters() ParamMap

	// Port returns the Javascript MessagePort of the worklet.
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

// NewWorklet creates a new AudioWorkletNode with given AudioContext and module name.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/AudioWorkletNode/AudioWorkletNode.
func NewWorkletNode(ctx Context, name string) WorkletNode {
	val := builtin.AudioWorkletNode.New(ctx.JSValue(), name)
	return &workletNode{node: node(val)}
}

// -----------------------------------------------------------------------------

// WorkleNodetOf converts to an AudioWorkletNode from given Javascript value v.
func WorkleNodetOf(v js.Value) WorkletNode {
	if !builtin.AudioWorkletNode.Is(v) {
		panic(js.ValueError{
			Method: "WorkleNodetOf",
			Type:   v.Type(),
		})
	}
	return &workletNode{node: node(v)}
}
