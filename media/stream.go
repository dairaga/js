//go:build js && wasm

package media

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

type Stream interface {
	js.Wrapper
	ID() string
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

func StreamOf(v js.Value) Stream {
	if !builtin.MediaStream.Is(v) {
		panic(js.ValueError{
			Method: "StreamOf",
			Type:   v.Type(),
		})
	}

	return stream(v)
}
