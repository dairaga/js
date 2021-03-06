package media

import (
	"github.com/dairaga/js"
)

var (
	window    = js.Global().Get("window")
	navigator = window.Get("navigator")
	stream    *Stream
)

// StreamConstrains ...
type StreamConstrains struct {
	Audio bool
	Video bool
}

func (constrains StreamConstrains) toJSObject() map[string]interface{} {
	return map[string]interface{}{
		"audio": constrains.Audio,
		"video": constrains.Video,
	}
}

// GetUserMedia ...
func GetUserMedia(constrains StreamConstrains, success func(*Stream), fail func(*js.Error)) {
	if stream != nil && stream.Ready() {
		success(stream)
		return
	}

	promise := js.Value{}
	if d := navigator.Get("mediaDevices"); d.Truthy() {
		promise = d.Call("getUserMedia", constrains.toJSObject())
	} else if navigator.Get("getUserMedia").Truthy() {
		promise = navigator.Call("getUserMedia", constrains.toJSObject())
	} else if navigator.Get("webkitGetUserMedia").Truthy() {
		promise = navigator.Call("webkitGetUserMedia", constrains.toJSObject())
	} else if navigator.Get("mozGetUserMedia").Truthy() {
		promise = navigator.Call("mozGetUserMedia", constrains.toJSObject())
	} else if navigator.Get("msGetUserMedia").Truthy() {
		promise = navigator.Call("msGetUserMedia", constrains.toJSObject())
	} else {
		fail(js.ErrorOf(js.ValueOf("user media unsupported!")))
		return
	}

	promise.Call("then", js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		stream = StreamOf(args[0])
		success(stream)
		return nil
	}))
	promise.Call("catch", js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		fail(js.ErrorOf(args[0]))
		return nil
	}))
}
