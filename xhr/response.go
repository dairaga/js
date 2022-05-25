//go:build js && wasm

package xhr

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dairaga/js/v2"
)

type HandleFunc = func(*Response, error)

// -----------------------------------------------------------------------------

type Response struct {
	code    int
	headers map[string]string
	body    []byte
}

func (r *Response) String() string {
	return fmt.Sprintf("%d %s", r.StatusCode(), r.StatusText())
}

// -----------------------------------------------------------------------------

func (r *Response) StatusCode() int {
	return r.code
}

// -----------------------------------------------------------------------------

func (r *Response) StatusText() string {
	return StatusText(r.code)
}

// -----------------------------------------------------------------------------

func (r *Response) Body() []byte {
	return r.body
}

// -----------------------------------------------------------------------------

func (r *Response) Header(key string) string {
	key = strings.ToLower(key)
	return r.headers[key]
}

// -----------------------------------------------------------------------------

func (r *Response) OK() bool {
	return r.code >= StatusOK && r.code <= StatusIMUsed
}

// -----------------------------------------------------------------------------

func (r *Response) Unmarshal(x any) error {
	return json.Unmarshal(r.body, x)
}

// -----------------------------------------------------------------------------

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