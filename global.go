package js

import (
	"syscall/js"
	gojs "syscall/js"
)

// ----------------------------------------------------------------------------

// Error ...
type Error struct {
	ref gojs.Value
}

// ErrorOf ...
func ErrorOf(x gojs.Value) Error {
	return Error{ref: x}
}

// JSValue ...
func (err Error) JSValue() gojs.Value {
	return err.ref
}

func (err Error) Error() string {
	return err.ref.String()
}

// ----------------------------------------------------------------------------

// Event ...
type Event struct {
	ref js.Value
}

// JSValue ...
func (e Event) JSValue() js.Value {
	return e.ref
}

// EventOf ...
func EventOf(x js.Value) Event {
	return Event{ref: x}
}

// Target ...
func (e Event) Target() js.Value {
	return e.ref.Get("target")
}
