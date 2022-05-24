//go:build js && wasm

package upload

import "github.com/dairaga/js/v2"

var (
	defaultWithCredentials = true
)

// -----------------------------------------------------------------------------

type uploader struct {
	ref      js.Value
	listener js.Listener
	lastErr  error
}

// -----------------------------------------------------------------------------

type Client struct {
	ref      js.Value
	lastErr  error
	released bool
	js.Listener
}

// -----------------------------------------------------------------------------
