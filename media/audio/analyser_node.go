//go:build js && wasm

package audio

import (
	"github.com/dairaga/js/v3"
	"github.com/dairaga/js/v3/builtin"
)

// AnalyserNode is Javascript AnalyserNode.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode.
type AnalyserNode interface {
	Node

	// FFTSize returns the value of the fftSize property.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/fftSize.
	FFTSize() uint

	// FrequencyBinCount returns the value of the frequencyBinCount property.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/frequencyBinCount.
	FrequencyBinCount() uint

	// MaxDecibels returns the value of the maxDecibels property.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/maxDecibels.
	MaxDecibels() float64

	// MinDecibels returns the value of the minDecibels property.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/minDecibels.
	MinDecibels() float64

	// SmoothingTimeConstant returns the value of the smoothingTimeConstant property.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/smoothingTimeConstant.
	SmoothingTimeConstant() float64

	// GetByteFrequencyData returns the current frequency data into bytes.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/getByteFrequencyData.
	GetByteFrequencyData() []byte

	// GetByteTimeDomainData returns he current waveform, or time-domain, data into bytes.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/getByteTimeDomainData.
	GetByteTimeDomainData() []byte

	// GetFloatFrequencyData returns the current frequency data into float32 array.
	//
	// https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/getFloatFrequencyData.
	GetFloatFrequencyData() []float32

	// GetFloatTimeDomainData returns the current waveform, or time-domain, data into float32 array.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/AnalyserNode/getFloatTimeDomainData.
	GetFloatTimeDomainData() []float32
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

func (a *analyserNode) GetFloatTimeDomainData() []float32 {
	size := a.FFTSize()
	buf := builtin.Float32Array.New(size)
	a.JSValue().Call("getFloatTimeDomainData", buf)
	return js.Float32Array(buf)
}

// -----------------------------------------------------------------------------

// AnalyserNodeOf converts to AnalyserNode from given Javascript value v.
func AnalyserNodeOf(v js.Value) AnalyserNode {
	if !builtin.AnalyserNode.Is(v) {
		panic(js.ValueError{
			Method: "AnalyserNodeOf",
			Type:   v.Type(),
		})
	}
	return &analyserNode{node: node(v)}
}
