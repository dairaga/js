//go:build js && wasm

package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaders(t *testing.T) {

	headers := NewHeaders([2]string{"Content-Type", "image/jpeg"}, [2]string{"X-Custom-Header", "value"}, [2]string{"Accept-Encoding", "deflate"})

	assert.True(t, headers.Has("Content-Type"))
	assert.False(t, headers.Has("X-Token"))

	val, ok := headers.Get("Content-Type")
	assert.True(t, ok)
	assert.Equal(t, "image/jpeg", val)

	val, ok = headers.Get("X-Token")
	assert.False(t, ok)
	assert.Equal(t, "", val)

	headers.Append("Accept-Encoding", "gzip")

	val, _ = headers.Get("Accept-Encoding")
	assert.Equal(t, "deflate, gzip", val)

	headers.Set("X-Custom-Header", "value2")
	val, _ = headers.Get("X-Custom-Header")
	assert.Equal(t, "value2", val)

	headers.Set("X-Token", "token")
	val, _ = headers.Get("X-Token")
	assert.Equal(t, "token", val)

	headers.Delete("X-Token")
	assert.False(t, headers.Has("X-Token"))

	count := 0
	headers.Foreach(func(_, _ string) {
		count++
	})

	assert.Equal(t, 3, count)

}
