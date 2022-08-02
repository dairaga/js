//go:build js && wasm

package http

import (
	"syscall/js"

	"github.com/dairaga/js/v3/builtin"
)

type Credential string

func (c Credential) String() string {
	return string(c)
}

const (
	Omit       Credential = "omit"
	SameOrigin Credential = "same-origin"
	Include    Credential = "include"
)

// -----------------------------------------------------------------------------

type Headers js.Value

func (h Headers) JSValue() js.Value {
	return js.Value(h)
}

// -----------------------------------------------------------------------------

func (h Headers) Append(key, val string) Headers {
	js.Value(h).Call("append", key, val)
	return h
}

// -----------------------------------------------------------------------------

func (h Headers) Add(pairs ...[2]string) Headers {
	for i := range pairs {
		h.Append(pairs[i][0], pairs[i][1])
	}
	return h
}

// -----------------------------------------------------------------------------

func (h Headers) Delete(name string) Headers {
	js.Value(h).Call("delete", name)
	return h
}

// -----------------------------------------------------------------------------

func (h Headers) Get(name string) (string, bool) {
	result := js.Value(h).Call("get", name)
	if result.IsNull() {
		return "", false
	}
	return result.String(), true
}

// -----------------------------------------------------------------------------

func (h Headers) Has(name string) bool {
	return js.Value(h).Call("has", name).Bool()
}

// -----------------------------------------------------------------------------

func (h Headers) Set(name, val string) Headers {
	js.Value(h).Call("set", name, val)
	return h
}

// -----------------------------------------------------------------------------

func (h Headers) Foreach(fn func(string, string)) {
	it := js.Value(h).Call("entries")

	for next := it.Call("next"); !next.Get("done").Bool(); next = it.Call("next") {
		val := next.Get("value")
		fn(val.Index(0).String(), val.Index(1).String())
	}
}

// -----------------------------------------------------------------------------

func NewHeaders(init ...[2]string) Headers {
	ret := Headers(builtin.Headers.New())

	for i := range init {
		ret.Append(init[i][0], init[i][1])
	}

	return ret
}

// -----------------------------------------------------------------------------

func HeadersOf(v js.Value) Headers {
	if !builtin.Is(v, builtin.Headers) {
		panic(js.ValueError{
			Method: "HeadersOf",
			Type:   v.Type(),
		})
	}

	return Headers(v)
}
