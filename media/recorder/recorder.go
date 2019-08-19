package recorder

import (
	"github.com/dairaga/js"
	"github.com/dairaga/js/media"
)

// ----------------------------------------------------------------------------

// Recorder is a recorder component.
type Recorder struct {
	ctx      media.AudioContext
	source   media.AudioNode
	size     media.BufSize
	channels int

	recording bool
	cb        func([]float32)
	processor media.AudioNode
	onProcess js.Func
}

// Recording returns true when recording.
func (r Recorder) Recording() bool {
	return r.recording
}

// New returns a recorder component.
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
	r.onProcess = js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
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
				// interleave
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
	})

	r.processor.JSValue().Set("onaudioprocess", r.onProcess)

	source.Connect(r.processor)
	r.processor.Connect(r.ctx.Destination())

	return r
}

// OnRecording callback invokeing when recording.
func (r *Recorder) OnRecording(cb func([]float32)) {
	r.cb = cb
}

// Record starts to record.
func (r *Recorder) Record() {
	r.recording = true
}

// Stop stops to record.
func (r *Recorder) Stop() {
	r.recording = false
}

// Release frees up resources. Recorder must be not used after calling Release.
func (r *Recorder) Release() {
	r.Stop()
	r.processor.DisconnectAll()
	r.source.Disconnect(r.processor)
	r.onProcess.Release()
}
