package source

import "github.com/dairaga/js"

var _src = js.Global().Get("MediaSource")

// ----------------------------------------------------------------------------

// TypeSupported returns a Boolean value which is true if the given MIME type is likely to be supported by the current user agent.
func TypeSupported(mine string) bool {
	return _src.Call("isTypeSupported", mine).Bool()
}

// ----------------------------------------------------------------------------

// Source represents javascript MediaSource. https://developer.mozilla.org/en-US/docs/Web/API/MediaSource
type Source struct {
	ref js.Value
}

// New returns a media source.
func New() *Source {
	return &Source{ref: _src.New()}
}

// JSValue ...
func (s *Source) JSValue() js.Value {
	return s.ref
}

// AddSourceBuffer creates a new SourceBuffer of the given MIME type and adds it to the MediaSource's sourceBuffers list.
func (s Source) AddSourceBuffer(mine string) *Buffer {
	v := s.ref.Call("addSourceBuffer", mine)
	return BufferOf(v)
}
