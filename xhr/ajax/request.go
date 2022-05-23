//go:build js && wasm

package ajax

import (
	"encoding/json"
)

type Request struct {
	User     string
	Password string

	method  string
	url     string
	mime    string
	headers map[string]string
	body    []byte
}

// -----------------------------------------------------------------------------

func (req *Request) Method() string {
	return req.method
}

// -----------------------------------------------------------------------------

func (req *Request) URL() string {
	return req.url
}

// -----------------------------------------------------------------------------

func (req *Request) MimeType() string {
	return req.mime
}

// -----------------------------------------------------------------------------

func (req *Request) Header(key string) string {
	return req.headers[key]
}

// -----------------------------------------------------------------------------

func (req *Request) Body() []byte {
	return req.body
}

// -----------------------------------------------------------------------------

// SetHeader set request header.
func (req *Request) SetHeader(key, value string) *Request {

	if req.headers == nil {
		req.headers = make(map[string]string)
	}

	req.headers[key] = value
	return req
}

// -----------------------------------------------------------------------------

func guess(x any) (mime string, data []byte, err error) {
	if x != nil {
		switch v := x.(type) {
		case string:
			mime = MimeText
			data = []byte(v)
		case []byte:
			mime = MimeStream
			data = v
		default:
			data, err = json.Marshal(x)
			if err == nil {
				mime = MimeJSON
			}
		}
	}

	return
}

// -----------------------------------------------------------------------------

// NewRequest returns a request.
func NewRequest(method, url string, x any) (req *Request, err error) {
	mime, data, err := guess(x)
	if err != nil {
		return nil, err
	}

	req = new(Request)
	req.method = method
	req.url = url
	req.mime = mime
	req.body = data
	if req.mime != "" {
		req.SetHeader("Content-Type", req.mime)
	}

	return
}
