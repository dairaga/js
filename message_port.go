//go:build js && wasm

package js

import (
	"github.com/dairaga/js/v2/builtin"
)

type MessagePort Value

// -----------------------------------------------------------------------------

func (m MessagePort) JSValue() Value {
	return Value(m)
}

// -----------------------------------------------------------------------------

func (m MessagePort) Start() {
	Value(m).Call("start")
}

// -----------------------------------------------------------------------------

func (m MessagePort) PostMessage(v Value) {
	Value(m).Call("postMessage", v)
}

// -----------------------------------------------------------------------------

func (m MessagePort) Close() {
	Value(m).Call("close")
}

// -----------------------------------------------------------------------------

func (m MessagePort) OnMessage(fn func(Value)) {
	Value(m).Call("addEventListener", "message", FuncOf(func(_ Value, args []Value) any {
		fn(args[0])
		return nil
	}))
}

// -----------------------------------------------------------------------------

func (m MessagePort) OnError(fn func(Value)) {
	Value(m).Call("addEventListener", "messageerror", FuncOf(func(this Value, args []Value) any {
		fn(args[0])
		return nil
	}))
}

// -----------------------------------------------------------------------------

func MessagePortOf(v Value) MessagePort {
	if !builtin.MessagePort.Is(v) {
		panic(ValueError{
			Method: "MessagePortOf",
			Type:   v.Type(),
		})
	}

	return MessagePort(v)
}
