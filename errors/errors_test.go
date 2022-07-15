//go:build js && wasm

package errors

import (
	"errors"
	"testing"

	"github.com/dairaga/js/v2/builtin"
	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {

	var err error = TypeError(errors.New("type error"))

	jserr, ok := AsJSError(err)
	assert.True(t, ok)
	assert.True(t, Is(jserr, builtin.TypeError))

	err = errors.New("general error")
	_, ok = AsJSError(err)
	assert.False(t, ok)
}
