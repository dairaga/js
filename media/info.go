//go:build js && wasm

package media

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

// Kind represents the kind of media device.
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

// DeviceInfo is Javascript MediaDeviceInfo.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/MediaDeviceInfo.
type DeviceInfo js.Value

func (i DeviceInfo) JSValue() js.Value {
	return js.Value(i)
}

// -----------------------------------------------------------------------------

// ID returns device identifier.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/MediaDeviceInfo/deviceId.
func (i DeviceInfo) ID() string {
	return js.Value(i).Get("deviceId").String()
}

// -----------------------------------------------------------------------------

// GroupID returns group identifier.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/MediaDeviceInfo/groupId.
func (i DeviceInfo) GroupID() string {
	return js.Value(i).Get("groupId").String()
}

// -----------------------------------------------------------------------------

// Kind returns an enumerated value that represents the kind of media device.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/MediaDeviceInfo/kind.
func (i DeviceInfo) Kind() Kind {
	return Kind(js.Value(i).Get("kind").String())
}

// -----------------------------------------------------------------------------

// Label returns label describing the device.
//
// https://developer.mozilla.org/en-US/docs/Web/API/MediaDeviceInfo/label
func (i DeviceInfo) Label() string {
	return js.Value(i).Get("label").String()
}

// -----------------------------------------------------------------------------

// DeviceInfoOf wraps a Javascript MediaDeviceInfo object.
func DeviceInfoOf(v js.Value) DeviceInfo {
	if !builtin.MediaDeviceInfo.Is(v) {
		panic(js.ValueError{
			Method: "InfoOf",
			Type:   v.Type(),
		})
	}

	return DeviceInfo(v)
}
