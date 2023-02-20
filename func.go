//go:build js && wasm

package js

import (
	"syscall/js"
)

var onceFunc map[string]js.Func

// -----------------------------------------------------------------------------

// FuncOf 原 syscall/js 的 FuncOf。
func FuncOf(fn JSFunc) Func {
	return js.FuncOf(fn)
}

// -----------------------------------------------------------------------------

func OnceFuncOf(fn JSFunc) Func {
	id := tattoo(10)

	ret := js.FuncOf(func(this Value, args []Value) any {
		defer func() {
			f, ok := onceFunc[id]
			if ok && f.Truthy() {
				delete(onceFunc, id)
				f.Release()
			}
		}()
		return fn(this, args)
	})

	return ret
}
