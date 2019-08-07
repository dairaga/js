package bs

import (
	"syscall/js"

	"github.com/dairaga/js/dom"
)

// Object ...
type Object struct {
	dom.Element
	id        string
	callbacks map[string]js.Func
}

// JSValue ...
func (obj Object) JSValue() js.Value {
	return obj.Element.JSValue()
}

// Attach ...
func Attach(id string) Object {
	elm := dom.S(id)
	if !elm.Truthy() {
		panic("can not found " + id)
	}

	return Object{id: id, Element: elm, callbacks: make(map[string]js.Func)}
}

// On ...
func (obj Object) On(event string, fn func(Object, dom.Event)) {
	cb := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		fn(obj, dom.EventOf(args[0]))
		return nil
	})

	old, ok := obj.callbacks[event]
	if ok {
		old.Release()
	}
	obj.callbacks[event] = cb
	dom.Call("$", obj.id).Call("on", event, cb)
}

/*
func fromHTML(html string) Object {
	p := js.CreateElement("div")
	p.SetHTML(html)
	js.AppendChild(p)

	x := p.Get("firstChild")
	p.Call("removeChild", x)
	js.RemoveChild(p)
	elm := js.ElementOf(x)
	js.AppendChild(elm)

	return Object{ref: elm}
}

func fromTmpl(tmpl string, data interface{}) Object {
	html := js.HTML(tmpl, data)
	return fromHTML(html)
}
*/
