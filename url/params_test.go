//go:build js && wasm

package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {

	myurl := New(link)

	params := myurl.Params()

	assert.Equal(t, []string{"1", "2", "3"}, params.GetAll("aa"))

	v, ok := params.Get("a")
	assert.True(t, params.Has("a"))
	assert.True(t, ok)
	assert.Equal(t, "b", v)

	v, ok = params.Get("x")
	assert.False(t, params.Has("x"))
	assert.False(t, ok)
	assert.Equal(t, "", v)

}
