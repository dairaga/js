//go:build js && wasm

package audio

import "github.com/dairaga/js/v2"

// Node is a Javascript AudioNode object.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/AudioNode.
type Node interface {
	js.Wrapper

	// ChannelCount returns the number of channels used.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/channelCount.
	ChannelCount() int

	// Context returns the associated AudioContext.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/context.
	Context() Context

	// NumberOfInputs returns the number of inputs feeding the node.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/numberOfInputs.
	NumberOfInputs() int

	// NumberOfOutputs returns the number of outputs coming out of the node.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/numberOfOutputs
	NumberOfOutputs() int

	// Connect connects to another node.
	//
	// https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/connect.
	Connect(Node)

	// Disconnect disconnects other nodes from it.
	//
	// See https://developer.mozilla.org/en-US/docs/Web/API/AudioNode/disconnect.
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
