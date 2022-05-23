//go:build js && wasm

package ajax

import (
	"encoding/json"

	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
	"github.com/dairaga/js/v2/xhr"
)

const (
	MimeText   string = "text/plain; charset=utf8"
	MimeJSON   string = "application/json; charset=utf8"
	MimeStream string = "application/octet-stream"

	Unsent          = 0
	Opened          = 1
	HeadersReceived = 2
	Loading         = 3
	Done            = 4

	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	PATCH  = "PATCH"
)

// -----------------------------------------------------------------------------

type HandleFunc = func(*Response, error)

// -----------------------------------------------------------------------------

type Client struct {
	ref      js.Value
	lastErr  error
	released bool
	listener js.Listener
}

// -----------------------------------------------------------------------------

func (cli *Client) JSValue() js.Value {
	return cli.ref
}

// -----------------------------------------------------------------------------

func (cli *Client) do(method string, url string, x ...any) (err error) {

	var data []byte

	switch len(x) {
	case 0:
		data = nil
	case 1:
		data, err = json.Marshal(x[0])
	default:
		data, err = json.Marshal(x)
	}

	if err != nil {
		return
	}

	var req *Request
	req, err = NewRequest(method, url, data)
	if err != nil {
		return
	}

	cli.Do(req)
	return
}

// -----------------------------------------------------------------------------

func (cli *Client) Release() {
	if cli.released {
		return
	}
	cli.released = true
	cli.listener.Release()
}

// -----------------------------------------------------------------------------

func (cli *Client) Released() bool {
	return cli.released
}

// -----------------------------------------------------------------------------

func (cli *Client) Do(req *Request) error {
	if cli.released {
		return xhr.ErrReleased
	}

	if "" != req.User && "" != req.Password {
		cli.ref.Call("open", req.method, req.url, true, req.User, req.Password)
	} else {
		cli.ref.Call("open", req.method, req.url, true)
	}

	for key, val := range req.headers {
		cli.ref.Call("setRequestHeader", key, val)
	}

	if len(req.body) > 0 {
		data := js.Uint8Array(req.body)
		cli.ref.Call("send", data)
	} else {
		cli.ref.Call("send")
	}

	return nil
}

// -----------------------------------------------------------------------------

func New(fn HandleFunc, timeout ...int64) *Client {
	cli := new(Client)
	cli.ref = builtin.XMLHttpRequest.New()
	cli.listener = make(js.Listener)
	cli.released = false

	cli.ref.Set("responseType", "arraybuffer")
	if len(timeout) > 0 {
		cli.ref.Set("timeout", timeout)
	}

	cli.listener.Add(cli.ref, "timeout", func(js.Value, []js.Value) any {
		cli.lastErr = xhr.ErrTimeout
		return nil
	})

	cli.listener.Add(cli.ref, "error", func(js.Value, []js.Value) any {
		cli.lastErr = xhr.ErrFailed
		return nil
	})

	cli.listener.Add(cli.ref, "abort", func(js.Value, []js.Value) any {
		cli.lastErr = xhr.ErrAbort
		return nil
	})

	cli.listener.Add(cli.ref, "loadend", func(js.Value, []js.Value) any {
		resp := fill(cli.ref)
		fn(resp, cli.lastErr)
		return nil
	})

	return cli
}
