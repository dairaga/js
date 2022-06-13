//go:build js && wasm

package js

// query 呼叫 Document 或 HTMLElement 的 querySelector。
func query(v Value, selector string) Value {
	return v.Call("querySelector", selector)
}

// -----------------------------------------------------------------------------

// queryAll 呼叫 Document 或 HTMLElement 的 querySelectorAll。
func queryAll(v Value, selector string) Value {
	return v.Call("querySelectorAll", selector)
}

// -----------------------------------------------------------------------------

func fragment(node Value) Value {
	return node.Call("cloneNode", true).Get("firstElementChild")
}
