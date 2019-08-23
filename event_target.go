package js

// EventTarget represents javascript EventTarget.
type EventTarget struct {
	ref Value
	cb  map[string]Func
}

// EventTargetOf returns a event target object.
func EventTargetOf(v Value) *EventTarget {
	return &EventTarget{ref: v}
}

// JSValue ...
func (e *EventTarget) JSValue() Value {
	return e.ref
}

// Release frees up resoruces if the object is not used anymore.
func (e *EventTarget) Release() {
	for _, v := range e.cb {
		v.Release()
	}
}

// Register ...
func (e *EventTarget) Register(event string, cb Func) *EventTarget {
	old, ok := e.cb[event]
	if ok && old.Truthy() {
		old.Release()
	}

	e.cb[event] = cb
	return e
}

// AddEventListener adds callback for some event.
func (e *EventTarget) AddEventListener(event string, cb Func) *EventTarget {
	if e.cb == nil {
		e.cb = make(map[string]Func)
	}
	e.Register(event, cb)
	e.ref.Call("addEventListener", event, cb)
	return e
}
