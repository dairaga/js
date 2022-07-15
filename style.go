//go:build js && wasm

package js

type Style Value

// -----------------------------------------------------------------------------

func (s Style) JSValue() Value {
	return Value(s)
}

// -----------------------------------------------------------------------------

func (s Style) GetPropertyValue(name string) string {
	return Value(s).Call("getPropertyValue", name).String()
}

// -----------------------------------------------------------------------------

func (s Style) RemoveProperty(name string) string {
	return Value(s).Call("removeProperty", name).String()
}

// -----------------------------------------------------------------------------

func (s Style) SetProperty(name, value string) {
	Value(s).Call("setProperty", name, value)
}

// -----------------------------------------------------------------------------

func (s Style) Length() int {
	return Value(s).Get("length").Int()
}

// -----------------------------------------------------------------------------

func (s Style) Item(idx int) string {
	return Value(s).Call("item", idx).String()
}
