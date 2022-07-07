//go:build js && wasm

package audio

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
	"github.com/dairaga/js/v2/media"
)

// -----------------------------------------------------------------------------

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

type Context js.Value

// -----------------------------------------------------------------------------

func (c Context) JSValue() js.Value {
	return js.Value(c)
}

// -----------------------------------------------------------------------------

func (c Context) AudioWorklet() Worklet {
	return Worklet(js.Value(c).Get("audioWorklet"))
}

// -----------------------------------------------------------------------------

func (c Context) BaseLatency() float64 {
	return js.Value(c).Get("baseLatency").Float()
}

// -----------------------------------------------------------------------------

func (c Context) OutputLatency() float64 {
	return js.Value(c).Get("outputLatency").Float()
}

// -----------------------------------------------------------------------------

func (c Context) CurrentTime() float64 {
	return js.Value(c).Get("currentTime").Float()
}

// -----------------------------------------------------------------------------

func (c Context) SampleRate() float64 {
	return js.Value(c).Get("sampleRate").Float()
}

func (c Context) State() State {
	return State(js.Value(c).Get("state").String())
}

// -----------------------------------------------------------------------------

func (c Context) Destination() DestinationNode {
	return DestinationNodeOf(js.Value(c).Get("destination"))
}

// -----------------------------------------------------------------------------

func (c Context) CreateMediaStreamSource(s media.Stream) SourceNode {
	node := js.Value(c).Call("createMediaStreamSource", s.JSValue())
	return SourceNodeOf(node)
}

// -----------------------------------------------------------------------------

func (c Context) CreateAnalyser() AnalyserNode {
	node := js.Value(c).Call("createAnalyser")
	return AnalyserNodeOf(node)
}

// -----------------------------------------------------------------------------

func (c Context) Resume() js.Promise {
	return js.PromiseOf(js.Value(c).Call("resume"))
}

// -----------------------------------------------------------------------------

func (c Context) Suspend() js.Promise {
	return js.PromiseOf(js.Value(c).Call("suspend"))
}

// -----------------------------------------------------------------------------

func (c Context) Close() js.Promise {
	return js.PromiseOf(js.Value(c).Call("close"))
}

// -----------------------------------------------------------------------------

func NewContext(sampleRate float64) Context {
	ctx := builtin.AudioContext.New(js.Obj{"sampleRate": sampleRate})
	return Context(ctx)
}

// -----------------------------------------------------------------------------

func ContextOf(v js.Value) Context {
	if !builtin.AudioContext.Is(v) {
		panic(js.ValueError{
			Method: "ContextOf",
			Type:   v.Type(),
		})
	}
	return Context(v)
}
