//go:build js && wasm

package js

import (
	"github.com/dairaga/js/v3/builtin"
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

func (m MessagePort) OnMessage(fn func(Event)) Func {
	jsfn := FuncOf(func(_ Value, args []Value) any {
		fn(event(args[0]))
		return nil
	})

	Value(m).Call("addEventListener", "message", jsfn)
	return jsfn
}

// -----------------------------------------------------------------------------

func (m MessagePort) OnError(fn func(Value)) Func {
	jsfn := FuncOf(func(this Value, args []Value) any {
		fn(args[0])
		return nil
	})
	Value(m).Call("addEventListener", "messageerror", jsfn)
	return jsfn
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
