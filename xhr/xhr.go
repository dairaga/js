package xhr

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/dairaga/js"
)

// State xhr state
type State int

// ReadyState ...
const (
	Unsent          State = 0
	Opened          State = 1
	HeadersReceived State = 2
	Loading         State = 3
	Done            State = 4
)

// content types
const (
	TextType   string = "text/plain"
	JSONType   string = "application/json"
	StreamType string = "application/octet-stream"
)

func (s State) String() string {
	switch s {
	case Unsent:
		return "Unsent"
	case Opened:
		return "Opened"
	case HeadersReceived:
		return "Headers Received"
	case Loading:
		return "Loading"
	case Done:
		return "Done"
	default:
		return fmt.Sprintf("Unknown(%d)", s)
	}
}

var xhrConstructor = js.Global().Get("XMLHttpRequest")

//-----------------------------------------------------------------------------

// Request ...
type Request struct {
	method  string
	url     string
	headers map[string]string
	body    []byte
}

// NewRequest returns a request.
func NewRequest(method, url string, body []byte) *Request {
	return &Request{
		method: method,
		url:    url,
		body:   body,
	}
}

// SetHeader set request header.
func (req *Request) SetHeader(key, value string) *Request {

	if req.headers == nil {
		req.headers = make(map[string]string)
	}

	req.headers[key] = value
	return req
}

//-----------------------------------------------------------------------------

// Response ...
type Response struct {
	code    int
	headers map[string]string
	body    []byte
}

func (resp *Response) fill(xhr js.Value) *Response {
	resp.code = xhr.Get("status").Int()

	if xhr.Get("readyState").Int() == int(Done) {
		resp.body = js.Bytes(xhr.Get("response"))
	}

	headers := xhr.Call("getAllResponseHeaders").String()

	tmp := strings.Split(headers, "\r\n")
	if size := len(tmp); size > 0 {
		resp.headers = make(map[string]string, size)
		for _, x := range tmp {
			parts := strings.Split(x, ": ")
			resp.headers[parts[0]] = strings.Join(parts[1:], ": ")
		}
	}
	return resp
}

// Code ...
func (resp *Response) Code() int {
	return resp.code
}

// Body ...
func (resp *Response) Body() []byte {
	return resp.body
}

// GetHeader ...
func (resp *Response) GetHeader(key string) string {
	if resp.headers != nil {
		return resp.headers[key]
	}
	return ""
}

//-----------------------------------------------------------------------------

// Client ...
type Client struct {
	ref        js.Value
	resp       *Response
	lastErr    error
	onDone     func(*Response)
	onFail     func(*Response)
	onAlways   func(*Response)
	onProgress func(int, int, bool)
}

// NewClient returns Client
func NewClient() Client {
	cli := Client{
		ref: xhrConstructor.New(),
	}

	cli.ref.Set("responseType", "arraybuffer")

	cli.ref.Call("addEventListener", "load", js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		if cli.onDone != nil {
			if cli.resp == nil {
				cli.resp = new(Response)
			}
			cli.onDone(cli.resp.fill(cli.ref))
		}
		return nil
	}))

	cli.ref.Call("addEventListener", "error", js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		if cli.onFail != nil {
			if cli.resp == nil {
				cli.resp = new(Response)
			}
			cli.onFail(cli.resp.fill(cli.ref))
		}
		return nil
	}))

	cli.ref.Call("addEventListener", "loadend", js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		if cli.onAlways != nil {
			if cli.resp == nil {
				cli.resp = new(Response)
				cli.resp.fill(cli.ref)
			}
			cli.onAlways(cli.resp)
		}
		return nil
	}))

	cli.ref.Call("addEventListener", "progress", js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		if cli.onProgress != nil {
			cli.onProgress(args[0].Get("loaded").Int(), args[0].Get("total").Int(), args[0].Get("lengthComputable").Bool())
		}
		return nil
	}))

	return cli
}

// JSValue ...
func (cli Client) JSValue() js.Value {
	return cli.ref
}

// ReadyState https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/readyState
func (cli Client) ReadyState() State {
	return State(cli.ref.Get("readyState").Int())
}

// Do ...
func (cli Client) Do(req *Request) Client {
	cli.ref.Call("open", req.method, req.url, true)
	if len(req.headers) > 0 {
		for k, v := range req.headers {
			cli.ref.Call("setRequestHeader", k, v)
		}
	}

	if len(req.body) > 0 {
		arr := js.TypedArrayOf(req.body)
		cli.ref.Call("send", arr)
		arr.Release()
	} else {
		cli.ref.Call("send")
	}

	return cli
}

// Done ...
func (cli Client) Done(cb func(*Response)) Client {
	cli.onDone = cb
	return cli
}

// Fail ...
func (cli Client) Fail(cb func(*Response)) Client {
	cli.onFail = cb
	return cli
}

// Always ...
func (cli Client) Always(cb func(*Response)) Client {
	cli.onAlways = cb
	return cli
}

// Progress ...
func (cli Client) Progress(cb func(int, int, bool)) Client {
	cli.onProgress = cb
	return cli
}

//-----------------------------------------------------------------------------

func handle(data interface{}) (string, []byte, error) {
	if data == nil {
		return "", nil, nil
	}

	typ := reflect.TypeOf(data)
	if typ.Kind() == reflect.Struct {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return "", nil, err
		}

		return JSONType, dataBytes, nil
	}

	switch v := data.(type) {
	case string:
		return TextType, []byte(v), nil
	case []byte:
		return StreamType, v, nil
	}

	return "", nil, fmt.Errorf("unknown type: %v", typ)
}

func do(method, url string, data interface{}) Client {
	cli := NewClient()
	ctype, dataBytes, err := handle(data)
	if err != nil {
		cli.lastErr = err
		return cli
	}

	req := NewRequest(method, url, dataBytes)
	req.SetHeader("Content-Type", ctype)
	cli.Do(req)
	return cli
}

// Get ...
func Get(url string, data interface{}) Client {
	return do("GET", url, data)
}

// Post ...
func Post(url string, data interface{}) Client {
	return do("POST", url, data)
}

// Put ...
func Put(url string, data interface{}) Client {
	return do("PUT", url, data)
}

// Delete ...
func Delete(url string, data interface{}) Client {
	return do("DELETE", url, data)
}
