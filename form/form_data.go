//go:build js && wasm

package form

import (
	"fmt"

	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

// ValueType is Form Value type. There are two types: String and Binary.
type ValueType int

const (
	TypeString ValueType = ValueType(js.TypeString) // StringType
	TypeBinary ValueType = ValueType(js.TypeObject) // BinaryType
)

// -----------------------------------------------------------------------------

// FormValue represents Javascript Form Value.
type FormValue interface {
	// Type returns the value type.
	Type() ValueType

	// ValueRef returns javascript value.
	ValueRef() js.Value
}

// -----------------------------------------------------------------------------

// String represents String value.
type String string

// ValueRef returns javascript value.
func (s String) ValueRef() js.Value {
	return js.ValueOf(string(s))
}

// -----------------------------------------------------------------------------

// Type returns string type.
func (s String) Type() ValueType {
	return TypeString
}

// -----------------------------------------------------------------------------

// Binary represents Binary value.
type Binary js.Value

// ValueRef returns javascript value.
func (b Binary) ValueRef() js.Value {
	v := js.Value(b)
	if builtin.Blob.Is(v) {
		return v
	}

	panic(fmt.Sprintf("unsupported type %v", v.Type()))
}

// -----------------------------------------------------------------------------

// Type return binary type.
func (b Binary) Type() ValueType {
	return TypeBinary
}

// -----------------------------------------------------------------------------

// FormData is javascript FormData.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/FormData.
type FormData js.Value

// -----------------------------------------------------------------------------

// JSValue returns javascript value.
func (f FormData) JSValue() js.Value {
	return js.Value(f)
}

// -----------------------------------------------------------------------------

// Get returns value associated with the given name.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/FormData/get.
func (f FormData) Get(name string) FormValue {
	val := js.Value(f).Call("get", name)
	if js.TypeString == val.Type() {
		return String(val.String())
	}

	return Binary(val)
}

// -----------------------------------------------------------------------------

// Set sets value associated with the given name to the given value val.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/FormData/set.
func (f FormData) Set(name string, val FormValue, filename ...string) {
	if len(filename) > 0 {
		js.Value(f).Call("set", name, val.ValueRef(), filename[0])
	} else {
		js.Value(f).Call("set", name, val.ValueRef())
	}
}

// -----------------------------------------------------------------------------

// Append appends the given value val onto the given key name.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/FormData/append.
func (f FormData) Append(name string, val FormValue, filename ...string) {
	if len(filename) > 0 {
		js.Value(f).Call("append", name, val.ValueRef(), filename[0])
	} else {
		js.Value(f).Call("append", name, val.ValueRef())
	}
}

// -----------------------------------------------------------------------------

// Delete removes all values associated with the given name.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/FormData/delete.
func (f FormData) Delete(name string) {
	js.Value(f).Call("delete", name)
}

// -----------------------------------------------------------------------------

// FormDataOf returns a new FormData instance. Construct a empty FormData if no arguments.
//
// Given x can be:
//
// 1. string as selector criteria.
//
// 2. javascript value or wrapper of javascript value.
func FormDataOf(x ...any) FormData {
	if len(x) <= 0 {
		return FormData(builtin.FormData.New())
	}

	val := js.Null()
	switch v := x[0].(type) {
	case string:
		val = js.Window().Get("document").Call("querySelector", v)
	case js.Wrapper:
		val = v.JSValue()
	case js.Value:
		val = v
	}

	if val.Truthy() && builtin.HTMLFormElement.Is(val) {
		return FormData(builtin.FormData.New(val))
	}

	panic(fmt.Sprintf("unsupported type %T", x))
}
