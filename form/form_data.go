//go:build js && wasm

package form

import (
	"fmt"

	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

type ValueType int

const (
	TypeString ValueType = ValueType(js.TypeString)
	TypeBinary ValueType = ValueType(js.TypeObject)
)

// -----------------------------------------------------------------------------

type FormValue interface {
	Type() ValueType
	ValueRef() js.Value
}

// -----------------------------------------------------------------------------

type String string

func (s String) ValueRef() js.Value {
	return js.ValueOf(string(s))
}

// -----------------------------------------------------------------------------

func (s String) Type() ValueType {
	return TypeString
}

// -----------------------------------------------------------------------------

type Binary js.Value

func (b Binary) ValueRef() js.Value {
	v := js.Value(b)
	if builtin.IsBlob(v) {
		return v
	}

	panic(fmt.Sprintf("unsupported type %v", v.Type()))
}

// -----------------------------------------------------------------------------

func (b Binary) Type() ValueType {
	return TypeBinary
}

// -----------------------------------------------------------------------------

type FormData js.Value

// -----------------------------------------------------------------------------

func (f FormData) JSValue() js.Value {
	return js.Value(f)
}

// -----------------------------------------------------------------------------

func (f FormData) Get(name string) FormValue {
	val := js.Value(f).Call("get", name)
	if js.TypeString == val.Type() {
		return String(val.String())
	}

	return Binary(val)
}

// -----------------------------------------------------------------------------

func (f FormData) Set(name string, val FormValue, filename ...string) {
	if len(filename) > 0 {
		js.Value(f).Call("set", name, val.ValueRef(), filename[0])
	} else {
		js.Value(f).Call("set", name, val.ValueRef())
	}
}

// -----------------------------------------------------------------------------

func (f FormData) Append(name string, val FormValue, filename ...string) {
	if len(filename) > 0 {
		js.Value(f).Call("append", name, val.ValueRef(), filename[0])
	} else {
		js.Value(f).Call("append", name, val.ValueRef())
	}
}

// -----------------------------------------------------------------------------

func (f FormData) Delete(name string) {
	js.Value(f).Call("delete", name)
}

// -----------------------------------------------------------------------------

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

	if val.Truthy() && builtin.IsForm(val) {
		return FormData(builtin.FormData.New(val))
	}

	panic(fmt.Sprintf("unsupported type %T", x))
}
