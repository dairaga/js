package bs

import (
	"github.com/dairaga/js"
)

// Modal ...
type Modal struct {
	Object
}

// AttachModal ...
func AttachModal(id string) Modal {
	return Modal{Attach(id)}
}

// ----------------------------------------------------------------------------

// SetTitle ...
func (m Modal) SetTitle(title string) Modal {
	x := m.S(".modal-title")
	if x.Truthy() {
		x.SetText(title)
	}
	return m
}

// SetText ...
func (m Modal) SetText(content string) Modal {
	x := m.S(".modal-body")
	if x.Truthy() {
		x.SetText(content)
	}

	return m
}

// SetHTML ....
func (m Modal) SetHTML(html string) Modal {
	x := m.S(".modal-body")
	if x.Truthy() {
		x.SetHTML(html)
	}

	return m
}

// ----------------------------------------------------------------------------

// Show ...
func (m Modal) Show() {
	js.Call("$", m.id).Call("modal", "show")
}

// Hide ...
func (m Modal) Hide() {
	js.Call("$", m.id).Call("modal", "hide")
}

// Dispose ...
func (m Modal) Dispose() {
	js.Call("$", m.id).Call("modal", "dispose")
}

// ----------------------------------------------------------------------------

// Showing ...
func (m Modal) Showing(fn func(Modal, js.Event)) Modal {
	m.On("show.bs.modal", func(_ Object, e js.Event) {
		fn(m, e)
	})
	return m
}

// Shown ...
func (m Modal) Shown(fn func(Modal, js.Event)) Modal {
	m.On("shown.bs.modal", func(_ Object, e js.Event) {
		fn(m, e)
	})
	return m
}

// Hidding ...
func (m Modal) Hidding(fn func(Modal, js.Event)) Modal {
	m.On("hide.bs.modal", func(_ Object, e js.Event) {
		fn(m, e)
	})
	return m
}

// Hidden ...
func (m Modal) Hidden(fn func(Modal, js.Event)) Modal {
	m.On("hidden.bs.modal", func(_ Object, e js.Event) {
		fn(m, e)
	})
	return m
}
