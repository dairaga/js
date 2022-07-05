//go:build js && wasm

package audio

import "github.com/dairaga/js/v2"

type Node interface {
	js.Wrapper

	ChannelCount() int
	Context() Context

	NumberOfInputs() int
	NumberOfOutputs() int

	Connect(Node)
	Disconnect()
}

type node js.Value

var _ Node = node{}

// -----------------------------------------------------------------------------

func (n node) JSValue() js.Value {
	return js.Value(n)
}

// -----------------------------------------------------------------------------

func (n node) ChannelCount() int {
	return js.Value(n).Get("channelCount").Int()
}

// -----------------------------------------------------------------------------

func (n node) Context() Context {
	return Context(js.Value(n).Get("context"))
}

// -----------------------------------------------------------------------------

func (n node) NumberOfInputs() int {
	return js.Value(n).Get("numberOfInputs").Int()
}

// -----------------------------------------------------------------------------

func (n node) NumberOfOutputs() int {
	return js.Value(n).Get("numberOfOutputs").Int()
}

// -----------------------------------------------------------------------------

func (n node) Connect(other Node) {
	js.Value(n).Call("connect", other.JSValue())
}

// -----------------------------------------------------------------------------

func (n node) Disconnect() {
	js.Value(n).Call("disconnect")
}
