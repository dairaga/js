//go:build js && wasm

package js

import "fmt"

type Appendable interface {
	Wrapper
	Ref() Value
}

// -----------------------------------------------------------------------------

var document = global.Get("document")
var body = document.Get("body")

// -----------------------------------------------------------------------------

func Query(selector string) Element {
	return elementOf(query(document, selector))
}

// -----------------------------------------------------------------------------

func QueryAll(selector string) Elements {
	return ElementsOf(queryAll(document, selector))
}

// -----------------------------------------------------------------------------

func Append(a Appendable) {
	body.Call("append", a.Ref())
}

// -----------------------------------------------------------------------------

func Prepend(a Appendable) {
	body.Call("prepend", a.Ref())
}

// -----------------------------------------------------------------------------

func RemoveChild(x any) {
	switch v := x.(type) {
	case string:
		query(document, v).Call("remove")
	case Value:
		v.Call("remove")
	case Wrapper:
		v.JSValue().Call("remove")
	default:
		panic(fmt.Sprintf("unsupport type %T", x))
	}
}

// -----------------------------------------------------------------------------

func createElement(tag string) Value {
	return document.Call("createElement", tag)
}

// -----------------------------------------------------------------------------

func CreateElement(tag string) Element {
	return elementOf(createElement(tag))
}
