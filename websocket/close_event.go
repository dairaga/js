package websocket

import "github.com/dairaga/js"

// CloseEvent represents javascript close event. https://developer.mozilla.org/en-US/docs/Web/API/CloseEvent.
type CloseEvent struct {
	*js.Event
}

// CloseEventOf returns a close event.
func CloseEventOf(x js.Value) *CloseEvent {
	return &CloseEvent{js.EventOf(x)}
}

// Code returns close code from server.
func (e *CloseEvent) Code() uint16 {
	return uint16(e.JSValue().Get("code").Int())
}

// Reason returns eht reason the server closed the connection.
func (e *CloseEvent) Reason() string {
	return e.JSValue().Get("reason").String()
}

// WasClean returns a boolean that indicates whether or not the connection was cleanly closed.
func (e *CloseEvent) WasClean() bool {
	return e.JSValue().Get("wasClean").Bool()
}
