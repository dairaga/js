//go:build js && wasm

package http

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/errors"
	"github.com/dairaga/js/v2/form"
	"github.com/dairaga/js/v2/promise"
)

func Do(req Request) js.Promise {
	return js.PromiseOf(js.Window().Call("fetch", req.JSValue()))
}

// -----------------------------------------------------------------------------

func fetch(method, url, contentType string, body any, headers ...[2]string) js.Promise {
	req, err := NewRequest(method, url, body)
	if err != nil {
		return promise.Reject(errors.TypeError(err).Value)
	}

	if contentType != "" {
		req.Headers().Set("Content-Type", contentType)
	}

	req.Headers().Add(headers...)
	return Do(req)
}

// -----------------------------------------------------------------------------

func Get(url string, headers ...[2]string) js.Promise {
	opt := js.ValueOf(map[string]any{
		"method":  MethodGet,
		"headers": NewHeaders(headers...).JSValue(),
	})
	return js.PromiseOf(js.Window().Call("fetch", url, opt))
}

// -----------------------------------------------------------------------------

func Head(url string, headers ...[2]string) js.Promise {
	opt := js.ValueOf(map[string]any{
		"method":  MethodHead,
		"headers": NewHeaders(headers...).JSValue(),
	})

	return js.PromiseOf(js.Window().Call("fetch", url, opt))
}

// -----------------------------------------------------------------------------

func Post(url, contentType string, body any, headers ...[2]string) js.Promise {
	return fetch(MethodPost, url, contentType, body, headers...)
}

// -----------------------------------------------------------------------------

func PostJSON(url string, data any, headers ...[2]string) js.Promise {
	return Post(url, "application/json; charset=utf-8", data, headers...)
}

// -----------------------------------------------------------------------------

func PostForm(url string, formData form.FormData, headers ...[2]string) js.Promise {
	return Do(NewRequestWithFormData(MethodPost, url, formData, headers...))
}

// -----------------------------------------------------------------------------

func JSONCall(method, url string, data any, headers ...[2]string) js.Promise {
	return fetch(method, url, "application/json; charset=utf-8", data, headers...)
}
