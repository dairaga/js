//go:build js && wasm

package js

func query(v Value, selector string) Value {
	return v.Call("querySelector", selector)
}

// -----------------------------------------------------------------------------

func queryAll(v Value, selector string) Value {
	return v.Call("querySelectorAll", selector)
}

// -----------------------------------------------------------------------------

func fragment(node Value) Value {
	return node.Call("cloneNode", true).Get("firstElementChild")
}
