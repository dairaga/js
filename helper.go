//go:build js && wasm

package js

import "syscall/js"

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

func appendNode(parent Value, child Appendable) Value {
	parent.Call("append", child.Ref())
	return parent
}

// -----------------------------------------------------------------------------

func prependNode(parent Value, child Appendable) Value {
	parent.Call("prepend", child.Ref())
	return parent
}

// -----------------------------------------------------------------------------

func attr(v Value, a string) string {
	val := v.Call("getAttribute", a)
	if val.Truthy() && val.Type() == js.TypeString {
		return val.String()
	}

	return ""
}

// -----------------------------------------------------------------------------

func setAttr(v Value, a, val string) Value {
	if a != _tattoo {
		v.Call("setAttribute", a, val)
	}
	return v
}

// -----------------------------------------------------------------------------

func addClz(v Value, clz string) Value {
	v.Get("classList").Call("add", clz)
	return v
}

// -----------------------------------------------------------------------------

func removeClz(v Value, clz string) Value {
	v.Get("classList").Call("remove", clz)
	return v
}

func hasClz(v Value, clz string) bool {
	return v.Get("classList").Call("contains", clz).Bool()
}

// -----------------------------------------------------------------------------

func replaceClz(v Value, old, new string) Value {
	v.Get("classList").Call("replace", old, new)
	return v
}

// -----------------------------------------------------------------------------

func toggleClz(v Value, clz string) Value {
	v.Get("classList").Call("toggle", clz)
	return v
}

// -----------------------------------------------------------------------------

func fragment(node Value) Value {
	//return node.Call("cloneNode", true).Get("firstElementChild")
	return node.Get("firstElementChild")
}

// -----------------------------------------------------------------------------

func at(parent Value, selector ...string) Value {
	if len(selector) > 0 {
		child := query(parent, selector[0])
		tattoos(child)
		return child
	}
	return parent
}
