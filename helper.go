//go:build js && wasm

package js

func query(v Value, selector string) Element {
	return ElementOf(v.Call("querySelector", selector))
}

// -----------------------------------------------------------------------------

func queryAll(v Value, selector string) Elements {
	return ElementsOf(v.Call("querySelectorAll", selector))
}
