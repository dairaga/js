// +build js,wasm

package toast

import (
	"syscall/js"

	"github.com/dairaga/js/dom"
	"github.com/dairaga/js/dom/bs"
)

// Toast represents bootstrap toast.
type Toast struct {
	*bs.Component
}

// New returns alert component.
func New(msg string, autoHide bool, header ...interface{}) *Toast {
	a := &Toast{bs.ComponentOf(dom.CreateElement("div"))}

	a.AddClass("toast").
		SetAttr("role", "alert").
		SetAttr("aria-live", "assertive").
		SetAttr("aria-atomic", "true")

	if autoHide {
		a.SetAttr("data-autohide", "true")
		a.SetAttr("data-delay", "3000")
	} else {
		a.SetAttr("data-autohide", "false")
	}

	h := dom.CreateElement("div")
	h.AddClass("toast-header")
	if len(header) > 0 {
		h.Append(header...)
	}

	closeBtn := dom.CreateElement("button")
	closeBtn.AddClass("ml-2", "mb-1", "close").SetAttr("type", "button").SetAttr("data-dismiss", "toast").SetAttr("aria-label", "Close")
	closeBtn.SetHTML(`<span aria-hidden="true">&times;</span>`)
	h.Append(closeBtn)

	a.Append(h)

	b := dom.CreateElement("div")
	b.AddClass("toast-body")
	b.SetText(msg)

	a.Append(b)
	return a
}

// Show ...
func Show(parent *bs.Component, msg string, autoHide bool, header ...interface{}) *Toast {
	a := New(msg, autoHide, header...)
	parent.Prepend(a)
	a.Show()
	return a
}

// Show ...
func (a *Toast) Show() *Toast {
	js.Global().Call("$", a.JSValue()).Call("toast", "show")
	return a
}

// Close hide and dispose alert component.
func (a *Toast) Close() *Toast {
	js.Global().Call("$", a.JSValue()).Call("toast", "hide")
	return a
}
