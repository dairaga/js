package websocket

import (
	"fmt"
	"syscall/js"

	djs "github.com/dairaga/js"
)

// State ...
type State uint16

func (s State) String() string {
	switch s {
	case Connecting:
		return "CONNECTING"
	case Open:
		return "OPEN"
	case Closing:
		return "CLOSING"
	case Closed:
		return "CLOSED"
	default:
		return "UNKNOW"
	}
}

// websocket constant
const (
	Connecting State = 0
	Open       State = 1
	Closing    State = 2
	Closed     State = 3
)

// WebSocket ...
type WebSocket struct {
	ref js.Value

	onOpen    js.Func
	onClose   js.Func
	onError   js.Func
	onMessage js.Func
	onText    func(string)
	onBinary  func([]byte)
}

// JSValue ...
func (ws WebSocket) JSValue() js.Value {
	return ws.ref
}

// URL ...
func (ws WebSocket) URL() string {
	return ws.ref.Get("url").String()
}

// ReadyState ...
func (ws WebSocket) ReadyState() State {
	return State(ws.ref.Get("readyState").Int())
}

// OnOpen ...
func (ws WebSocket) OnOpen(cb func(djs.Event)) WebSocket {
	ws.onOpen = js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		cb(djs.EventOf(args[0]))
		return nil
	})
	ws.ref.Set("onopen", ws.onOpen)
	return ws
}

// OnClose ...
func (ws WebSocket) OnClose(cb func(CloseEvent)) WebSocket {
	ws.onClose = js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		cb(CloseEventOf(args[0]))
		return nil
	})

	ws.ref.Set("onclose", ws.onClose)
	return ws
}

// OnError ...
func (ws WebSocket) OnError(cb func(djs.Event)) WebSocket {
	ws.onError = js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		cb(djs.EventOf(args[0]))
		return nil
	})

	ws.ref.Set("onerror", ws.onError)
	return ws
}

// OnMessage ...
func (ws WebSocket) OnMessage(cb func([]byte)) WebSocket {
	ws.onMessage = js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		data := args[0].Get("data")

		if data.InstanceOf(arrayBuffer) {
			cb(toBytes(data))
		} else {
			cb([]byte(data.String()))
		}
		return nil
	})

	ws.ref.Set("onmessage", ws.onMessage)
	return ws
}

// Closed ...
func (ws WebSocket) Closed() bool {
	s := ws.ReadyState()
	return s == Closed || s == Closing
}

// Close ...
func (ws WebSocket) Close() {
	ws.ref.Call("close")
}

// Release ...
func (ws WebSocket) Release() {
	if !ws.Closed() {
		ws.Close()
	}

	if ws.onOpen.Truthy() {
		ws.onOpen.Release()
	}

	if ws.onClose.Truthy() {
		ws.onClose.Release()
	}

	if ws.onError.Truthy() {
		ws.onError.Release()
	}

	if ws.onMessage.Truthy() {
		ws.onMessage.Release()
	}
}

// SendText ...
func (ws WebSocket) SendText(data string) {
	ws.ref.Call("send", data)
}

// SendBinary data must be []int8, []int16, []int32, []uint8, []uint16, []uint32, []float32 and []float64.
func (ws WebSocket) SendBinary(data interface{}) {
	switch v := data.(type) {
	case []int8, []int16, []int32, []uint8, []uint16, []uint32, []float32, []float64:
		arr := js.TypedArrayOf(v)
		ws.ref.Call("send", arr)
		arr.Release()
	default:
		fmt.Println("send data type not supported!")
	}
}

// ----------------------------------------------------------------------------

var stringValue = js.Global().Get("String")
var arrayBuffer = js.Global().Get("ArrayBuffer")

func toBytes(v js.Value) []byte {
	size := v.Get("byteLength")
	if !size.Truthy() {
		return nil
	}

	ret := make([]uint8, size.Int())
	destArray := js.TypedArrayOf(ret)

	srcArray := js.Global().Get("Uint8Array").New(v)

	destArray.Call("set", srcArray, 0)

	destArray.Release()

	return []byte(ret)
}

// Connect ...
func Connect(url string) WebSocket {
	ws := WebSocket{ref: js.Global().Get("WebSocket").New(url)}
	ws.ref.Set("binaryType", "arraybuffer")
	return ws
}
