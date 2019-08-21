package bs

import (
	"github.com/dairaga/js"
	"github.com/dairaga/js/dom"
)

// Component represents bootstrap component.
type Component struct {
	*dom.Element
	id string
}

// Attach binds some html component on page.
func Attach(id string) *Component {
	elm := dom.S(id)
	if !elm.Truthy() {
		panic("can not found " + id)
	}

	return &Component{id: id, Element: elm}
}

// ID returns component id attribute.
func (obj Component) ID() string {
	return obj.id
}

// On adds bootstrap listener.
func (obj *Component) On(event string, fn func(*Component, *js.Event)) {
	cb := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		fn(obj, js.EventOf(args[0]))
		return nil
	})
	obj.Register(event, cb)
	js.Call("$", obj.id).Call("on", event, cb)
}
