//go:build js && wasm

package audio

import "github.com/dairaga/js/v3"

type Param js.Value

// -----------------------------------------------------------------------------

func (p Param) JSValue() js.Value {
	return js.Value(p)
}

// -----------------------------------------------------------------------------

func (p Param) DefaultValue() js.Value {
	return js.Value(p).Get("defaultValue")
}

// -----------------------------------------------------------------------------

func (p Param) MaxValue() js.Value {
	return js.Value(p).Get("maxValue")
}

// -----------------------------------------------------------------------------

func (p Param) MinValue() js.Value {
	return js.Value(p).Get("minValue")
}

// -----------------------------------------------------------------------------

func (p Param) Value() js.Value {
	return js.Value(p).Get("value")
}

// -----------------------------------------------------------------------------

func (p Param) SetValueAtTime(value js.Value, start float64) {
	js.Value(p).Call("setValueAtTime", value, start)
}

// -----------------------------------------------------------------------------

type ParamMap js.Value

func (m ParamMap) JSValue() js.Value {
	return js.Value(m)
}

// -----------------------------------------------------------------------------

func (m ParamMap) Get(name string) Param {
	return Param(js.Value(m).Call("get", name))
}
