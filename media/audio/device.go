//go:build js && wasm

package audio

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/media"
)

type Device js.Value

// -----------------------------------------------------------------------------

func (d Device) JSValue() js.Value {
	return js.Value(d)
}

// -----------------------------------------------------------------------------

func (d Device) ID() string {
	return js.Value(d).Get("deviceID").String()
}

// -----------------------------------------------------------------------------

func (d Device) SampleRate() float64 {
	return js.Value(d).Get("sampleRate").Float()
}

// -----------------------------------------------------------------------------

func (d Device) GetByteFrequencyData() []byte {
	n := analyserNode{
		node: node(js.Value(d).Get("analyserNode")),
	}

	return n.GetByteFrequencyData()
}

// -----------------------------------------------------------------------------

func (d Device) Record(url, name string, cb func(js.Value), c ...js.Credential) js.Promise {
	return media.GetUserMedia(js.Obj{
		"video": false,
		"audio": d.ID(),
	}).Then(func(stream js.Value) any {
		ctx := NewContext(d.SampleRate())
		src := ctx.CreateMediaStreamSource(media.StreamOf(stream))
		analyser := ctx.CreateAnalyser()
		src.Connect(analyser)

		js.Value(d).Set("context", ctx.JSValue())
		js.Value(d).Set("sourceNode", src.JSValue())
		js.Value(d).Set("analyserNode", analyser.JSValue())

		ctx.AudioWorklet().AddModule(url, c...).Then(func(js.Value) any {
			worker := NewWorklet(ctx, name)
			src.Connect(worker)
			worker.Connect(ctx.Destination())

			worker.Port().OnMessage(func(value js.Value) {
				cb(value)
			})

			worker.Port().Start()
			js.Value(d).Set("workletNode", worker.JSValue())
			js.Value(d).Set("ready", true)
			return nil
		})
		return nil
	})
}

// -----------------------------------------------------------------------------

func (d Device) Ready() bool {
	ready := js.Value(d).Get("ready")
	return ready.Truthy() && ready.Bool()
}

// -----------------------------------------------------------------------------

func (d Device) Closed() bool {
	ctx := js.Value(d).Get("context")
	return !ctx.Truthy() || Context(ctx).State() == StateClosed
}

// -----------------------------------------------------------------------------

func (d Device) Close() js.Promise {
	dv := js.Value(d)

	dv.Get("workletNode").Call("disconnect")
	dv.Get("analyserNode").Call("disconnect")
	dv.Get("sourceNode").Call("disconnect")
	return Context(dv.Get("context")).Close()
}

// -----------------------------------------------------------------------------

func (d Device) Truthy() bool {
	return js.Value(d).Get("deviceID").Truthy() && js.Value(d).Get("sampleRate").Truthy()
}

// -----------------------------------------------------------------------------

func NewDevice(id string, sampleRate float64) Device {
	return Device(js.ValueOf(map[string]any{
		"deviceID":   id,
		"sampleRate": sampleRate,
	}))
}
