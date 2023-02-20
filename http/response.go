//go:build js && wasm

package http

import (
	"fmt"

	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

type Response js.Value

// -----------------------------------------------------------------------------

func (r Response) JSValue() js.Value {
	return js.Value(r)
}

// -----------------------------------------------------------------------------

func (r Response) String() string {
	return fmt.Sprintf(`{"code": %d, "message": %q}`, r.Status(), r.StatusText())
}

// -----------------------------------------------------------------------------

func (r Response) Used() bool {
	return js.Value(r).Get("bodyUsed").Bool()
}

// -----------------------------------------------------------------------------

func (r Response) Headers() Headers {
	return Headers(js.Value(r).Get("headers"))
}

// -----------------------------------------------------------------------------

func (r Response) OK() bool {
	return js.Value(r).Get("ok").Bool()
}

// -----------------------------------------------------------------------------

func (r Response) Redirected() bool {
	return js.Value(r).Get("redirected").Bool()
}

// -----------------------------------------------------------------------------

func (r Response) Status() int {
	return js.Value(r).Get("status").Int()
}

// -----------------------------------------------------------------------------

func (r Response) StatusText() string {
	return js.Value(r).Get("statusText").String()
}

// -----------------------------------------------------------------------------

func (r Response) Type() string {
	return js.Value(r).Get("type").String()
}

// -----------------------------------------------------------------------------

func (r Response) URL() string {
	return js.Value(r).Get("url").String()
}

// -----------------------------------------------------------------------------

func (r Response) ArrayBuffer() js.Promise {
	return js.PromiseOf(js.Value(r).Call("arrayBuffer"))
}

// -----------------------------------------------------------------------------

func (r Response) Blob() js.Promise {
	return js.PromiseOf(js.Value(r).Call("blob"))
}

// -----------------------------------------------------------------------------

func (r Response) FormData() js.Promise {
	return js.PromiseOf(js.Value(r).Call("formData"))
}

// -----------------------------------------------------------------------------

func (r Response) JSON() js.Promise {
	return js.PromiseOf(js.Value(r).Call("json"))
}

// -----------------------------------------------------------------------------

func (r Response) Text() js.Promise {
	return js.PromiseOf(js.Value(r).Call("text"))
}

// -----------------------------------------------------------------------------

func (r Response) Clone() Response {
	return Response(js.Value(r).Call("clone"))
}

// -----------------------------------------------------------------------------

func Redirect(url string, status ...int) Response {
	for _, s := range status {
		return Response(builtin.Response.JSValue().Call("redirect", url, s))
	}
	return Response(builtin.Response.JSValue().Call("redirect", url))
}
