package media

import "syscall/js"

// Stream ...
type Stream struct {
	ref js.Value
}

// JSValue ...
func (stream Stream) JSValue() js.Value {
	return stream.ref
}

// Ready ...
func (stream Stream) Ready() bool {
	return stream.ref.Truthy()
}
