//go:build js && wasm

package ajax

import (
	"time"

	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
	"github.com/dairaga/js/v2/xhr"
)

const (
	MimeText   string = "text/plain; charset=utf8"
	MimeJSON   string = "application/json; charset=utf8"
	MimeStream string = "application/octet-stream"

	//Unsent          = 0
	//Opened          = 1
	//HeadersReceived = 2
	//Loading         = 3
	//Done            = 4

	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	PATCH  = "PATCH"
)

var (
	defaultWithCredentials = true
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
	var req *Request
	req, err = NewRequest(method, url, x...)
	if err != nil {
		return
	}

	err = cli.Do(req)
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

func (cli *Client) WithCredentials(flag bool) {
	cli.ref.Set("withCredentials", flag)
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

func (cli *Client) Get(url string, x ...any) error {
	return cli.do(GET, url, x...)
}

// -----------------------------------------------------------------------------

func (cli *Client) Post(url string, x ...any) error {
	return cli.do(POST, url, x...)
}

// -----------------------------------------------------------------------------

func (cli *Client) Put(url string, x ...any) error {
	return cli.do(PUT, url, x...)
}

// -----------------------------------------------------------------------------

func (cli *Client) Delete(url string, x ...any) error {
	return cli.do(DELETE, url, x...)
}

// -----------------------------------------------------------------------------

func (cli *Client) Patch(url string, x ...any) error {
	return cli.do(PATCH, url, x...)
}

// -----------------------------------------------------------------------------

func New(fn HandleFunc, timeout ...time.Duration) *Client {
	cli := new(Client)
	cli.ref = builtin.XMLHttpRequest.New()
	cli.listener = make(js.Listener)
	cli.released = false

	cli.ref.Set("responseType", "arraybuffer")
	if len(timeout) > 0 {
		cli.ref.Set("timeout", timeout[0].Milliseconds())
	}
	cli.ref.Set("withCredentials", defaultWithCredentials)

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
		var resp *Response
		if cli.lastErr == nil {
			resp = fill(cli.ref)
		}
		fn(resp, cli.lastErr)
		return nil
	})

	return cli
}

// -----------------------------------------------------------------------------

func do(method, url string, fn HandleFunc, x ...any) (cli *Client, err error) {
	cli = New(fn)
	err = cli.do(method, url, x...)
	return
}

// -----------------------------------------------------------------------------

func Get(url string, fn HandleFunc, x ...any) (cli *Client, err error) {
	return do(GET, url, fn, x...)
}

// -----------------------------------------------------------------------------

func Post(url string, fn HandleFunc, x ...any) (cli *Client, err error) {
	return do(POST, url, fn, x...)
}

// -----------------------------------------------------------------------------

func Put(url string, fn HandleFunc, x ...any) (cli *Client, err error) {
	return do(PUT, url, fn, x...)
}

// -----------------------------------------------------------------------------

func Delete(url string, fn HandleFunc, x ...any) (cli *Client, err error) {
	return do(DELETE, url, fn, x...)
}

// -----------------------------------------------------------------------------

func Patch(url string, fn HandleFunc, x ...any) (cli *Client, err error) {
	return do(PATCH, url, fn, x...)
}
