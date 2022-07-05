//go:build js && wasm

package js

import "github.com/dairaga/js/v2/builtin"

type Listener map[string]Func

func (l Listener) Add(target Value, typ string, fn JSFunc) Func {
	if !builtin.EventTarget.Is(target) {
		panic("target is not an EventTarget")
	}

	cb := FuncOf(fn)
	if old, ok := l[typ]; ok && old.Truthy() {
		old.Release()
	}
	l[typ] = cb
	target.Call("addEventListener", typ, cb)
	return cb
}

// -----------------------------------------------------------------------------

func (l Listener) Remove(target Value, typ string) {
	if !builtin.EventTarget.Is(target) {
		panic("target is not an EventTarget")
	}

	if old, ok := l[typ]; ok && old.Truthy() {
		delete(l, typ)
		target.Call("removeEventListener", typ, old)
		old.Release()
	}
}

// -----------------------------------------------------------------------------

func (l Listener) Release() {
	for _, fn := range l {
		fn.Release()
	}
}
