package dom

import (
	"strings"
	gojs "syscall/js"

	"github.com/dairaga/js"
)

// Element ...
type Element struct {
	*js.EventTarget
}

// ElementOf ...
func ElementOf(v js.Value) *Element {
	return &Element{js.EventTargetOf(v)}
}

// ----------------------------------------------------------------------------

// S ...
func (e *Element) S(selector string) *Element {
	return ElementOf(e.JSValue().Call("querySelector", selector))
}

// SS ...
func (e *Element) SS(selector string) NodeList {
	return NodeListOf(e.JSValue().Call("querySelectorAll", selector))
}

// ----------------------------------------------------------------------------

// Truthy ...
func (e *Element) Truthy() bool {
	return e.JSValue().Truthy()
}

// ----------------------------------------------------------------------------

// Attr ...
func (e *Element) Attr(name string) string {
	return e.JSValue().Call("getAttribute", name).String()
}

// SetAttr ...
func (e *Element) SetAttr(name, value string) *Element {
	e.JSValue().Call("setAttribute", name, value)
	return e
}

// ----------------------------------------------------------------------------

// Prop ...
func (e *Element) Prop(name string) js.Value {
	return e.JSValue().Get(name)
}

// SetProp ...
func (e *Element) SetProp(name string, val interface{}) *Element {
	e.JSValue().Set(name, val)
	return e
}

// ----------------------------------------------------------------------------

// SetText ...
func (e *Element) SetText(text string) *Element {
	e.JSValue().Set("innerText", text)
	return e
}

// Text ...
func (e *Element) Text() string {
	return e.JSValue().Get("innerText").String()
}

// SetHTML ...
func (e *Element) SetHTML(html string) *Element {
	e.JSValue().Set("innerHTML", html)
	return e
}

// HTML ...
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

// AddClass ...
func (e *Element) AddClass(names ...string) *Element {
	return e.clz("add", names...)
}

// RemoveClass ...
func (e *Element) RemoveClass(names ...string) *Element {
	return e.clz("remove", names...)
}

// ToggleClass ...
func (e *Element) ToggleClass(name string) *Element {
	return e.clz("toggle", name)
}

// ReplaceClass ...
func (e *Element) ReplaceClass(oldName, newName string) *Element {
	return e.clz("replace", oldName, newName)
}

// HasClass ...
func (e *Element) HasClass(name string) bool {
	return e.JSValue().Get("classList").Call("contains").Bool()
}

// ----------------------------------------------------------------------------

// TagName ...
func (e *Element) TagName() string {
	return strings.ToLower(e.JSValue().Get("tagName").String())
}

// ----------------------------------------------------------------------------

// Val ...
func (e *Element) Val() string {

	switch e.TagName() {
	case "input", "select":
		return e.JSValue().Get("value").String()
	case "textarea":
		return e.JSValue().Get("innerText").String()
	}
	return gojs.Undefined().String()
}

// SetVal ...
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

// Append ...
func (e *Element) Append(child interface{}) *Element {
	e.JSValue().Call("append", child)
	return e
}

// Clone https://developer.mozilla.org/en-US/docs/Web/API/Node/cloneNode
func (e *Element) Clone() *Element {
	return ElementOf(e.JSValue().Call("cloneNode", true))
}

// ----------------------------------------------------------------------------

// On ...
func (e *Element) On(event string, fn func(*Element, *js.Event)) *Element {
	cb := js.FuncOf(func(_this js.Value, args []js.Value) interface{} {
		fn(ElementOf(_this), js.EventOf(args[0]))
		return nil
	})

	e.EventTarget.On(event, cb)
	return e
}

// OnClick ...
func (e *Element) OnClick(fn func(*Element, *js.Event)) *Element {
	return e.On("click", fn)
}

// OnChange ...
func (e *Element) OnChange(fn func(*Element, *js.Event)) *Element {
	return e.On("change", fn)
}

// ----------------------------------------------------------------------------

// Call ...
func (e *Element) Call(name string, args ...interface{}) js.Value {
	return e.JSValue().Call(name, args...)
}
