//go:build js && wasm

package xhr

import (
	"fmt"
)

type Error int

const (
	ErrAbort    Error = -1
	ErrFailed   Error = -2
	ErrTimeout  Error = -3
	ErrReleased Error = -4
)

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
