//go:build js && wasm

package js

import "fmt"

// Appendable 代表物件是可被加到 Body 或 HTMLElement 內。目前有實作的物件有： Element, HTML, Plain, Template.
type Appendable interface {
	Wrapper
	Ref() Value
}

// -----------------------------------------------------------------------------

// Query 呼叫 Document.querySelector，並回傳 Element。
func Query(selector string) Element {
	return elementOf(query(document, selector))
}

// -----------------------------------------------------------------------------

// QueryAll 呼叫 Document.querySelectorAll，並回傳 Elements。
func QueryAll(selector string) Elements {
	return ElementsOf(queryAll(document, selector))
}

// -----------------------------------------------------------------------------

// Append 將物件加到 Body 的尾端。
func Append(child Appendable, selector ...string) {
	appendNode(at(body, selector...), child)
}

// -----------------------------------------------------------------------------

// Prepend 將物件加到 Body 的前面。
func Prepend(child Appendable, selector ...string) {
	//body.Call("prepend", a.Ref())
	prependNode(at(body, selector...), child)
}

// -----------------------------------------------------------------------------

// RemoveChild 將 Body 內的物件移除。
// x 可以是：
// 1. string: 需遵守 Selector 規則。
// 2. Value: 呼叫 remove 函式，Value 必須是 Javascript Element。
// 3. Wrapper: 呼叫 JSValue() 取得 Value 後，呼叫 remove，因此取得的 Value 物件，必須是 Javascript Element.
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
