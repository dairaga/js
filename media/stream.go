package media

import "github.com/dairaga/js"

// Stream https://developer.mozilla.org/en-US/docs/Web/API/MediaStream
type Stream struct {
	ref js.Value
}

// StreamOf returns a media stream.
func StreamOf(v js.Value) *Stream {
	return &Stream{ref: v}
}

// JSValue ...
func (stream *Stream) JSValue() js.Value {
	return stream.ref
}

// Ready ...
func (stream *Stream) Ready() bool {
	return stream.ref.Truthy()
}
