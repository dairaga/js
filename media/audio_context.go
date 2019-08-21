package media

import (
	"fmt"

	"github.com/dairaga/js"
)

// BufSize buffer size for script processor.
type BufSize int

// buffer https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/createScriptProcessor#Parameters
const (
	Size256   BufSize = 256
	Size512   BufSize = 512
	Size1024  BufSize = 1024
	Size2048  BufSize = 2048
	Size4096  BufSize = 4096
	Size8192  BufSize = 8192
	Size16384 BufSize = 16384
)

// AudioState state for AudioConext.
type AudioState string

// AudioContext State https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/state
const (
	Suspended AudioState = "suspended"
	Running   AudioState = "running"
	Closed    AudioState = "closed"
)

func toFloat32Slice(srcFloat32Array js.Value) []float32 {
	size := srcFloat32Array.Get("length")
	if !size.Truthy() {
		return nil
	}

	ret := make([]float32, size.Int())
	destArray := js.TypedArrayOf(ret)

	destArray.Call("set", srcFloat32Array, 0)

	destArray.Release()

	return ret
}

// ----------------------------------------------------------------------------

// AudioContext https://developer.mozilla.org/en-US/docs/Web/API/AudioContext
type AudioContext struct {
	ref       js.Value
	onClose   js.Func
	onSuspend js.Func
	onResume  js.Func
}

var audioContextConstructor = js.Value{}

func init() {

	if c := window.Get("AudioContext"); c.Truthy() {
		audioContextConstructor = c
	} else if c := window.Get("webkitAudioContext"); c.Truthy() {
		audioContextConstructor = c
	} else if c := window.Get("oAudioContext"); c.Truthy() {
		audioContextConstructor = c
	} else if c := window.Get("msAudioContext"); c.Truthy() {
		audioContextConstructor = c
	} else {
		panic("AudioContext not supported")
	}
}

// NewAudioContext return a audio context with specific sample rate.
func NewAudioContext(sampleRate float64) *AudioContext {
	if sampleRate > 0 {
		option := map[string]interface{}{
			"sampleRate": sampleRate,
		}
		return &AudioContext{ref: audioContextConstructor.New(option)}
	}
	return &AudioContext{ref: audioContextConstructor.New()}
}

// JSValue ...
func (ctx *AudioContext) JSValue() js.Value {
	return ctx.ref
}

// CreateMediaStreamSource https://developer.mozilla.org/en-US/docs/Web/API/AudioContext/createMediaStreamSource
func (ctx *AudioContext) CreateMediaStreamSource(stream *Stream) *AudioNode {
	return AudioNodeOf(ctx.ref.Call("createMediaStreamSource", stream.ref))
}

// CreateScriptProcessor https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/createScriptProcessor
func (ctx *AudioContext) CreateScriptProcessor(size BufSize, in, out int) *AudioNode {
	return AudioNodeOf(ctx.ref.Call("createScriptProcessor", int(size), in, out))
}

// Destination https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/destination
func (ctx *AudioContext) Destination() *AudioNode {
	return AudioNodeOf(ctx.ref.Get("destination"))
}

// State https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/state
func (ctx *AudioContext) State() AudioState {
	return AudioState(ctx.ref.Get("state").String())
}

// SampleRate https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/sampleRate
func (ctx *AudioContext) SampleRate() float64 {
	return ctx.ref.Get("sampleRate").Float()
}

// Close https://developer.mozilla.org/en-US/docs/Web/API/AudioContext/close
func (ctx *AudioContext) Close() *AudioContext {
	promise := ctx.ref.Call("close")
	if ctx.onClose.Truthy() {
		promise.Call("then", ctx.onClose)
	}
	return ctx
}

// Suspend https://developer.mozilla.org/en-US/docs/Web/API/AudioContext/suspend
func (ctx *AudioContext) Suspend() *AudioContext {
	promise := ctx.ref.Call("suspend")
	if ctx.onSuspend.Truthy() {
		promise.Call("then", ctx.onSuspend)
	}
	return ctx
}

// Resume https://developer.mozilla.org/en-US/docs/Web/API/AudioContext/resume
func (ctx *AudioContext) Resume() *AudioContext {
	promise := ctx.ref.Call("resume")
	if ctx.onResume.Truthy() {
		promise.Call("then", ctx.onResume)
	}
	return ctx
}

// Release frees up callback functions.
func (ctx *AudioContext) Release() {
	if ctx.State() != Closed {
		return
	}

	if ctx.onClose.Truthy() {
		ctx.onClose.Release()
	}

	if ctx.onSuspend.Truthy() {
		ctx.onSuspend.Release()
	}

	if ctx.onResume.Truthy() {
		ctx.onResume.Release()
	}
}

// OnClose invoked when context is closed.
func (ctx *AudioContext) OnClose(cb func()) *AudioContext {
	ctx.onClose = js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		cb()
		return nil
	})

	return ctx
}

// OnSuspend invoked when context is suspended.
func (ctx *AudioContext) OnSuspend(cb func()) *AudioContext {
	ctx.onSuspend = js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		cb()
		return nil
	})
	return ctx
}

// OnResume invoked when context is resumed.
func (ctx *AudioContext) OnResume(cb func()) *AudioContext {
	ctx.onResume = js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		cb()
		return nil
	})
	return ctx
}

// ----------------------------------------------------------------------------

// AudioNode https://developer.mozilla.org/en-US/docs/Web/API/AudioNode
type AudioNode struct {
	ref js.Value
}

// AudioNodeOf returns a audio node.
func AudioNodeOf(x js.Value) *AudioNode {
	return &AudioNode{ref: x}
}

// JSValue ...
func (n *AudioNode) JSValue() js.Value {
	return n.ref
}

// Context https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/context
func (n *AudioNode) Context() *AudioContext {
	return &AudioContext{ref: n.ref.Get("context")}
}

// Connect https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/connect
func (n *AudioNode) Connect(dest *AudioNode, index ...int) *AudioNode {
	size := len(index)
	switch size {
	case 0:
		n.ref.Call("connect", dest)
	case 1:
		n.ref.Call("connect", dest, index[0])
	case 2:
		n.ref.Call("connect", dest, index[0], index[1])
	}
	return n
}

// DisconnectAll disconnects all destination nodes.
func (n *AudioNode) DisconnectAll() {
	n.ref.Call("disconnect")
}

// Disconnect https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/disconnect
func (n *AudioNode) Disconnect(dest *AudioNode, index ...int) {
	size := len(index)
	switch size {
	case 0:
		n.ref.Call("disconnect", dest)
	case 1:
		n.ref.Call("disconnect", dest, index[0])
	case 2:
		n.ref.Call("disconnect", dest, index[0], index[1])
	}
}

// ----------------------------------------------------------------------------

// AudioBuffer https://developer.mozilla.org/en-US/docs/Web/API/AudioBuffer
type AudioBuffer struct {
	ref js.Value
}

// AudioBufferOf returns audio buffer.
func AudioBufferOf(x js.Value) *AudioBuffer {
	return &AudioBuffer{ref: x}
}

// JSValue ...
func (buf *AudioBuffer) JSValue() js.Value {
	return buf.ref
}

// SampleRate https://developer.mozilla.org/en-US/docs/Web/API/AudioBuffer/sampleRate
func (buf *AudioBuffer) SampleRate() float64 {
	return buf.ref.Get("sampleRate").Float()
}

// Length https://developer.mozilla.org/en-US/docs/Web/API/AudioBuffer/length
func (buf *AudioBuffer) Length() int {
	return buf.ref.Get("length").Int()
}

// Duration https://developer.mozilla.org/en-US/docs/Web/API/AudioBuffer/duration
func (buf *AudioBuffer) Duration() float64 {
	return buf.ref.Get("duration").Float()
}

// NumberOfChannels https://developer.mozilla.org/en-US/docs/Web/API/AudioBuffer/numberOfChannels
func (buf *AudioBuffer) NumberOfChannels() int {
	return buf.ref.Get("numberOfChannels").Int()
}

func (buf *AudioBuffer) String() string {
	return fmt.Sprintf(`{"sample_rate": %v, "length": %d, "duration": %v, "number_of_channels": %v}`, buf.SampleRate(), buf.Length(), buf.Duration(), buf.NumberOfChannels())
}

// ChannelData https://developer.mozilla.org/en-US/docs/Web/API/AudioBuffer/getChannelData
func (buf *AudioBuffer) ChannelData(idx int) []float32 {
	if idx >= buf.NumberOfChannels() {
		return nil
	}
	return toFloat32Slice(buf.ref.Call("getChannelData", idx))
}
