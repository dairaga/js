package dom

import (
	"strings"
	"syscall/js"
)

// Value ...
type Value struct {
	ref js.Value
}

// ValueOf ...
func ValueOf(x interface{}) Value {
	switch v := x.(type) {
	case Value:
		return v
	case js.Value:
		return Value{ref: v}
	default:
		return Value{ref: js.ValueOf(x)}
	}
}

// ----------------------------------------------------------------------------

// JSValue ...
func (v Value) JSValue() js.Value {
	return v.ref
}

// ----------------------------------------------------------------------------

func (v Value) call(m string, args ...interface{}) js.Value {
	return v.ref.Call(m, args...)
}

// Call ...
func (v Value) Call(m string, args ...interface{}) Value {
	return ValueOf(v.call(m, args...))
}

// ----------------------------------------------------------------------------

// Get ...
func (v Value) Get(name string) Value {
	return ValueOf(v.ref.Get(name))
}

// GetDeep ...
func (v Value) GetDeep(name string) Value {
	names := strings.Split(name, ".")

	result := v.ref
	for _, x := range names {
		result = result.Get(x)
		if !result.Truthy() {
			return undefined
		}
	}
	return Value{ref: result}
}

// Set ...
func (v Value) Set(name string, val interface{}) Value {
	v.ref.Set(name, val)
	return v
}

// ----------------------------------------------------------------------------

// Truthy ...
func (v Value) Truthy() bool {
	return v.ref.Truthy()
}

// ----------------------------------------------------------------------------

// String ...
func (v Value) String() string {
	return v.ref.String()
}

// Float64 ...
func (v Value) Float64() float64 {
	return v.ref.Float()
}

// Int ...
func (v Value) Int() int {
	return v.ref.Int()
}

// Bool ...
func (v Value) Bool() bool {
	return v.ref.Bool()
}

// ----------------------------------------------------------------------------

// Index ...
func (v Value) Index(i int) Value {
	return ValueOf(v.ref.Index(i))
}

// ----------------------------------------------------------------------------

// On ...
func (v Value) On(event string, fn func(Value, Event)) Value {
	cb := js.FuncOf(func(_this js.Value, args []js.Value) interface{} {
		fn(ValueOf(_this), EventOf(args[0]))
		return nil
	})

	v.call("addEventListener", event, cb)
	return v
}
