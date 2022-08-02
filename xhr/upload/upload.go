//go:build js && wasm

package upload

import (
	"fmt"

	"github.com/dairaga/js/v3"
	"github.com/dairaga/js/v3/builtin"
	"github.com/dairaga/js/v3/form"
	"github.com/dairaga/js/v3/xhr"
)

type ProgressState int

const (
	StateStart    ProgressState = 1
	StateProgress ProgressState = 2
	StateFinish   ProgressState = 3
)

func (s ProgressState) String() string {
	switch s {
	case StateStart:
		return "started"
	case StateProgress:
		return "progressing"
	case StateFinish:
		return "finished"
	default:
		return fmt.Sprintf("unknown type (%d)", s)
	}
}

// -----------------------------------------------------------------------------

var (
	defaultWithCredentials = true
)

// -----------------------------------------------------------------------------

type upload struct {
	ref      js.Value
	listener js.Listener
}

// -----------------------------------------------------------------------------

func newUpload(ref js.Value) *upload {
	return &upload{
		ref:      ref,
		listener: make(js.Listener),
	}
}

// -----------------------------------------------------------------------------

type Client struct {
	ref      js.Value
	lastErr  error
	released bool
	listener js.Listener

	upload *upload
	jobID  string
}

// -----------------------------------------------------------------------------

func (cli *Client) Upload(url, jobID string, form form.FormData, headers ...[2]string) error {
	if cli.released {
		return xhr.ErrReleased
	}
	cli.jobID = jobID
	cli.ref.Call("open", xhr.POST, url, true)
	for _, header := range headers {
		cli.ref.Call("setRequestHeader", header[0], header[1])
	}
	cli.ref.Call("send", form.JSValue())
	return nil
}

// -----------------------------------------------------------------------------

func (cli *Client) Abort() {
	cli.ref.Call("abort")
}

// -----------------------------------------------------------------------------

func (cli *Client) WithCredentials(flag bool) {
	cli.ref.Set("withCredentials", flag)
}

// -----------------------------------------------------------------------------

func (cli *Client) Release() {
	if cli.released {
		return
	}
	cli.released = true
	cli.upload.listener.Release()
	cli.listener.Release()
}

// -----------------------------------------------------------------------------

func progressEvt(v js.Value) (computable bool, loaded uint64, total uint64) {
	computable = v.Get("lengthComputable").Bool()
	loaded = uint64(v.Get("loaded").Int())
	total = uint64(v.Get("total").Int())
	if !computable {
		total = 0
	}

	return
}

// -----------------------------------------------------------------------------

func New(h Handler) *Client {
	cli := new(Client)
	cli.ref = builtin.XMLHttpRequest.New()
	cli.ref.Set("withCredentials", defaultWithCredentials)

	cli.upload = newUpload(cli.ref.Get("upload"))
	cli.listener = make(js.Listener)
	cli.released = false
	cli.ref.Set("responseType", "arraybuffer")

	cli.upload.listener.Add(cli.upload.ref, "loadstart", func(_ js.Value, args []js.Value) any {
		_, loaded, total := progressEvt(args[0])
		h.Progress(cli.jobID, StateStart, loaded, total)
		return nil
	})

	cli.upload.listener.Add(cli.upload.ref, "progress", func(_ js.Value, args []js.Value) any {
		_, loaded, total := progressEvt(args[0])
		h.Progress(cli.jobID, StateProgress, loaded, total)
		return nil
	})

	cli.upload.listener.Add(cli.upload.ref, "load", func(_ js.Value, args []js.Value) any {
		_, loaded, total := progressEvt(args[0])
		h.Progress(cli.jobID, StateFinish, loaded, total)
		return nil
	})

	cli.listener.Add(cli.ref, "abort", func(_ js.Value, args []js.Value) any {
		cli.lastErr = xhr.ErrAbort
		return nil
	})

	cli.listener.Add(cli.ref, "error", func(_ js.Value, args []js.Value) any {
		cli.lastErr = xhr.ErrFailed
		return nil
	})

	cli.listener.Add(cli.ref, "timeout", func(_ js.Value, args []js.Value) any {
		cli.lastErr = xhr.ErrTimeout
		return nil
	})

	cli.listener.Add(cli.ref, "loadend", func(_ js.Value, args []js.Value) any {
		var resp *xhr.Response = nil
		if cli.lastErr == nil {
			resp = xhr.ResponseOf(cli.ref)
		}
		h.Completed(resp, cli.lastErr)
		return nil
	})

	return cli
}
