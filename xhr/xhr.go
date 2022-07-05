//go:build js && wasm

// Package xhr is a wrapper of javascript XMLHttpRequest (XHR), and provides a simple way to make a JSON AJAX.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest.
package xhr

import (
	"fmt"
)

type Error int

const (
	ErrAbort    Error = -1 // Request aborted.
	ErrFailed   Error = -2 // Request failed.
	ErrTimeout  Error = -3 // Request timeout.
	ErrReleased Error = -4 // Request released.

	GET    = "GET"    // HTTP GET method.
	POST   = "POST"   // HTTP POST method.
	PUT    = "PUT"    // HTTP PUT method.
	DELETE = "DELETE" // HTTP DELETE method.
	PATCH  = "PATCH"  // HTTP PATCH method.
)

// -----------------------------------------------------------------------------

func (e Error) Error() string {
	switch e {
	case ErrAbort:
		return "abort"
	case ErrFailed:
		return "error"
	case ErrTimeout:
		return "timeout"
	case ErrReleased:
		return "released"
	default:
		if e >= StatusBadRequest {
			return fmt.Sprintf("code: %d, status: %s", e, StatusText(int(e)))
		}
		return fmt.Sprintf("unknown error (%d)", e)
	}
}
