//go:build js && wasm

package audio

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/media"
)

// Device represents an AudioDevice. It is not a built-in Javascript object.
// Device is to attach an device id and process data with a custom worklet.
type Device interface {
	js.Wrapper

	ID() string
	SampleRate() float64
	GetByteFrequencyData() []byte
	Process(url, name string, cb func(js.Event), c ...js.Credential) js.Promise

	Truthy() bool
	Ready() bool
	Closed() bool

	Close() js.Promise
}

type device js.Value

// -----------------------------------------------------------------------------

// JSvalue returns the underlying Javascript value.
func (d device) JSValue() js.Value {
	return js.Value(d)
}

// -----------------------------------------------------------------------------

// ID returns the device id. Device id can be retrieved from the MediaDevices.enumerateDevices() method.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/MediaDevices/enumerateDevices.
func (d device) ID() string {
	return js.Value(d).Get("deviceID").String()
}

// -----------------------------------------------------------------------------

// SampleRate returns the sample rate inputted when attaching the device.
func (d device) SampleRate() float64 {
	return js.Value(d).Get("sampleRate").Float()
}

// -----------------------------------------------------------------------------

// GetByteFrequencyData returns current frequency data from underlying AnalyzerNode.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/getByteFrequencyData.
func (d device) GetByteFrequencyData() []byte {
	n := analyserNode{
		node: node(js.Value(d).Get("analyserNode")),
	}

	return n.GetByteFrequencyData()
}

// -----------------------------------------------------------------------------

// Process process attached pyhisical device.
// Give url and name to create a worklet to load module.
// Give callback cb to handle data from worklet.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/Worklet/addModule.
// See https://developer.mozilla.org/en-US/docs/Web/API/AudioWorkletProcessor/AudioWorkletProcessor.
// See https://developer.mozilla.org/en-US/docs/Web/API/MessagePort/message_event.
func (d device) Process(url, name string, cb func(js.Event), c ...js.Credential) js.Promise {
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
			worker := NewWorkletNode(ctx, name)
			src.Connect(worker)
			worker.Connect(ctx.Destination())

			cb := worker.Port().OnMessage(func(evt js.Event) {
				cb(evt)
			})

			worker.Port().Start() // must call start after addEventListener, see https://developer.mozilla.org/en-US/docs/Web/API/MessagePort/start.
			js.Value(d).Set("workletNode", worker.JSValue())
			js.Value(d).Set("ready", true)
			js.Value(d).Set("msgcb", cb.Value)
			return nil
		})
		return nil
	})
}

// -----------------------------------------------------------------------------

// Ready returns true if it is ready to process.
func (d device) Ready() bool {
	ready := js.Value(d).Get("ready")
	return ready.Truthy() && ready.Bool()
}

// -----------------------------------------------------------------------------

// Closed returns true if the device is closed.
func (d device) Closed() bool {
	ctx := js.Value(d).Get("context")
	return !ctx.Truthy() || Context(ctx).State() == StateClosed
}

// -----------------------------------------------------------------------------

// Close detaches the device and releases resources used.
func (d device) Close() js.Promise {
	dv := js.Value(d)
	dv.Set("ready", false)

	js.Func{Value: dv.Get("msgcb")}.Release()

	dv.Get("workletNode").Get("port").Call("close")
	dv.Get("workletNode").Call("disconnect")
	dv.Get("analyserNode").Call("disconnect")
	dv.Get("sourceNode").Call("disconnect")

	return Context(dv.Get("context")).Close()
}

// -----------------------------------------------------------------------------

// Truthy returns true if it is valid.
func (d device) Truthy() bool {
	return js.Value(d).Get("deviceID").Truthy() && js.Value(d).Get("sampleRate").Truthy()
}

// -----------------------------------------------------------------------------

// NewDevice attachs to pyhisical device with given id.
// The device id can be retrieved from the MediaDevices.enumerateDevices() method.
// Give sample rate to process the device.
func NewDevice(id string, sampleRate float64) Device {
	return device(js.ValueOf(map[string]any{
		"deviceID":   id,
		"sampleRate": sampleRate,
	}))
}
