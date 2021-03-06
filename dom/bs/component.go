package bs

import (
	"github.com/dairaga/js"
	"github.com/dairaga/js/dom"
)

// Component represents Bootstrap component.
type Component struct {
	*dom.Element
	show func(*Component)
	hide func(*Component)
}

// ComponentOf returns a Bootstrap component.
func ComponentOf(elm *dom.Element) *Component {
	return &Component{elm, nil, nil}
}

// Attach binds some html component on page.
func Attach(id string) *Component {
	elm := dom.S(id)
	if !elm.Truthy() {
		panic("can not found " + id)
	}

	return &Component{Element: elm}
}

// ----------------------------------------------------------------------------

// On adds Bootstrap listener.
func (obj *Component) On(event string, fn func(*js.Event)) {
	cb := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		fn(js.EventOf(args[0]))
		return nil
	})
	obj.Register(event, cb)
	js.Call("$", "#"+obj.Attr("id")).Call("on", event, cb)
}
