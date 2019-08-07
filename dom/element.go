package dom

import (
	"strings"
	"syscall/js"
)

// Element ...
type Element struct {
	Value
}

// ElementOf ...
func ElementOf(x interface{}) Element {
	return Element{ValueOf(x)}
}

// ----------------------------------------------------------------------------

// S ...
func (e Element) S(selector string) Element {
	return ElementOf(e.ref.Call("querySelector", selector))
}

// SS ...
func (e Element) SS(selector string) NodeList {
	return NodeListOf(e.ref.Call("querySelectorAll", selector))
}

// ----------------------------------------------------------------------------

// Attr ...
func (e Element) Attr(name string) string {
	return e.call("getAttribute", name).String()
}

// SetAttr ...
func (e Element) SetAttr(name, value string) Element {
	e.call("setAttribute", name, value)
	return e
}

// ----------------------------------------------------------------------------

// Prop ...
func (e Element) Prop(name string) Value {
	return e.Get(name)
}

// SetProp ...
func (e Element) SetProp(name string, val interface{}) Element {
	e.Set(name, val)
	return e
}

// ----------------------------------------------------------------------------

// SetText ...
func (e Element) SetText(text string) Element {
	e.ref.Set("innerText", text)
	return e
}

// Text ...
func (e Element) Text() string {
	return e.ref.Get("innerText").String()
}

// SetHTML ...
func (e Element) SetHTML(html string) Element {
	e.ref.Set("innerHTML", html)
	return e
}

// HTML ...
func (e Element) HTML() string {
	return e.ref.Get("innerHTML").String()
}

// ----------------------------------------------------------------------------

func (e Element) clz(m string, args ...string) Element {
	size := len(args)
	if size <= 0 {
		return e
	}

	if size == 1 {
		e.ref.Get("classList").Call(m, args[0])
	} else if size == 2 {
		e.ref.Get("classList").Call(m, args[0], args[1])
	} else if size == 3 {
		e.ref.Get("classList").Call(m, args[0], args[1], args[2])
	} else {
		x := make([]interface{}, size)
		for i, str := range args {
			x[i] = str
		}
		e.ref.Get("classList").Call(m, x...)
	}

	return e
}

// AddClass ...
func (e Element) AddClass(names ...string) Element {
	return e.clz("add", names...)
}

// RemoveClass ...
func (e Element) RemoveClass(names ...string) Element {
	return e.clz("remove", names...)
}

// ToggleClass ...
func (e Element) ToggleClass(name string) Element {
	return e.clz("toggle", name)
}

// ReplaceClass ...
func (e Element) ReplaceClass(oldName, newName string) Element {
	return e.clz("replace", oldName, newName)
}

// HasClass ...
func (e Element) HasClass(name string) bool {
	return e.ref.Get("classList").Call("contains").Bool()
}

// ----------------------------------------------------------------------------

// TagName ...
func (e Element) TagName() string {
	return strings.ToLower(e.ref.Get("tagName").String())
}

// ----------------------------------------------------------------------------

// Val ...
func (e Element) Val() string {

	switch e.TagName() {
	case "input", "select":
		return e.Get("value").String()
	case "textarea":
		return e.Get("innerText").String()
	}
	return undefined.String()
}

// SetVal ...
func (e Element) SetVal(val interface{}) Element {
	switch e.TagName() {
	case "input", "select":
		e.Set("value", val)
	case "textarea":
		e.Set("innerText", val)
	}
	return e
}

// ----------------------------------------------------------------------------

// Append ...
func (e Element) Append(child interface{}) Element {
	e.ref.Call("append", child)
	return e
}

// ----------------------------------------------------------------------------

// On ...
func (e Element) On(event string, fn func(Element, Event)) Element {
	cb := js.FuncOf(func(_this js.Value, args []js.Value) interface{} {
		fn(ElementOf(_this), EventOf(args[0]))
		return nil
	})

	e.call("addEventListener", event, cb)
	return e
}

// OnClick ...
func (e Element) OnClick(fn func(Element, Event)) Element {
	return e.On("click", fn)
}

// OnChange ...
func (e Element) OnChange(fn func(Element, Event)) Element {
	return e.On("change", fn)
}

// ----------------------------------------------------------------------------

// Call ...
func (e Element) Call(name string, args ...interface{}) Value {
	return ValueOf(e.ref.Call(name, args...))
}
