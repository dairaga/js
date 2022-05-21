//go:build js && wasm

package builtin

import (
	"fmt"
	"syscall/js"
)

type EventTarget interface {
	AddEventListener(typ string, fn func(js.Value, []js.Value) any)
	RemoveEventListener(typ string)
	DispatchEvent(event js.Value)
	Release()
	Released() bool
}

// -----------------------------------------------------------------------------

type eventTarget struct {
	ref       js.Value
	listeners map[string]js.Func
	released  bool
}

func (t *eventTarget) JSValue() js.Value {
	return t.ref
}

// -----------------------------------------------------------------------------

func (t *eventTarget) AddEventListener(typ string, fn func(js.Value, []js.Value) any) {
	cb := js.FuncOf(fn)
	old, ok := t.listeners[typ]
	if ok && old.Truthy() {
		old.Release()
	}
	t.listeners[typ] = cb
	t.ref.Call("addEventListener", cb)
}

// -----------------------------------------------------------------------------

func (t *eventTarget) RemoveEventListener(typ string) {
	old, ok := t.listeners[typ]
	if ok && old.Truthy() {
		t.ref.Call("removeEventListener", old)
		delete(t.listeners, typ)
		old.Release()
	}
}

// -----------------------------------------------------------------------------

func (t *eventTarget) DispatchEvent(event js.Value) {
	t.ref.Call("dispatchEvent", event)
}

// -----------------------------------------------------------------------------

func (t *eventTarget) Release() {
	if t.released {
		return
	}

	t.released = true

	for _, v := range t.listeners {
		v.Release()
	}
}

// -----------------------------------------------------------------------------

func (t *eventTarget) Released() bool {
	return t.released
}

// -----------------------------------------------------------------------------

func EventTargetOf(x any) EventTarget {
	switch v := x.(type) {
	case js.Value:
		return &eventTarget{
			ref:       v,
			listeners: make(map[string]js.Func, 0),
			released:  false,
		}
	case Wrapper:
		return EventTargetOf(v.JSValue())
	}
	panic(fmt.Sprintf("unsupport type %T", x))
}
