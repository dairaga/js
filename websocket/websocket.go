package websocket

import (
	"github.com/dairaga/js"
)

// State represents websocket ready state.
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

// websocket ready state valeus.
const (
	Connecting State = 0
	Open       State = 1
	Closing    State = 2
	Closed     State = 3
)

// WebSocket represents javascript web socket component.
type WebSocket struct {
	ref      *js.EventTarget
	onText   func(string)
	onBinary func([]byte)
}

// URL returns current connecting url.
func (ws *WebSocket) URL() string {
	return ws.ref.JSValue().Get("url").String()
}

// ReadyState returns current ready state.
func (ws *WebSocket) ReadyState() State {
	return State(ws.ref.JSValue().Get("readyState").Int())
}

// OnOpen add callback function when opening.
func (ws *WebSocket) OnOpen(cb func(*js.Event)) *WebSocket {
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		cb(js.EventOf(args[0]))
		return nil
	})
	ws.ref.On("open", fn)
	return ws
}

// OnClose add callback function when closed.
func (ws *WebSocket) OnClose(cb func(*CloseEvent)) *WebSocket {
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		cb(CloseEventOf(args[0]))
		return nil
	})

	ws.ref.On("close", fn)
	return ws
}

// OnError add callback function when error happening.
func (ws *WebSocket) OnError(cb func(*js.Event)) *WebSocket {
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		cb(js.EventOf(args[0]))
		return nil
	})

	ws.ref.On("error", fn)
	return ws
}

// OnMessage add callback function when receiving message.
func (ws *WebSocket) OnMessage(cb func([]byte)) *WebSocket {
	fn := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		data := args[0].Get("data")

		if data.InstanceOf(arrayBuffer) {
			cb(js.Bytes(data))
		} else {
			cb([]byte(data.String()))
		}
		return nil
	})

	ws.ref.On("message", fn)
	return ws
}

// Closed returns boolean that indicates whether or not the connection was closed.
func (ws *WebSocket) Closed() bool {
	s := ws.ReadyState()
	return s == Closed
}

// Close closes the connection.
func (ws *WebSocket) Close() {
	ws.ref.JSValue().Call("close")
}

// Release closes connection and frees up resources.
func (ws *WebSocket) Release() {
	if !ws.Closed() {
		ws.Close()
	}
	ws.ref.Release()
}

// SendText sends text message to server.
func (ws *WebSocket) SendText(data string) *WebSocket {
	ws.ref.JSValue().Call("send", data)
	return ws
}

// SendBinary sends binary data to server. Parameter data must be []int8, []int16, []int32, []uint8, []uint16, []uint32, []float32 and []float64.
func (ws *WebSocket) SendBinary(data interface{}) *WebSocket {
	arr := js.TypedArrayOf(data)
	ws.ref.JSValue().Call("send", arr)
	arr.Release()
	return ws
}

// ----------------------------------------------------------------------------

var stringValue = js.Global().Get("String")
var arrayBuffer = js.Global().Get("ArrayBuffer")

// Connect connects to server.
func Connect(url string) *WebSocket {
	ws := &WebSocket{ref: js.EventTargetOf(js.Global().Get("WebSocket").New(url))}
	ws.ref.JSValue().Set("binaryType", "arraybuffer")
	return ws
}
