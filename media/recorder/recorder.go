package recorder

import (
	"syscall/js"

	"github.com/dairaga/js/media"
)

// ----------------------------------------------------------------------------

func resample(buf []float32, srcRate, destRate float64) []float32 {
	return nil
}

// ----------------------------------------------------------------------------

// Recorder ...
type Recorder struct {
	ctx      media.AudioContext
	source   media.AudioNode
	size     media.BufSize
	channels int

	recording bool
	cb        func([]float32)
	processor media.AudioNode
}

// Recording ...
func (r Recorder) Recording() bool {
	return r.recording
}

// New ...
func New(source media.AudioNode, size media.BufSize, channels int) *Recorder {

	if size <= 0 {
		size = media.Size4096
	}

	if channels <= 0 {
		channels = 1
	}

	r := &Recorder{
		ctx:       source.Context(),
		source:    source,
		size:      size,
		channels:  channels,
		recording: false,
	}

	r.processor = r.ctx.CreateScriptProcessor(r.size, r.channels, r.channels)

	r.processor.JSValue().Set("onaudioprocess", js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		if !r.recording {
			return nil
		}

		if r.cb != nil {
			evt := args[0]
			input := media.AudioBufferOf(evt.Get("inputBuffer"))
			ch := input.NumberOfChannels()

			buf := make([][]float32, ch)

			for i := 0; i < ch; i++ {
				buf[i] = input.ChannelData(i)
			}

			if ch == 2 {
				newLen := len(buf[0]) + len(buf[1])
				newbuf := make([]float32, newLen)
				idx := 0

				for i := 0; i < newLen; i++ {
					newbuf[i] = buf[0][idx]
					i++
					newbuf[i] = buf[1][idx]
					idx++
				}
				r.cb(newbuf)
			} else {
				r.cb(buf[0])
			}

		}

		return nil
	}))

	source.Connect(r.processor)
	r.processor.Connect(r.ctx.Destination())

	return r
}

// OnRecording ...
func (r *Recorder) OnRecording(cb func([]float32)) {
	r.cb = cb
}

// Record ...
func (r *Recorder) Record() {
	r.recording = true
}

// Stop ...
func (r *Recorder) Stop() {
	r.recording = false
}

// Release ...
func (r *Recorder) Release() {
	r.Stop()
	r.processor.DisconnectAll()
	r.source.Disconnect(r.processor)

}
