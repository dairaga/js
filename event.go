//go:build js && wasm

package js

import (
	"github.com/dairaga/js/v3/builtin"
)

type Event interface {
	Wrapper
	Type() string
	Get(name string) Value
	PreventDefault()
}

// -----------------------------------------------------------------------------

type event Value

func (e event) JSValue() Value {
	return Value(e)
}

// -----------------------------------------------------------------------------

func (e event) Get(name string) Value {
	return Value(e).Get(name)
}

// -----------------------------------------------------------------------------

func (e event) Type() string {
	return e.Get("type").String()
}

// -----------------------------------------------------------------------------

func (e event) PreventDefault() {
	Value(e).Call("preventDefault")
}

// -----------------------------------------------------------------------------

func EventOf(v Value) Event {
	if !builtin.Event.Is(v) {
		panic(ValueError{
			Method: "EventOf",
			Type:   v.Type(),
		})
	}
	return event(v)
}
