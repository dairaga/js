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
	KindVideoInput  Kind = "videoinput"
	KindAudioInput  Kind = "audioinput"
	KindAudioOutput Kind = "audiooutput"
)

// -----------------------------------------------------------------------------

type DeviceInfo js.Value

func (i DeviceInfo) JSValue() js.Value {
	return js.Value(i)
}

// -----------------------------------------------------------------------------

func (i DeviceInfo) ID() string {
	return js.Value(i).Get("deviceId").String()
}

// -----------------------------------------------------------------------------

func (i DeviceInfo) GroupID() string {
	return js.Value(i).Get("groupId").String()
}

// -----------------------------------------------------------------------------

func (i DeviceInfo) Kind() Kind {
	return Kind(js.Value(i).Get("kind").String())
}

// -----------------------------------------------------------------------------

func (i DeviceInfo) Label() string {
	return js.Value(i).Get("label").String()
}

// -----------------------------------------------------------------------------

func DeviceInfoOf(v js.Value) DeviceInfo {
	if !builtin.MediaDeviceInfo.Is(v) {
		panic(js.ValueError{
			Method: "InfoOf",
			Type:   v.Type(),
		})
	}

	return DeviceInfo(v)
}
