//go:build js && wasm

package audio

import (
	"github.com/dairaga/js/v3"
	"github.com/dairaga/js/v3/builtin"
	"github.com/dairaga/js/v3/media"
)

// SourceNode is Javascript MediaStreamAudioSourceNode.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/MediaStreamAudioSourceNode.
type SourceNode interface {
	Node

	// MediaStream returns the value of mediaStream property.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/MediaStreamAudioSourceNode/mediaStream.
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

// SourceNodeOf converts to an MediaStreamAudioSourceNode from given Javascript value v.
func SourceNodeOf(v js.Value) SourceNode {
	if !builtin.MediaStreamAudioSourceNode.Is(v) {
		panic(js.ValueError{
			Method: "SourceNodeOf",
			Type:   v.Type(),
		})
	}
	return &sourceNode{node: node(v)}
}
