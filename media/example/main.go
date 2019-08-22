package main

import (
	"fmt"

	"github.com/dairaga/js"
	"github.com/dairaga/js/dom"
	"github.com/dairaga/js/media"
	"github.com/dairaga/js/media/recorder"
)

//var stream = media.Stream{}
var rec *recorder.Recorder
var ctx *media.AudioContext

func main() {

	dom.S("#record").OnClick(func(_ *dom.Element, _ *js.Event) {
		media.GetUserMedia(media.StreamConstrains{Audio: true, Video: false},
			func(s *media.Stream) {
				ctx = media.NewAudioContext(16000)
				rec = recorder.New(ctx.CreateMediaStreamSource(s), media.Size4096, 1)

				rec.OnRecording(func(data []float32) {
					fmt.Println(data)
				})
				fmt.Println("sample rate from audio context:", ctx.SampleRate())
				rec.Record()
			},
			func(err *js.Error) {
				fmt.Println(err)
			},
		)
	})

	dom.S("#stop").OnClick(func(_ *dom.Element, _ *js.Event) {
		rec.Stop()
		ctx.Close()
	})
	select {}
}
