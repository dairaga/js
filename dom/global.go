package dom

import (
	"github.com/dairaga/js"
)

var document = ElementOf(js.Global().Get("document"))

// ----------------------------------------------------------------------------

// document

// CreateElement returns a HTML element.
func CreateElement(tag string) *Element {
	return ElementOf(document.Call("createElement", tag))
}

// AppendChild appends child to document.
func AppendChild(child interface{}) {
	document.Prop("body").Call("appendChild", child)
}

// RemoveChild removes child form document.
func RemoveChild(child *Element) {
	//document.Call("removeChild", child)
	document.Prop("body").Call("removeChild", child)
}

// S returns one element in document with query selector condition.
func S(selector string) *Element {
	return document.S(selector)
}

// SS returns elements in document with query selector condition.
func SS(selector string) ElementList {
	return document.SS(selector)
}
