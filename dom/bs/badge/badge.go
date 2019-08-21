package badge

import (
	"github.com/dairaga/js/dom"
	"github.com/dairaga/js/dom/bs"
)

// Badge represents Bootstrap badge component.
type Badge struct {
	*bs.Component
}

// Attach binds Bootstrap badge component.
func Attach(id string) *Badge {
	return &Badge{bs.Attach(id)}
}

func generate(tag string, style bs.Style, content interface{}) *Badge {
	b := &Badge{bs.ComponentOf(dom.CreateElement(tag), "")}

	b.AddClass("badge", "badge-"+style)
	b.Append(content)

	return b
}

// New returns a Bootstrap badge with span tag.
func New(style bs.Style, content interface{}) *Badge {
	return generate("span", style, content)
}

// Link returns a Bootstrap badge with link (<a>).
func Link(style bs.Style, content interface{}) *Badge {
	return generate("a", style, content)
}

// Pill adds pill style.
func (b *Badge) Pill() *Badge {
	b.AddClass("badge-pill")
	return b
}
