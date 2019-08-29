// +build js,wasm

package url_test

import (
	"testing"

	"github.com/dairaga/js/url"
)

func TestURL(t *testing.T) {
	//link := `https://developer.mozilla.org:4097/en-US/docs/Web/API/URL/host?a=b&c=d#tag`
	link := `blob:https://localhost:8088/16bb5745-6242-4527-8cd8-90ab53bc1682`

	myurl := url.New(link)

	t.Log(myurl.Host())
	t.Log(myurl.Hostname())
	t.Log(myurl.Pathname())
	t.Log(myurl.Protocol())
	t.Log(myurl)

	myurl.Params().Foreach(func(k, v string) {
		t.Logf("key: %q, val: %q", k, v)
	})
}
