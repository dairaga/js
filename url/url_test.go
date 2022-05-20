//go:build js && wasm

package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const link = `https://developer.mozilla.org:4097/en-US/docs/Web/API/URL/host?a=b&c=d&aa=1&aa=2&aa=3#tag`

//link := `blob:https://localhost:8088/16bb5745-6242-4527-8cd8-90ab53bc1682`

func TestURL(t *testing.T) {

	myurl := New(link)

	assert.Equal(t, "https:", myurl.Protocol())
	assert.Equal(t, "developer.mozilla.org:4097", myurl.Host())
	assert.Equal(t, "developer.mozilla.org", myurl.Hostname())
	assert.Equal(t, "/en-US/docs/Web/API/URL/host", myurl.QueryPath())
	assert.Equal(t, "?a=b&c=d&aa=1&aa=2&aa=3", myurl.QueryString())
	assert.Equal(t, "#tag", myurl.Hash())
	assert.Equal(t, link, myurl.String())
}
