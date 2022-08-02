//go:build js && wasm

package audio

import (
	"github.com/dairaga/js/v3"
	"github.com/dairaga/js/v3/builtin"
	"github.com/dairaga/js/v3/media"
)

// -----------------------------------------------------------------------------

// State presents context state.
type State string

const (
	StateSuspended State = "suspended"
	StateRunning   State = "running"
	StateClosed    State = "closed"
)

func (s State) String() string {
	return string(s)
}

// -----------------------------------------------------------------------------

// Context is Javascript AudioContext.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/AudioContext.
type Context js.Value

// -----------------------------------------------------------------------------

// JSValue returns the underlying Javascript value.
func (c Context) JSValue() js.Value {
	return js.Value(c)
}

// -----------------------------------------------------------------------------

// AudioWorklet returns the value of audioWorlet property.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/audioWorklet.
func (c Context) AudioWorklet() Worklet {
	return Worklet(js.Value(c).Get("audioWorklet"))
}

// -----------------------------------------------------------------------------

// BaseLatency returns the value of baseLatency property.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/AudioContext/baseLatency.
func (c Context) BaseLatency() float64 {
	return js.Value(c).Get("baseLatency").Float()
}

// -----------------------------------------------------------------------------

// OutputLatency returns the value of outputLatency property.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/AudioContext/outputLatency.
func (c Context) OutputLatency() float64 {
	return js.Value(c).Get("outputLatency").Float()
}

// -----------------------------------------------------------------------------

// CurrentTime returns the value of currentTime property.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/currentTime.
func (c Context) CurrentTime() float64 {
	return js.Value(c).Get("currentTime").Float()
}

// -----------------------------------------------------------------------------

// SampleRate returns the value of sampleRate property.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/sampleRate.
func (c Context) SampleRate() float64 {
	return js.Value(c).Get("sampleRate").Float()
}

func (c Context) State() State {
	return State(js.Value(c).Get("state").String())
}

// -----------------------------------------------------------------------------

// Destination returns the value of destination property.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/destination.
func (c Context) Destination() DestinationNode {
	return DestinationNodeOf(js.Value(c).Get("destination"))
}

// -----------------------------------------------------------------------------

// CreateMediaStreamSource creates a new source from given stream s.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/AudioContext/createMediaStreamSource.
func (c Context) CreateMediaStreamSource(s media.Stream) SourceNode {
	node := js.Value(c).Call("createMediaStreamSource", s.JSValue())
	return SourceNodeOf(node)
}

// -----------------------------------------------------------------------------

// CreateAnalyser returns a Javascript AnalyserNode to create data visualizations.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/createAnalyser.
func (c Context) CreateAnalyser() AnalyserNode {
	node := js.Value(c).Call("createAnalyser")
	return AnalyserNodeOf(node)
}

// -----------------------------------------------------------------------------

// Resume resumes the context.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/AudioContext/resume.
func (c Context) Resume() js.Promise {
	return js.PromiseOf(js.Value(c).Call("resume"))
}

// -----------------------------------------------------------------------------

// Suspend suspends the context.
//
// https://developer.mozilla.org/en-US/docs/Web/API/AudioContext/suspend.
func (c Context) Suspend() js.Promise {
	return js.PromiseOf(js.Value(c).Call("suspend"))
}

// -----------------------------------------------------------------------------

// Close closes the context and releases resources that it uses.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/AudioContext/close.
func (c Context) Close() js.Promise {
	return js.PromiseOf(js.Value(c).Call("close"))
}

// -----------------------------------------------------------------------------

// NewContext creates a new AudioContext with given sample rate.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/AudioContext/AudioContext.
func NewContext(sampleRate float64) Context {
	ctx := builtin.AudioContext.New(js.Obj{"sampleRate": sampleRate})
	return Context(ctx)
}

// -----------------------------------------------------------------------------

// ContextOf converts to Context from given value v.
func ContextOf(v js.Value) Context {
	if !builtin.AudioContext.Is(v) {
		panic(js.ValueError{
			Method: "ContextOf",
			Type:   v.Type(),
		})
	}
	return Context(v)
}
