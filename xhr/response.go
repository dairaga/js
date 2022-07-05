//go:build js && wasm

package xhr

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dairaga/js/v2"
)

// HandlerFunc is a function to handle XHR response.
type HandlerFunc = func(*Response, error)

// -----------------------------------------------------------------------------

// Respopnse is a XHR response to a XHR request.
// Response imitate golang http Response.
type Response struct {
	code    int               // http status code.
	headers map[string]string // response headers.
	body    []byte            // response body as bytes default.
}

func (r *Response) String() string {
	return fmt.Sprintf("%d %s", r.StatusCode(), r.StatusText())
}

// -----------------------------------------------------------------------------

// StatusCode returns the response http status code.
func (r *Response) StatusCode() int {
	return r.code
}

// -----------------------------------------------------------------------------

// StatusText returns the response status text string.
func (r *Response) StatusText() string {
	return StatusText(r.code)
}

// -----------------------------------------------------------------------------

// Body returns response body.
func (r *Response) Body() []byte {
	return r.body
}

// -----------------------------------------------------------------------------

// Header returns the value of the named header.
func (r *Response) Header(key string) string {
	key = strings.ToLower(key)
	return r.headers[key]
}

// -----------------------------------------------------------------------------

// Headers returns all headers.
func (r *Response) Headers() map[string]string {
	return r.headers
}

// -----------------------------------------------------------------------------

// OK returns true if the response is successful (code is between 200 and 299).
func (r *Response) OK() bool {
	return r.code >= StatusOK && r.code <= StatusIMUsed
}

// -----------------------------------------------------------------------------

// Unmarshal decodes the response body to the given value x.
func (r *Response) Unmarshal(x any) error {
	return json.Unmarshal(r.body, x)
}

// -----------------------------------------------------------------------------

// ResponseOf returns a Response from the given XHR request.
func ResponseOf(xhr js.Value) *Response {
	resp := new(Response)
	resp.code = xhr.Get("status").Int()

	resp.body = js.ArrayBufferToBytes(xhr.Get("response"))
	headers := strings.TrimSpace(xhr.Call("getAllResponseHeaders").String())
	tmp := strings.Split(headers, "\r\n")
	resp.headers = make(map[string]string)
	for _, x := range tmp {
		x := strings.TrimSpace(x)
		if strings.Contains(x, ":") {
			parts := strings.Split(x, ": ")
			resp.headers[parts[0]] = strings.Join(parts[1:], ": ")
		}
	}

	return resp
}
