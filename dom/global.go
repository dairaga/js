package dom

import (
	"github.com/dairaga/js"
)

var document = ElementOf(js.Global().Get("document"))

// ----------------------------------------------------------------------------

// document

// CreateElement ...
func CreateElement(tag string) Element {
	return ElementOf(document.Call("createElement", tag))
}

// AppendChild ...
func AppendChild(child Element) {
	document.Call("appendChild", child)
}

// RemoveChild ...
func RemoveChild(child Element) {
	document.Call("removeChild", child)
}

// S ...
func S(selector string) Element {
	return document.S(selector)
}

// SS ...
func SS(selector string) NodeList {
	return document.SS(selector)
}
