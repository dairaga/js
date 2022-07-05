//go:build js && wasm

// Packages media provides common media-related functionality.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/MediaDevices
package media

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

var navigator = js.Window().Get("navigator")
var devices = navigator.Get("mediaDevices")

// -----------------------------------------------------------------------------

// Authorize is to request permission to use media devices.
// It is an async function and returns a Promise.
func Authorize(video, audio bool) js.Promise {
	return js.PromiseOf(devices.Call("getUserMedia", js.Obj{"video": video, "audio": audio}))
}

// -----------------------------------------------------------------------------

// AwaitAuthorize returns true if user authorizes to access devices.
func AwaitAuthorize(video, audio bool) bool {
	/*ch := make(chan bool)

	Authorize(video, audio).Then(func(result js.Value) any {
		ch <- true
		return nil
	}).Catch(func(err js.Value) any {
		ch <- false
		return nil
	}).Finally(func() any {
		close(ch)
		return nil
	})

	return ch*/

	result := Authorize(video, audio).Await()
	return result.Truthy() && builtin.MediaStream.Is(result)
}

// -----------------------------------------------------------------------------

// EnumerateDevices is to get media devices user authorized.
// It is an async function and put resulted to given callback function fn.
func EnumerateDevices(fn func([]DeviceInfo), fails ...func(err js.Error)) {
	js.PromiseOf(devices.Call("enumerateDevices")).Then(func(result js.Value) any {
		size := result.Length()
		infos := make([]DeviceInfo, size)
		for i := 0; i < size; i++ {
			infos[i] = DeviceInfoOf(result.Index(i))
		}
		fn(infos)
		return nil
	}).Catch(func(err js.Value) any {
		valErr := js.Error{Value: err}
		for i := range fails {
			fails[i](valErr)
		}
		return nil
	})
}

// -----------------------------------------------------------------------------

// AwaitEnumerateDevices is to get media devices user authorized.
func AwaitEnumerateDevices() []DeviceInfo {

	ch := make(chan []DeviceInfo)
	defer func() {
		close(ch)
	}()

	EnumerateDevices(
		func(infos []DeviceInfo) {
			ch <- infos
		},
		func(js.Error) {
			ch <- nil
		},
	)
	return <-ch
}

// -----------------------------------------------------------------------------

// OnDevices is to listen to media devices.
func OnDevicecChange(fn func([]DeviceInfo)) {
	devices.Call("addEventListener", "devicechange", js.FuncOf(func(_ js.Value, args []js.Value) any {
		EnumerateDevices(fn)
		return nil
	}))
}

// -----------------------------------------------------------------------------

// GetUserMedia returns a Promise of a MediaStream object contains a video and/or audio track.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/MediaDevices/getUserMedia
func GetUserMedia(opt js.Obj) js.Promise {
	return js.PromiseOf(devices.Call("getUserMedia", opt))
}

// -----------------------------------------------------------------------------

// AwaitGetUserMedia returns a MediaStream object that contains a video and/or audio track.
func AwaitGetUserMedia(opt js.Obj) <-chan Stream {
	ch := make(chan Stream)

	GetUserMedia(opt).Then(func(result js.Value) any {
		ch <- StreamOf(result)
		return nil
	}).Catch(func(err js.Value) any {
		ch <- nil
		return nil
	}).Finally(func() any {
		close(ch)
		return nil
	})

	return ch
}
