//go:build js && wasm

package audio

import (
	"github.com/dairaga/js/v3"
	"github.com/dairaga/js/v3/builtin"
)

// DestinationNode is Javascript AudioDestinationNode.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/AudioDestinationNode.
type DestinationNode interface {
	Node

	// MaxChannelCount returns the maximum amount of channels that the physical device can handle.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/AudioDestinationNode/maxChannelCount.
	MaxChannelCount() uint64
}

// -----------------------------------------------------------------------------

type destinationNode struct {
	node
}

var _ DestinationNode = &destinationNode{}

// -----------------------------------------------------------------------------

func (d *destinationNode) MaxChannelCount() uint64 {
	return uint64(d.JSValue().Get("maxChannelCount").Int())
}

// -----------------------------------------------------------------------------

// DestinationNodeOf converts to an AudioDestinationNode from given Javascript value v.
func DestinationNodeOf(v js.Value) DestinationNode {
	if !builtin.AudioDestinationNode.Is(v) {
		panic(js.ValueError{
			Method: "DestinationNodeOf",
			Type:   v.Type(),
		})
	}
	return &destinationNode{node: node(v)}
}
