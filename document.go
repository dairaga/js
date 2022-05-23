//go:build js && wasm

package js

type Appendable interface {
	Wrapper
	Ref() Value
}

// -----------------------------------------------------------------------------

var document = global.Get("document")
var body = document.Get("body")

// -----------------------------------------------------------------------------

func Query(selector string) Element {
	return query(document, selector)
}

// -----------------------------------------------------------------------------

func QueryAll(selector string) Elements {
	return queryAll(document, selector)
}

// -----------------------------------------------------------------------------

func Append(a Appendable) {
	body.Call("append", a.Ref())
}

// -----------------------------------------------------------------------------

func Prepend(a Appendable) {
	body.Call("prepend", a.Ref())
}
