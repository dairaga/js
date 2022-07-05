//go:build js && wasm

package js

import (
	"syscall/js"

	"github.com/dairaga/js/v2/builtin"
)

type Promise interface {
	Wrapper

	Then(func(Value) any) Promise
	Catch(func(Value) any) Promise
	Finally(func() any) Promise
	Await() Value
}

type promise js.Value

var _ Promise = promise{}

// -----------------------------------------------------------------------------

func (p promise) JSValue() Value {
	return Value(p)
}

// -----------------------------------------------------------------------------

func (p promise) call(method string, f func(Value) any) Promise {
	result := Value(p).Call(method, js.FuncOf(func(_ js.Value, args []js.Value) any {
		return f(args[0])
	}))
	return promise(result)
}

// -----------------------------------------------------------------------------

func (p promise) Then(f func(Value) any) Promise {
	return p.call("then", f)
}

// -----------------------------------------------------------------------------

func (p promise) Catch(f func(Value) any) Promise {
	return p.call("catch", f)
}

// -----------------------------------------------------------------------------

func (p promise) Finally(f func() any) Promise {
	result := Value(p).Call("finally", js.FuncOf(func(js.Value, []js.Value) any {
		return f()
	}))
	return promise(result)
}

// -----------------------------------------------------------------------------

func (p promise) Await() Value {
	ch := make(chan Value)
	defer func() {
		close(ch)
	}()

	p.Then(func(v Value) any {
		ch <- v
		return nil
	}).Catch(func(v Value) any {
		ch <- v
		return nil
	})

	return <-ch
}

// -----------------------------------------------------------------------------

func PromiseOf(x Value) Promise {
	if !builtin.Promise.Is(x) {
		panic(ValueError{
			Method: "PromiseOf",
			Type:   x.Type(),
		})
	}

	return promise(x)
}
