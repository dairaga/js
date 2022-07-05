//go:build js && wasm

package promise

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

func Resolve(x any) js.Promise {
	result := builtin.Promise.JSValue().Call("resolve", x)
	return js.PromiseOf(result)
}

// -----------------------------------------------------------------------------

func Reject(x any) js.Promise {
	result := builtin.Promise.JSValue().Call("reject", x)
	return js.PromiseOf(result)
}
