package js

// Event represents javascript Event. https://developer.mozilla.org/en-US/docs/Web/API/Event
type Event struct {
	ref Value
}

// JSValue ...
func (e *Event) JSValue() Value {
	return e.ref
}

// EventOf returns event.
func EventOf(x Value) *Event {
	return &Event{ref: x}
}

// Target https://developer.mozilla.org/en-US/docs/Web/API/Event/target
func (e *Event) Target() Value {
	return e.ref.Get("target")
}

// Type https://developer.mozilla.org/en-US/docs/Web/API/Event/type
func (e *Event) Type() string {
	return e.ref.Get("type").String()
}
