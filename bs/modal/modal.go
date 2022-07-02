//go:build js && wasm

// Package modal is a wrapper of Bootstrap Modal component.
package modal

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/bs"
)

var object = bs.Global().Get("Modal") // bootstrap modal component.

// -----------------------------------------------------------------------------

// Class represents Bootstrap modal component.
//
// See https://getbootstrap.com/docs/5.2/components/modal/
type Class struct {
	js.Element
	val js.Value // bootstrap modal instance.
}

// -----------------------------------------------------------------------------

// Show manually opens a modal.
func (m *Class) Show() {
	m.val.Call("show")
}

// -----------------------------------------------------------------------------

// Hide manually hides a modal.
func (m *Class) Hide() {
	m.val.Call("hide")
}

// -----------------------------------------------------------------------------

// Toggle manually toggles a modal.
func (m *Class) Toggle() {
	m.val.Call("toggle")
}

// -----------------------------------------------------------------------------

// Dispose destroys the modal and removes it from the DOM.
func (m *Class) Dispose() {
	m.val.Call("dispose")
}

// -----------------------------------------------------------------------------

// onEvt add an event listener to modal component.
func (m *Class) onEvt(event string, fn func(*Class, js.Event)) {
	m.On(event, func(this js.Element, evt js.Event) {
		fn(m, evt)
	})
}

// -----------------------------------------------------------------------------

// OnShow adds a listener fired immediately when the show instance method is called.
func (m *Class) OnShow(fn func(*Class, js.Event)) {
	m.onEvt("show.bs.modal", fn)
}

// -----------------------------------------------------------------------------

// OnShown adds a listener fired when the modal has been made visible to the user (will wait for CSS transitions to complete).
func (m *Class) OnShown(fn func(*Class, js.Event)) {
	m.onEvt("shown.bs.modal", fn)
}

// -----------------------------------------------------------------------------

// OnHide adds a listener fired immediately when the hide instance method has been called.
func (m *Class) OnHide(fn func(*Class, js.Event)) {
	m.onEvt("hide.bs.modal", fn)
}

// -----------------------------------------------------------------------------

// OnHidden adds a listener fired when the modal has been hidden from the user (will wait for CSS transitions to complete).
func (m *Class) OnHidden(fn func(*Class, js.Event)) {
	m.onEvt("hidden.bs.modal", fn)
}

// -----------------------------------------------------------------------------

// Option is modal attributes.
type Option struct {
	Backdrop bool // includes a modal-backdrop element.
	Keyboard bool // closes the modal when escape key is pressed.
	Focus    bool // puts the focus on the modal when initialized.
}

// default modal attributes.
var _default = &Option{
	Backdrop: true,
	Keyboard: true,
	Focus:    true,
}

// -----------------------------------------------------------------------------

func (opt *Option) JSValue() js.Value {
	return js.ValueOf(map[string]any{
		"backdrop": opt.Backdrop,
		"keyboard": opt.Keyboard,
		"focus":    opt.Focus,
	})
}

// -----------------------------------------------------------------------------

// New returns a new modal component.
func New(x any, opt ...*Option) *Class {
	elm := js.ElementOf(x)

	if len(opt) > 0 {
		return &Class{
			Element: elm,
			val:     object.New(elm.JSValue(), opt[0].JSValue()),
		}
	} else {
		return &Class{
			Element: elm,
			val:     object.New(elm.JSValue(), _default.JSValue()),
		}
	}
}
