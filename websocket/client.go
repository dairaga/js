//go:build js && wasm

package websocket

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

// State represents websocket ready state.
type State uint16

// websocket ready state valeus.
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

type Handler interface {
	ServeText(msg string)
	ServeBinary(data []byte)
	Opened()
	Failed()
	Closed()
}

// -----------------------------------------------------------------------------

type Client struct {
	ref      js.Value
	listener js.Listener
}

// -----------------------------------------------------------------------------

func (cli *Client) JSValue() js.Value {
	return cli.ref
}

// -----------------------------------------------------------------------------

func (cli *Client) URL() string {
	return cli.ref.Get("url").String()
}

// -----------------------------------------------------------------------------

func (cli *Client) State() State {
	return State(cli.ref.Get("readyState").Int())
}

// -----------------------------------------------------------------------------

func (cli *Client) Close() {
	cli.ref.Call("close")
	cli.listener.Release()
}

// -----------------------------------------------------------------------------

func (cli *Client) Closed() bool {
	return cli.State() == Closed
}

// -----------------------------------------------------------------------------

func (cli *Client) SentText(msg string) {
	cli.ref.Call("send", msg)
}

// -----------------------------------------------------------------------------

func (cli *Client) SendBinary(buf []byte) {
	cli.ref.Call("send", js.Uint8Array(buf))
}

// -----------------------------------------------------------------------------

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
		if builtin.IsArrayBuffer(data) {
			handler.ServeBinary(js.GoBytes(builtin.Uint8Array.New(data)))
		} else {
			handler.ServeText(data.String())
		}
		return nil
	})

	return ret
}
