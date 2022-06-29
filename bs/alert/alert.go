//go:build js && wasm

package alert

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/bs"
)

var object = bs.Global().Get("Alert")

// -----------------------------------------------------------------------------

type Class struct {
	js.Element
	val js.Value
}

// -----------------------------------------------------------------------------

func (a *Class) onEvt(event string, fn func(*Class, js.Event)) {
	a.On(event, func(_ js.Element, evt js.Event) {
		fn(a, evt)
	})
}

// -----------------------------------------------------------------------------

func (a *Class) Close() {
	a.val.Call("close")
}

// -----------------------------------------------------------------------------

func (a *Class) Dispose() {
	a.val.Call("dispose")
}

// -----------------------------------------------------------------------------

func (a *Class) OnClose(fn func(*Class, js.Event)) {
	a.onEvt("close.bs.alert", fn)
}

// -----------------------------------------------------------------------------

func (a *Class) OnClosed(fn func(*Class, js.Event)) {
	a.onEvt("closed.bs.alert", fn)
}

// -----------------------------------------------------------------------------

func New(x any) *Class {
	elm := js.ElementOf(x)
	return &Class{
		Element: elm,
		val:     object.New(elm.JSValue()),
	}
}

// -----------------------------------------------------------------------------

func From(html js.HTML) *Class {
	tmpl := js.CreateTemplate(html)
	return New(tmpl.First())
}
