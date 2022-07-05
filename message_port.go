//go:build js && wasm

package js

import "syscall/js"

type MessagePort Value

// -----------------------------------------------------------------------------

func (m MessagePort) JSValue() Value {
	return Value(m)
}

// -----------------------------------------------------------------------------

func (m MessagePort) Start() {
	js.Value(m).Call("start")
}

// -----------------------------------------------------------------------------

func (m MessagePort) PostMessage(v Value) {
	js.Value(m).Call("postMessage", v)
}

// -----------------------------------------------------------------------------

func (m MessagePort) Close() {
	js.Value(m).Call("close")
}
