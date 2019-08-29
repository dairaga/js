// +build js,wasm

package js

import (
	"syscall/js"
	"testing"
)

func TestBytes(t *testing.T) {
	x := js.Global().Get("Uint8Array").New(2)
	x.SetIndex(0, uint8(21))
	x.SetIndex(1, uint8(31))

	y := Bytes(x)

	t.Log(y)
}
