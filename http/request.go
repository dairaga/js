//go:build js && wasm

package http

import (
	"fmt"

	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
	"github.com/dairaga/js/v2/form"
	"github.com/dairaga/js/v2/json"
)

type Mode string

func (m Mode) String() string {
	return string(m)
}

// -----------------------------------------------------------------------------

const (
	ModeSameOrigin Mode = "same-origin"
	ModeNoCORS     Mode = "no-cors"
	ModeCORS       Mode = "cors"
	ModeNavigate   Mode = "navigate"
	ModeWebSocket  Mode = "websocket"
)

// -----------------------------------------------------------------------------

type Credentials string // "omit", "same-origin", "include"

func (c Credentials) String() string {
	return string(c)
}

// -----------------------------------------------------------------------------

const (
	CredentialsOmit       Credentials = "omit"
	CredentialsSameOrigin Credentials = "same-origin"
	CredentialsInclude    Credentials = "include"
)

// -----------------------------------------------------------------------------

type CacheMode string // "default", "no-store", "reload", "no-cache", "force-cache", "only-if-cached"

func (c CacheMode) String() string {
	return string(c)
}

// -----------------------------------------------------------------------------

const (
	CacheDefault      CacheMode = "default"
	CacheNoStore      CacheMode = "no-store"
	CacheReload       CacheMode = "reload"
	CacheNoCache      CacheMode = "no-cache"
	CacheForce        CacheMode = "force-cache"
	CacheOnlyIfCached CacheMode = "only-if-cached"
)

// -----------------------------------------------------------------------------

type RedirectMode string

func (m RedirectMode) String() string {
	return string(m)
}

// -----------------------------------------------------------------------------

const (
	RedirectFollow RedirectMode = "follow"
	RedirectError  RedirectMode = "error"
	RedirectManual RedirectMode = "manual"
)

// -----------------------------------------------------------------------------

type Option struct {
	Method      string
	Headers     Headers
	Body        js.Value
	Mode        Mode
	Credentials Credentials
	Cache       CacheMode
	Redirect    RedirectMode
	Referrer    string
	Integrity   string
}

// -----------------------------------------------------------------------------

func (opt *Option) JSObject() map[string]any {
	return map[string]any{
		"method":      opt.Method,
		"headers":     opt.Headers.JSValue(),
		"body":        opt.Body,
		"mode":        opt.Mode.String(),
		"credentials": opt.Credentials.String(),
		"cache":       opt.Cache.String(),
		"redirect":    opt.Redirect.String(),
		"referrer":    opt.Referrer,
		"integrity":   opt.Integrity,
	}
}

// -----------------------------------------------------------------------------

func (opt *Option) String() string {
	return fmt.Sprint(opt.JSObject())
}

// -----------------------------------------------------------------------------

func (opt *Option) JSValue() js.Value {
	return js.ValueOf(opt.JSObject())
}

// -----------------------------------------------------------------------------

func validBody(v js.Value) bool {
	if v.Type() == js.TypeString {
		return true
	}

	if builtin.IsTypedArray(v) {
		return true
	}

	return builtin.In(v,
		builtin.Blob,
		builtin.ArrayBuffer,
		builtin.DataView,
		builtin.FormData,
		builtin.URLSearchParams,
		builtin.ReadableStream)
}

// -----------------------------------------------------------------------------

func NewOption(method string, body any, headers ...[2]string) (*Option, error) {
	bodyVal := js.Null()
	typ := ""

	if body != nil {
		switch v := body.(type) {
		case js.Value:
			bodyVal = v
		case js.Wrapper:
			bodyVal = v.JSValue()
		case json.Marshaler:
			if val, err := v.MarshalValue(); err != nil {
				return nil, err
			} else {
				bodyVal = val
			}
		case []byte:
			bodyVal = js.Uint8Array(v)
			typ = "application/octet-stream"
		case string:
			bodyVal = js.ValueOf(v)
		default:
			panic(fmt.Sprintf("unsupported body type: %T", body))
		}
	}

	opt := DefaultOpt()

	if validBody(bodyVal) {
		opt.Body = bodyVal
	} else {
		opt.Body = js.ValueOf(json.Stringify(bodyVal))
		typ = "application/json; charset=utf-8"
	}

	if typ != "" {
		opt.Headers.Set("Content-Type", typ)
	}

	opt.Method = method
	opt.Headers.Add(headers...)

	return opt, nil
}

// -----------------------------------------------------------------------------

func DefaultOpt() *Option {
	return &Option{
		Method:      "GET",
		Headers:     NewHeaders(),
		Body:        js.Null(),
		Mode:        ModeCORS,
		Credentials: CredentialsSameOrigin,
		Cache:       CacheDefault,
		Redirect:    RedirectFollow,
		Referrer:    "bout:client",
		Integrity:   "",
	}
}

// -----------------------------------------------------------------------------

type Request js.Value

// -----------------------------------------------------------------------------

func (r Request) JSValue() js.Value {
	return js.Value(r)
}

// -----------------------------------------------------------------------------

func (r Request) Used() bool {
	return js.Value(r).Get("bodyUsed").Bool()
}

// -----------------------------------------------------------------------------

func (r Request) Cache() CacheMode {
	return CacheMode(js.Value(r).Get("cache").String())
}

// -----------------------------------------------------------------------------

func (r Request) Credentials() Credentials {
	return Credentials(js.Value(r).Get("credentials").String())
}

// -----------------------------------------------------------------------------

func (r Request) Destination() string {
	return js.Value(r).Get("destination").String()
}

// -----------------------------------------------------------------------------

func (r Request) Headers() Headers {
	return Headers(js.Value(r).Get("headers"))
}

// -----------------------------------------------------------------------------

func (r Request) Integrity() string {
	return js.Value(r).Get("integrity").String()
}

// -----------------------------------------------------------------------------

func (r Request) Method() string {
	return js.Value(r).Get("method").String()
}

// -----------------------------------------------------------------------------

func (r Request) Mode() Mode {
	return Mode(js.Value(r).Get("mode").String())
}

// -----------------------------------------------------------------------------

func (r Request) Redirect() RedirectMode {
	return RedirectMode(js.Value(r).Get("redirect").String())
}

// -----------------------------------------------------------------------------

func (r Request) Referrer() string {
	return js.Value(r).Get("referrer").String()
}

// -----------------------------------------------------------------------------

func (r Request) ReferrerPolicy() string {
	return js.Value(r).Get("referrerPolicy").String()
}

// -----------------------------------------------------------------------------

func (r Request) URL() string {
	return js.Value(r).Get("url").String()
}

// -----------------------------------------------------------------------------

func (r Request) ArrayBuffer() js.Promise {
	return js.PromiseOf(js.Value(r).Call("arrayBuffer"))
}

// -----------------------------------------------------------------------------

func (r Request) Blob() js.Promise {
	return js.PromiseOf(js.Value(r).Call("blob"))
}

// -----------------------------------------------------------------------------

func (r Request) FormData() js.Promise {
	return js.PromiseOf(js.Value(r).Call("formData"))
}

// -----------------------------------------------------------------------------

func (r Request) JSON() js.Promise {
	return js.PromiseOf(js.Value(r).Call("json"))
}

// -----------------------------------------------------------------------------

func (r Request) Text() js.Promise {
	return js.PromiseOf(js.Value(r).Call("text"))
}

// -----------------------------------------------------------------------------

func (r Request) Clone() Request {
	return Request(js.Value(r).Call("clone"))
}

// -----------------------------------------------------------------------------

func NewRequest(method, url string, body any, headers ...[2]string) (Request, error) {
	opt, err := NewOption(method, body, headers...)
	if err != nil {
		return Request{}, err
	}

	return Request(builtin.Request.New(url, opt.JSValue())), nil
}

// -----------------------------------------------------------------------------

func NewRequestWithFormData(method, url string, formData form.FormData, headers ...[2]string) Request {
	opt := DefaultOpt()
	opt.Method = method
	opt.Body = formData.JSValue()
	opt.Headers.Add(headers...)

	return Request(builtin.Request.New(url, opt.JSValue()))
}
