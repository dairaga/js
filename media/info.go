//go:build js && wasm

package media

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

type Kind string

func (k Kind) String() string {
	return string(k)
}

const (
	VideoInput  Kind = "videoinput"
	AudioInput  Kind = "audioinput"
	AudioOutput Kind = "audiooutput"
)

// -----------------------------------------------------------------------------

type Info js.Value

func (i Info) JSValue() js.Value {
	return js.Value(i)
}

// -----------------------------------------------------------------------------

func (i Info) ID() string {
	return js.Value(i).Get("deviceId").String()
}

// -----------------------------------------------------------------------------

func (i Info) GroupID() string {
	return js.Value(i).Get("groupId").String()
}

// -----------------------------------------------------------------------------

func (i Info) Kind() Kind {
	return Kind(js.Value(i).Get("kind").String())
}

// -----------------------------------------------------------------------------

func (i Info) Label() string {
	return js.Value(i).Get("label").String()
}

// -----------------------------------------------------------------------------

func InfoOf(v js.Value) Info {
	if !builtin.MediaDeviceInfo.Is(v) {
		panic(js.ValueError{
			Method: "InfoOf",
			Type:   v.Type(),
		})
	}

	return Info(v)
}
