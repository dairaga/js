//go:build js && wasm

package app

import (
	"testing"
	"time"

	"github.com/dairaga/js/v2/url"
	"github.com/stretchr/testify/assert"
)

type serv struct {
	t  *testing.T
	ch chan struct{}
	a  string
	b  string
}

func (s *serv) ServeHash(curURL url.URL, oldHash, newHash string) {
	s.a = oldHash
	s.b = newHash
	s.ch <- struct{}{}
}

func TestHash(t *testing.T) {
	serv := &serv{
		t:  t,
		ch: make(chan struct{}, 1),
	}

	//curURL := url.New(js.Global().Get("location").Get("href").String())
	ServHash(serv)
	<-serv.ch
	assert.Equal(t, "", serv.a)
	assert.Equal(t, "", serv.b)

	SetHash("#a100")
	<-serv.ch
	assert.Equal(t, "", serv.a)
	assert.Equal(t, "#a100", serv.b)

	SetHash("#b100")
	<-serv.ch
	assert.Equal(t, "#a100", serv.a)
	assert.Equal(t, "#b100", serv.b)

	SetHash("#b100")
	select {
	case <-time.After(3 * time.Second):
	case <-serv.ch:
	}
	// the hashchange event is not triggered.
	assert.Equal(t, "#a100", serv.a)
	assert.Equal(t, "#b100", serv.b)
}
