//go:build js && wasm

// Package alert is a wrapper of Bootstrap Alert component.
package alert

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/bs"
)

// bootstrap alert component.
var object = bs.Global().Get("Alert")

// -----------------------------------------------------------------------------

// Class represents Bootstrap alert component.
// See https://getbootstrap.com/docs/5.2/components/alerts/
type Class struct {
	js.Element
	val js.Value // bootstrap alert instance.
}

// -----------------------------------------------------------------------------

// onEvt add an event listener to alert component.
func (a *Class) onEvt(event string, fn func(*Class, js.Event)) {
	a.On(event, func(_ js.Element, evt js.Event) {
		fn(a, evt)
	})
}

// -----------------------------------------------------------------------------

// Close closes the alert and removes it from the DOM.
func (a *Class) Close() {
	a.val.Call("close")
}

// -----------------------------------------------------------------------------

// Dispose destroys the alert and removes it from the DOM.
func (a *Class) Dispose() {
	a.val.Call("dispose")
}

// -----------------------------------------------------------------------------

// OnClose adds a listener fired immediately when the close instance method is called.
func (a *Class) OnClose(fn func(*Class, js.Event)) {
	a.onEvt("close.bs.alert", fn)
}

// -----------------------------------------------------------------------------

// OnClosed adds a listener fired when the alert has been closed and CSS transitions have completed.
func (a *Class) OnClosed(fn func(*Class, js.Event)) {
	a.onEvt("closed.bs.alert", fn)
}

// -----------------------------------------------------------------------------

// New returns a new alert component. Given x can be selector, HTML, or Element.
func New(x any) *Class {
	elm := js.ElementOf(x)
	return &Class{
		Element: elm,
		val:     object.New(elm.JSValue()),
	}
}

// -----------------------------------------------------------------------------

// From returns a new alert component from HTML.
func From(html js.HTML) *Class {
	tmpl := js.CreateTemplate(html)
	return New(tmpl.First())
}
