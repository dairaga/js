//go:build js && wasm

package audio

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
	"github.com/dairaga/js/v2/media"
)

type SourceNode interface {
	Node
	MediaStream() media.Stream
}

type sourceNode struct {
	node
}

var _ SourceNode = &sourceNode{}

// -----------------------------------------------------------------------------

func (s *sourceNode) MediaStream() media.Stream {
	return media.StreamOf(s.JSValue().Get("mediaStream"))
}

// -----------------------------------------------------------------------------

func SourceNodeOf(v js.Value) SourceNode {
	if !builtin.MediaStreamAudioSourceNode.Is(v) {
		panic(js.ValueError{
			Method: "SourceNodeOf",
			Type:   v.Type(),
		})
	}
	return &sourceNode{node: node(v)}
}
