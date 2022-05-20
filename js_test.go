//go:build js && wasm

package js

import (
	"os"
	"os/signal"
	"testing"

	"github.com/stretchr/testify/assert"

	gojs "syscall/js"
)

type jsv Value

func (v jsv) JSValue() Value {
	return Value(v)
}

// -----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	headless := os.Getenv("WASM_HEADLESS")
	exitVal := m.Run()

	if headless == "off" {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		// Block until a signal is received.
		<-c
	}

	os.Exit(exitVal)
}

// -----------------------------------------------------------------------------

func TestValueOf(t *testing.T) {

	var x any

	x = int(10)

	v := ValueOf(x)

	assert.Equal(t, 10, v.Int())

	v = ValueOf(gojs.ValueOf(10))

	assert.Equal(t, 10, v.Int())

	v = ValueOf(jsv(v))
	assert.Equal(t, 10, v.Int())
}
