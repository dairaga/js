//go:build js && wasm

// Package websocket is wrapper of Javascript WebSocket API. It sets binary type to arraybuffer as default.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API.
package websocket

import (
	"fmt"

	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

// State represents websocket ready state.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/readyState.
type State uint16

const (
	Connecting State = 0
	Open       State = 1
	Closing    State = 2
	Closed     State = 3
)

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

// -----------------------------------------------------------------------------

// Handler is a callback functions collection for websocket event.
type Handler interface {
	// ServeText is called when websocket receives text message.
	ServeText(msg string)

	// ServeBinary is called when websocket receives binary message.
	ServeBinary(data []byte)

	// Opened is called when websocket is opened.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/open_event.
	Opened()

	// Failed is called when connection is closed due to an error.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/error_event.
	Failed()

	// Closed is called when connection is closed.
	//
	// https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/close_event.
	Closed()
}

// -----------------------------------------------------------------------------

// Client is a wrapper of Javascript WebSocket class.
type Client struct {
	ref      js.Value    // Javascript WebSocket instance.
	listener js.Listener // websocket listeners.
}

// -----------------------------------------------------------------------------

// JSValue returns Javascript value.
func (cli *Client) JSValue() js.Value {
	return cli.ref
}

// -----------------------------------------------------------------------------

// URL returns url string of the websocket connection.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/url.
func (cli *Client) URL() string {
	return cli.ref.Get("url").String()
}

// -----------------------------------------------------------------------------

// State returns the websocket ready state.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/readyState.
func (cli *Client) State() State {
	return State(cli.ref.Get("readyState").Int())
}

// -----------------------------------------------------------------------------

// Close disconnects the websocket.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/close.
func (cli *Client) Close() {
	cli.ref.Call("close")
	cli.listener.Release()
}

// -----------------------------------------------------------------------------

// Closed returns true if the websocket is disconnected.
func (cli *Client) Closed() bool {
	return cli.State() == Closed
}

// -----------------------------------------------------------------------------

// SendText sends text message.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/send.
//func (cli *Client) SentText(msg string) {
//	cli.ref.Call("send", msg)
//}

// -----------------------------------------------------------------------------

// Send sends message.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/send.
func (cli *Client) Send(x any) {
	switch v := x.(type) {
	case string:
		cli.ref.Call("send", v)
	case []byte:
		cli.ref.Call("send", js.Uint8Array(v))
	case js.Wrapper:
		cli.Send(v.JSValue())
	case js.Value:
		if !(builtin.ArrayBuffer.Is(v) || builtin.IsArrayBufferView(v) || builtin.Blob.Is(v)) {
			panic(js.ValueError{
				Method: "WebSocket.Send",
				Type:   v.Type(),
			})
		}
		cli.ref.Call("send", v)
	default:
		panic(fmt.Sprintf("unsupported type: %T", x))
	}
}

// -----------------------------------------------------------------------------

// Connect connects to the service with the given url. Given a handler to serve all websocket events.
func Connect(url string, handler Handler) *Client {
	ret := &Client{
		ref:      builtin.WebSocket.New(url),
		listener: make(map[string]js.Func),
	}
	ret.ref.Set("binaryType", "arraybuffer")

	ret.listener.Add(ret.ref, "close", func(_ js.Value, _ []js.Value) any {
		handler.Closed()
		return nil
	})

	ret.listener.Add(ret.ref, "open", func(_ js.Value, _ []js.Value) any {
		handler.Opened()
		return nil
	})

	ret.listener.Add(ret.ref, "error", func(_ js.Value, _ []js.Value) any {
		handler.Failed()
		return nil
	})

	ret.listener.Add(ret.ref, "message", func(_this js.Value, args []js.Value) any {
		data := args[0].Get("data")
		if builtin.ArrayBuffer.Is(data) {
			handler.ServeBinary(js.GoBytes(builtin.Uint8Array.New(data)))
		} else {
			handler.ServeText(data.String())
		}
		return nil
	})

	return ret
}
