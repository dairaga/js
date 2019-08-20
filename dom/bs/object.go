package bs

import (
	"github.com/dairaga/js"
	"github.com/dairaga/js/dom"
)

// Object ...
type Object struct {
	*dom.Element
	id string
}

// Attach ...
func Attach(id string) *Object {
	elm := dom.S(id)
	if !elm.Truthy() {
		panic("can not found " + id)
	}

	return &Object{id: id, Element: elm}
}

// On ...
func (obj *Object) On(event string, fn func(*Object, *js.Event)) {
	cb := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		fn(obj, js.EventOf(args[0]))
		return nil
	})
	obj.Register(event, cb)
	js.Call("$", obj.id).Call("on", event, cb)
}
