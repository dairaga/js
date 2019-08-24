/*Package spinner wraps Bootstrap spinner component.*/
package spinner

import (
	"github.com/dairaga/js/dom"
	"github.com/dairaga/js/dom/bs"
)

// Spinner represents Bootstrap spinner component.
type Spinner struct {
	*bs.Component
}

// Attach binds a Bootstrap spinner on page.
func Attach(id string) *Spinner {
	return &Spinner{bs.Attach(id)}
}

func generate(typ string, fgColor string) *Spinner {
	s := &Spinner{bs.ComponentOf(dom.CreateElement("div"))}
	s.Color(fgColor).
		AddClass("spinner-"+typ).
		SetAttr("role", "status").
		SetHTML(`<span class="sr-only">Loading...</span>`)
	return s
}

// Border returns a border spinner.
func Border(fgColor string) *Spinner {
	return generate("border", fgColor)
}

// Grow returns grow spinner.
func Grow(fgColor string) *Spinner {
	return generate("grow", fgColor)
}

// Smaller makes a smaller spinner.
func (s *Spinner) Smaller() *Spinner {
	if s.HasClass("spinner-grow") {
		s.AddClass("spinner-grow-sm")
	} else if s.HasClass("spinner-border") {
		s.AddClass("spinner-border-sm")
	}
	return s
}
