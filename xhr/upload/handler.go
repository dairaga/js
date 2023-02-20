//go:build js && wasm

package upload

import "github.com/dairaga/js/v2/xhr"

type ProgressFunc func(string, ProgressState, uint64, uint64)

// -----------------------------------------------------------------------------

type Handler interface {
	Progress(jobID string, state ProgressState, loaded, total uint64)
	Completed(resp *xhr.Response, err error)
}
