//go:build js && wasm

package audio

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

type DestinationNode interface {
	Node
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

func DestinationNodeOf(v js.Value) DestinationNode {
	if !builtin.AudioDestinationNode.Is(v) {
		panic(js.ValueError{
			Method: "DestinationNodeOf",
			Type:   v.Type(),
		})
	}
	return &destinationNode{node: node(v)}
}
