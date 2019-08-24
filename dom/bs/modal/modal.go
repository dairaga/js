/*Package modal wraps Bootstrap Modal component.*/
package modal

import (
	"github.com/dairaga/js"
	"github.com/dairaga/js/dom/bs"
)

// Modal represents Bootstrap Modal component.
type Modal struct {
	*bs.Component
}

// Attach binds a Bootstrap modal component on page.
func Attach(id string) *Modal {
	return &Modal{bs.Attach(id)}
}

// ----------------------------------------------------------------------------

// SetTitle ...
func (m *Modal) SetTitle(title string) *Modal {
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
	js.Call("$", m.Attr("id")).Call("modal", "show")
}

// Hide ...
func (m Modal) Hide() {
	js.Call("$", m.Attr("id")).Call("modal", "hide")
}

// Dispose ...
func (m Modal) Dispose() {
	js.Call("$", m.Attr("id")).Call("modal", "dispose")
}

// ----------------------------------------------------------------------------

// Showing ...
func (m *Modal) Showing(fn func(*Modal, *js.Event)) *Modal {
	m.On("show.bs.modal", func(e *js.Event) {
		fn(m, e)
	})
	return m
}

// Shown ...
func (m *Modal) Shown(fn func(*Modal, *js.Event)) *Modal {
	m.On("shown.bs.modal", func(e *js.Event) {
		fn(m, e)
	})
	return m
}

// Hidding ...
func (m *Modal) Hidding(fn func(*Modal, *js.Event)) *Modal {
	m.On("hide.bs.modal", func(e *js.Event) {
		fn(m, e)
	})
	return m
}

// Hidden ...
func (m *Modal) Hidden(fn func(*Modal, *js.Event)) *Modal {
	m.On("hidden.bs.modal", func(e *js.Event) {
		fn(m, e)
	})
	return m
}
