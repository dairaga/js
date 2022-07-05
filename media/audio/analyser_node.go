//go:build js && wasm

package audio

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

type AnalyserNode interface {
	Node

	FFTSize() uint
	FrequencyBinCount() uint
	MaxDecibels() float64
	MinDecibels() float64
	SmoothingTimeConstant() float64

	GetByteFrequencyData() []byte
	GetByteTimeDomainData() []byte
	GetFloatFrequencyData() []float32
	getFloatTimeDomainData() []float32
}

type analyserNode struct {
	node
}

var _ AnalyserNode = &analyserNode{}

// -----------------------------------------------------------------------------

func (a *analyserNode) FFTSize() uint {
	return uint(a.JSValue().Get("fftSize").Int())
}

// -----------------------------------------------------------------------------

func (a *analyserNode) FrequencyBinCount() uint {
	return uint(a.JSValue().Get("frequencyBinCount").Int())
}

// -----------------------------------------------------------------------------

func (a *analyserNode) MaxDecibels() float64 {
	return a.JSValue().Get("maxDecibels").Float()
}

// -----------------------------------------------------------------------------

func (a *analyserNode) MinDecibels() float64 {
	return a.JSValue().Get("minDecibels").Float()
}

// -----------------------------------------------------------------------------

func (a *analyserNode) SmoothingTimeConstant() float64 {
	return a.JSValue().Get("smoothingTimeConstant").Float()
}

// -----------------------------------------------------------------------------

func (a *analyserNode) GetByteFrequencyData() []byte {
	size := a.FrequencyBinCount()
	buf := builtin.Uint8Array.New(size)
	a.JSValue().Call("getByteFrequencyData", buf)
	return js.GoBytes(buf)
}

// -----------------------------------------------------------------------------

func (a *analyserNode) GetByteTimeDomainData() []byte {
	size := a.FFTSize()
	buf := builtin.Uint8Array.New(size)
	a.JSValue().Call("getByteTimeDomainData", buf)
	return js.GoBytes(buf)
}

// -----------------------------------------------------------------------------

func (a *analyserNode) GetFloatFrequencyData() []float32 {
	size := a.FrequencyBinCount()
	buf := builtin.Float32Array.New(size)
	a.JSValue().Call("getFloatFrequencyData", buf)
	return js.Float32Array(buf)
}

// -----------------------------------------------------------------------------

func (a *analyserNode) getFloatTimeDomainData() []float32 {
	size := a.FFTSize()
	buf := builtin.Float32Array.New(size)
	a.JSValue().Call("getFloatTimeDomainData", buf)
	return js.Float32Array(buf)
}

// -----------------------------------------------------------------------------

func AnalyserNodeOf(v js.Value) AnalyserNode {
	if !builtin.AnalyserNode.Is(v) {
		panic(js.ValueError{
			Method: "AnalyserNodeOf",
			Type:   v.Type(),
		})
	}
	return &analyserNode{node: node(v)}
}
