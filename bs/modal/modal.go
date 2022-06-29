//go:build js && wasm

package modal

import (
	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/bs"
)

var object = bs.Global().Get("Modal")

// -----------------------------------------------------------------------------

type Class struct {
	js.Element
	val js.Value
}

// -----------------------------------------------------------------------------

func (m *Class) Show() {
	m.val.Call("show")
}

// -----------------------------------------------------------------------------

func (m *Class) Hide() {
	m.val.Call("hide")
}

// -----------------------------------------------------------------------------

func (m *Class) Toggle() {
	m.val.Call("toggle")
}

// -----------------------------------------------------------------------------

func (m *Class) Dispose() {
	m.val.Call("dispose")
}

// -----------------------------------------------------------------------------

func (m *Class) onEvt(event string, fn func(*Class, js.Event)) {
	m.On(event, func(this js.Element, evt js.Event) {
		fn(m, evt)
	})
}

// -----------------------------------------------------------------------------

func (m *Class) OnShow(fn func(*Class, js.Event)) {
	m.onEvt("show.bs.modal", fn)
}

// -----------------------------------------------------------------------------

func (m *Class) OnShown(fn func(*Class, js.Event)) {
	m.onEvt("shown.bs.modal", fn)
}

// -----------------------------------------------------------------------------

func (m *Class) OnHide(fn func(*Class, js.Event)) {
	m.onEvt("hide.bs.modal", fn)
}

// -----------------------------------------------------------------------------

func (m *Class) OnHidden(fn func(*Class, js.Event)) {
	m.onEvt("hidden.bs.modal", fn)
}

// -----------------------------------------------------------------------------

type Option struct {
	Backdrop bool
	Keyboard bool
	Focus    bool
}

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
