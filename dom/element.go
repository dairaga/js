package dom

import (
	"fmt"
	"strings"
	gojs "syscall/js"

	"github.com/dairaga/js"
)

// Element represents a HTML element.
type Element struct {
	*js.EventTarget
}

// ElementOf returns a HTML element.
func ElementOf(v js.Value) *Element {
	return &Element{js.EventTargetOf(v)}
}

// ----------------------------------------------------------------------------

func (e *Element) String() string {
	if !e.Truthy() {
		return e.JSValue().String()
	}

	id := e.Attr("id")
	tag := e.TagName()
	content := e.HTML()

	return fmt.Sprintf(`tag: %s, id:%q, content: %s`, tag, id, content)
}

// S returns child by quering selector condition.
func (e *Element) S(selector string) *Element {
	return ElementOf(e.JSValue().Call("querySelector", selector))
}

// SS returns children by query selector condition.
func (e *Element) SS(selector string) ElementList {
	return ElementListOf(e.JSValue().Call("querySelectorAll", selector))
}

// ----------------------------------------------------------------------------

// Truthy returns javascript truthy value.
func (e *Element) Truthy() bool {
	return e.JSValue().Truthy()
}

// ----------------------------------------------------------------------------

// Attr returns attribute value.
func (e *Element) Attr(name string) string {
	return e.JSValue().Call("getAttribute", name).String()
}

// SetAttr sets attribute.
func (e *Element) SetAttr(name, value string) *Element {
	e.JSValue().Call("setAttribute", name, value)
	return e
}

// ----------------------------------------------------------------------------

// Prop returns property of element.
func (e *Element) Prop(name string) js.Value {
	return e.JSValue().Get(name)
}

// SetProp sets property.
func (e *Element) SetProp(name string, val interface{}) *Element {
	e.JSValue().Set(name, val)
	return e
}

// ----------------------------------------------------------------------------

// SetText set inner text of element.
func (e *Element) SetText(text string) *Element {
	e.JSValue().Set("innerText", text)
	return e
}

// Text returns inner text of element.
func (e *Element) Text() string {
	return e.JSValue().Get("innerText").String()
}

// SetHTML set inner html of element.
func (e *Element) SetHTML(html string) *Element {
	e.JSValue().Set("innerHTML", html)
	return e
}

// HTML return inner html of element.
func (e *Element) HTML() string {
	return e.JSValue().Get("innerHTML").String()
}

// ----------------------------------------------------------------------------

func (e *Element) clz(m string, args ...string) *Element {
	size := len(args)
	if size <= 0 {
		return e
	}

	if size == 1 {
		e.JSValue().Get("classList").Call(m, args[0])
	} else if size == 2 {
		e.JSValue().Get("classList").Call(m, args[0], args[1])
	} else if size == 3 {
		e.JSValue().Get("classList").Call(m, args[0], args[1], args[2])
	} else {
		x := make([]interface{}, size)
		for i, str := range args {
			x[i] = str
		}
		e.JSValue().Get("classList").Call(m, x...)
	}

	return e
}

// AddClass add class to element.
func (e *Element) AddClass(names ...string) *Element {
	return e.clz("add", names...)
}

// RemoveClass remove class from element.
func (e *Element) RemoveClass(names ...string) *Element {
	return e.clz("remove", names...)
}

// ToggleClass toggle some class of element.
func (e *Element) ToggleClass(name string) *Element {
	return e.clz("toggle", name)
}

// ReplaceClass replace some class of element with new one.
func (e *Element) ReplaceClass(oldName, newName string) *Element {
	return e.clz("replace", oldName, newName)
}

// HasClass returns boolean indicates whether or not element has the class.
func (e *Element) HasClass(name string) bool {
	return e.JSValue().Get("classList").Call("contains", name).Bool()
}

// ----------------------------------------------------------------------------

// TagName returns tag name (all lower case) of element.
func (e *Element) TagName() string {
	return strings.ToLower(e.JSValue().Get("tagName").String())
}

// ----------------------------------------------------------------------------

// Val returns value of form input, select or textarea element.
func (e *Element) Val() string {

	switch e.TagName() {
	case "input", "select":
		return e.JSValue().Get("value").String()
	case "textarea":
		return e.JSValue().Get("innerText").String()
	}
	return gojs.Undefined().String()
}

// SetVal set value to input, select, or textarea element.
func (e *Element) SetVal(val interface{}) *Element {
	switch e.TagName() {
	case "input", "select":
		e.JSValue().Set("value", val)
	case "textarea":
		e.JSValue().Set("innerText", val)
	}
	return e
}

// ----------------------------------------------------------------------------

// Append add child to element.
func (e *Element) Append(child ...interface{}) *Element {
	for _, x := range child {
		e.JSValue().Call("append", x)
	}

	return e
}

// RemoveChild removes child element.
func (e *Element) RemoveChild(elm *Element) *Element {
	e.JSValue().Call("removeChild", elm)
	return e
}

// Clone clones the element. https://developer.mozilla.org/en-US/docs/Web/API/Node/cloneNode.
func (e *Element) Clone() *Element {
	return ElementOf(e.JSValue().Call("cloneNode", true))
}

// ----------------------------------------------------------------------------

// On add listener for some event.
func (e *Element) On(event string, fn func(*Element, *js.Event)) *Element {
	cb := js.FuncOf(func(_this js.Value, args []js.Value) interface{} {
		fn(ElementOf(_this), js.EventOf(args[0]))
		return nil
	})

	e.EventTarget.On(event, cb)
	return e
}

// OnClick adds callback function when clicking element.
func (e *Element) OnClick(fn func(*Element, *js.Event)) *Element {
	return e.On("click", fn)
}

// OnChange adds callback function when element value changed.
func (e *Element) OnChange(fn func(*Element, *js.Event)) *Element {
	return e.On("change", fn)
}

// ----------------------------------------------------------------------------

// Call invokes element method.
func (e *Element) Call(name string, args ...interface{}) js.Value {
	return e.JSValue().Call(name, args...)
}
