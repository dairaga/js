//go:build js && wasm

package audio

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/media"
)

type Device interface {
	ID() string
	SampleRate() int64

	Resume() js.Promise
	Suspend() js.Promise
	Release() js.Promise
}

type device struct {
	id         string
	sampleRate int64
	interval   js.IntervalID

	ctx      Context
	src      SourceNode
	analyser AnalyserNode
}

var _ Device = &device{}

// -----------------------------------------------------------------------------

func (d *device) ID() string {
	return d.id
}

// -----------------------------------------------------------------------------

func (d *device) SampleRate() int64 {
	return d.sampleRate
}

// -----------------------------------------------------------------------------

func (d *device) Resume() js.Promise {
	return d.ctx.Resume()
}

// -----------------------------------------------------------------------------

func (d *device) Suspend() js.Promise {
	return d.ctx.Suspend()
}

// -----------------------------------------------------------------------------

func (d *device) Release() js.Promise {
	js.ClearInterval(d.interval)
	d.analyser.Disconnect()
	d.src.Disconnect()
	return d.ctx.Close()
}

// -----------------------------------------------------------------------------

func AttachDevice(id string, sampleRate int64) <-chan Device {
	ch := make(chan Device)
	media.GetUserMedia(js.Obj{"video": false, "audio": js.Obj{"deviceId": id}}).Then(func(v js.Value) any {
		d := &device{
			id:         id,
			sampleRate: sampleRate,
		}

		d.ctx = NewContext(sampleRate)
		d.src = d.ctx.CreateMediaStreamSource(media.StreamOf(v))
		d.analyser = d.ctx.CreateAnalyser()
		d.src.Connect(d.analyser)

		ch <- d
		return nil
	}).Finally(func() any {
		close(ch)
		return nil
	})

	return ch
}
