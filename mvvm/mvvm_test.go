//go:build js && wasm

package mvvm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMVVM(t *testing.T) {

	a := false
	b := []string{}

	Add("a", &a)
	Add("b", &b)

	assert.Panics(t, func() { Add("xa", a) })
	assert.Panics(t, func() { Add("xa", b) })
	assert.Panics(t, func() { Add("a", b) })
	assert.Panics(t, func() { Add("b", a) })

	triggerA := false
	triggerB := false
	Watch("a", func(sender string, v bool) {
		triggerA = true
		assert.Equal(t, a, v)
	})

	Watch("b", func(sender string, v []string) {
		triggerB = true
		assert.Equal(t, b, v)
	})

	a = true
	Trigger("test", "a")
	assert.True(t, triggerA)

	b = append(b, "A")
	Trigger("test", "b")
	assert.True(t, triggerB)

}
