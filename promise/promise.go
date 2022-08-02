//go:build js && wasm

// Pacakage promise is provides static functions of Javascript Promise.
//
// See https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise.
package promise

import (
	"github.com/dairaga/js/v3"
	"github.com/dairaga/js/v3/builtin"
)

// Resolve returns a Promise that is resolved the given value x.
//
// See https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/resolve.
func Resolve(x any) js.Promise {
	result := builtin.Promise.JSValue().Call("resolve", x)
	return js.PromiseOf(result)
}

// -----------------------------------------------------------------------------

// Reoject returns a Promise that is rejected the given value x.
//
// See https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/reject.
func Reject(x any) js.Promise {
	result := builtin.Promise.JSValue().Call("reject", x)
	return js.PromiseOf(result)
}

// -----------------------------------------------------------------------------

// All takes given promises and returns a single promise. The single promise contains resolved values if the given promises are resolved,
// or contains the first rejected value.
//
// See https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/all.
func All(promises []js.Promise) js.Promise {
	values := make([]any, len(promises))
	for i := range promises {
		values[i] = promises[i].JSValue()
	}

	result := builtin.Promise.JSValue().Call("all", values)

	return js.PromiseOf(result)
}

// -----------------------------------------------------------------------------

// Any returns a single promise that is resolved with the first fulfilled value of the given promises.
//
// See https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/any.
func Any(promises []js.Promise) js.Promise {
	values := make([]any, len(promises))
	for i := range promises {
		values[i] = promises[i].JSValue()
	}

	result := builtin.Promise.JSValue().Call("any", values)

	return js.PromiseOf(result)
}
