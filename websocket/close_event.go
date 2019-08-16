package websocket

import "syscall/js"

// CloseEvent ...
type CloseEvent struct {
	ref js.Value
}

// CloseEventOf ...
func CloseEventOf(x js.Value) CloseEvent {
	return CloseEvent{ref: x}
}

// JSValue ...
func (e CloseEvent) JSValue() js.Value {
	return e.ref
}

// Code ...
func (e CloseEvent) Code() uint16 {
	return uint16(e.ref.Get("code").Int())
}

// Reason ...
func (e CloseEvent) Reason() string {
	return e.ref.Get("reason").String()
}

// WasClean ...
func (e CloseEvent) WasClean() bool {
	return e.ref.Get("wasClean").Bool()
}
