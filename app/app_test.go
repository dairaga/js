//go:build js && wasm

package app

import (
	"testing"
	"time"

	"github.com/dairaga/js/v2/url"
	"github.com/stretchr/testify/assert"
)

type serv struct {
	t   *testing.T
	ch  chan struct{}
	cur string
	old string
}

func (s *serv) Serve(oldURL, newURL url.URL) {
	s.old = oldURL.Hash()
	s.cur = newURL.Hash()
	s.ch <- struct{}{}
}

func TestHash(t *testing.T) {
	serv := &serv{
		t:  t,
		ch: make(chan struct{}, 1),
	}

	//curURL := url.New(js.Global().Get("location").Get("href").String())
	Init(serv)
	_app.handler.Serve(_app.currentURL, _app.currentURL)
	<-serv.ch
	assert.Equal(t, "", serv.old)
	assert.Equal(t, "", serv.cur)

	ChangeHash("#a100")
	<-serv.ch
	assert.Equal(t, "", serv.old)
	assert.Equal(t, "#a100", serv.cur)

	ChangeHash("#b100")
	<-serv.ch
	assert.Equal(t, "#a100", serv.old)
	assert.Equal(t, "#b100", serv.cur)

	ChangeHash("#b100")
	select {
	case <-time.After(1 * time.Second):
	case <-serv.ch:
	}
	// the hashchange event is not triggered.
	assert.Equal(t, "#a100", serv.old)
	assert.Equal(t, "#b100", serv.cur)

	Push("/test_state#1")
	<-serv.ch
	assert.Equal(t, "#b100", serv.old)
	assert.Equal(t, "#1", serv.cur)

	Push("/test_state#2")
	<-serv.ch
	assert.Equal(t, "#1", serv.old)
	assert.Equal(t, "#2", serv.cur)
}
