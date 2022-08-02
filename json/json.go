//go:build js && wasm

package json

import (
	"github.com/dairaga/js/v3"
	"github.com/dairaga/js/v3/builtin"
	"github.com/dairaga/js/v3/errors"
)

var jsjson = js.Window().Get("JSON") // Javascript JSON.

const (
	TimeFormat = "2006-01-02T15:04:05.999Z"
)

// -----------------------------------------------------------------------------

//type Convertable interface {
//	FromValue(js.Value)
//	JSValue() js.Value
//}

// -----------------------------------------------------------------------------

type Marshaler interface {
	MarshalValue() (js.Value, error)
}

// -----------------------------------------------------------------------------

type Unmarshaler interface {
	UnmarshalValue(value js.Value) error
}

// -----------------------------------------------------------------------------

func Stringify(src js.Value) string {
	return jsjson.Call("stringify", src).String()
}

// -----------------------------------------------------------------------------

func Parse(src string) js.Value {
	return jsjson.Call("parse", src)
}

// -----------------------------------------------------------------------------

func MarshalValue(x Marshaler) (js.Value, error) {
	if x == nil {
		return js.Null(), nil
	}
	return x.MarshalValue()
}

// -----------------------------------------------------------------------------

func Marshal(x Marshaler) ([]byte, error) {
	val, err := MarshalValue(x)
	if err != nil {
		return nil, err
	}
	return []byte(Stringify(val)), nil
}

// -----------------------------------------------------------------------------

func ValidValue(val js.Value) error {
	if val.IsNull() || val.IsUndefined() || val.IsNaN() {
		return errors.New("data is not valid", builtin.TypeError)
	}
	return nil
}

// -----------------------------------------------------------------------------

func UnmarshalValue(val js.Value, x Unmarshaler) error {
	if err := ValidValue(val); err != nil {
		return err
	}

	return x.UnmarshalValue(val)
}

// -----------------------------------------------------------------------------

func Unmarshal(data []byte, x Unmarshaler) error {
	if len(data) <= 0 {
		return errors.New("data is empty", builtin.TypeError)
	}

	if x == nil {
		return errors.New("x is nil", builtin.TypeError)
	}

	return UnmarshalValue(Parse(string(data)), x)
}
