//go:build js && wasm

package errors

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

func New(message string, constructor ...builtin.Constructor) js.Error {
	for i := range constructor {
		return js.Error{
			Value: constructor[i].New(message),
		}
	}

	return js.Error{
		Value: builtin.Error.New(message),
	}
}

// -----------------------------------------------------------------------------

func EvalError(err error) js.Error {
	return New(err.Error(), builtin.EvalError)
}

// -----------------------------------------------------------------------------

func RangeError(err error) js.Error {
	return New(err.Error(), builtin.RangeError)
}

// -----------------------------------------------------------------------------

func ReferenceError(err error) js.Error {
	return New(err.Error(), builtin.ReferenceError)
}

// -----------------------------------------------------------------------------

func SyntaxError(err error) js.Error {
	return New(err.Error(), builtin.SyntaxError)
}

// -----------------------------------------------------------------------------

func TypeError(err error) js.Error {
	return New(err.Error(), builtin.TypeError)
}

// -----------------------------------------------------------------------------

func URIError(err error) js.Error {
	return New(err.Error(), builtin.URIError)
}

// -----------------------------------------------------------------------------

func AsJSError(err error) (result js.Error, ok bool) {
	result, ok = err.(js.Error)
	return
}

// -----------------------------------------------------------------------------

func Is(v js.Error, constructor builtin.Constructor) bool {
	return builtin.Is(v.Value, constructor)
}
