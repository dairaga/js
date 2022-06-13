//go:build js && wasm

package js

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
