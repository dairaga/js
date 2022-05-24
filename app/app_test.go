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

func (s *serv) Serve(curURL url.URL, curHash, oldHash string) {
	s.old = oldHash
	s.cur = curHash
	s.ch <- struct{}{}
}

func TestHash(t *testing.T) {
	serv := &serv{
		t:  t,
		ch: make(chan struct{}, 1),
	}

	//curURL := url.New(js.Global().Get("location").Get("href").String())
	Init(serv)
	_app.handler.Serve(_app.currentURL, _app.currentURL.Hash(), "")
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
}
