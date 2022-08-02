//go:build js && wasm

package media

import (
	"github.com/dairaga/js/v3"
	"github.com/dairaga/js/v3/builtin"
)

// Stream is a Javascript MediaStream.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/MediaStream.
type Stream interface {
	js.Wrapper

	// ID returns a unique identifier for the stream.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/MediaStream/id
	ID() string

	// Active returns true if the stream is active.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/MediaStream/active
	Active() bool
}

// -----------------------------------------------------------------------------

type stream js.Value

// -----------------------------------------------------------------------------

func (s stream) JSValue() js.Value {
	return js.Value(s)
}

// -----------------------------------------------------------------------------

func (s stream) ID() string {
	return js.Value(s).Get("id").String()
}

// -----------------------------------------------------------------------------

func (s stream) Active() bool {
	return js.Value(s).Get("active").Bool()
}

// -----------------------------------------------------------------------------

// StreamOf converts to a MediaStream from given Javascript value v.
func StreamOf(v js.Value) Stream {
	if !builtin.MediaStream.Is(v) {
		panic(js.ValueError{
			Method: "StreamOf",
			Type:   v.Type(),
		})
	}

	return stream(v)
}
