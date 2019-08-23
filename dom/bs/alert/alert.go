package alert

import (
	"github.com/dairaga/js/dom"
	"github.com/dairaga/js/dom/bs"
)

// Alert represents Bootstrap alert component.
type Alert struct {
	*bs.Component
}

// New returns alert component.
func New(style bs.Style, close bool, content ...interface{}) *Alert {
	a := &Alert{bs.ComponentOf(dom.CreateElement("div"))}

	a.AddClass("alert", "alert-"+style).SetAttr("role", "alert").Append(content...)

	if close {
		a.AddClass("alert-dismissible", "fade", "show")
		closeBtn := dom.CreateElement("button").
			SetAttr("type", "button").
			SetAttr("data-dismiss", "alert").
			SetAttr("aria-label", "Close").
			AddClass("close").SetHTML(`<span aria-hidden="true">&times;</span>`)

		a.Append(closeBtn)
	}

	return a
}
