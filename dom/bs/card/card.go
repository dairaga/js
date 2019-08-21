package card

import (
	"github.com/dairaga/js/dom"
	"github.com/dairaga/js/dom/bs"
)

// Header represents Bootstrap card header.
type Header struct {
	*bs.Component
}

// NewHeader returns card header.
func NewHeader(content interface{}) *Header {
	h := &Header{bs.ComponentOf(dom.CreateElement("div"))}
	h.AddClass("card-header")
	if content != nil {
		h.Append(content)
	}
	return h
}

// ----------------------------------------------------------------------------

// Body represents Bootstrap card body.
type Body struct {
	*bs.Component
}

// NewBody returns card body.
func NewBody(content interface{}) *Body {
	b := &Body{bs.ComponentOf(dom.CreateElement("div"))}
	b.AddClass("card-body")
	if content != nil {
		b.Append(content)
	}
	return b
}

// Title adds title.
func (b *Body) Title(content interface{}) *Body {
	t := dom.CreateElement("h5")
	t.AddClass("card-title").Append(content)
	b.Append(t)
	return b
}

// SubTitle adds subtitle.
func (b *Body) SubTitle(content interface{}) *Body {
	st := dom.CreateElement("h6")
	st.AddClass("card-subtitle").Append(content)
	b.Append(st)
	return b
}

// Link adds link.
func (b *Body) Link(link string, content interface{}) *Body {
	ln := dom.CreateElement("a")
	ln.SetAttr("href", link).AddClass("card-link")
	if content != nil {
		ln.Append(content)
	}
	b.Append(ln)
	return b
}

// Text adds text content.
func (b *Body) Text(content string) *Body {
	p := dom.CreateElement("p")
	p.AddClass("card-text").SetText(content)
	b.Append(p)
	return b
}

// ----------------------------------------------------------------------------

// Footer represents Bootstrap card footer.
type Footer struct {
	*bs.Component
}

// NewFooter returns card footer.
func NewFooter(content interface{}) *Footer {
	f := &Footer{bs.ComponentOf(dom.CreateElement("div"))}
	f.AddClass("card-footer")
	if content != nil {
		f.Append(content)
	}
	return f
}

// ----------------------------------------------------------------------------

// Card represents Bootstrap card component.
type Card struct {
	*bs.Component
}

// Attach binds card component on page.
func Attach(id string) *Card {
	return &Card{bs.Attach(id)}
}

// New returns card.
func New() *Card {
	c := &Card{bs.ComponentOf(dom.CreateElement("div"))}
	c.AddClass("card")
	return c
}

// ----------------------------------------------------------------------------

// TopImage adds top image.
func (c *Card) TopImage(link string, alt string) *Card {
	img := dom.CreateElement("img")
	img.SetAttr("src", link).SetAttr("alt", alt).AddClass("card-img-top")
	c.Append(img)
	return c
}

// Header adds card header.
func (c *Card) Header(h *Header) *Card {
	c.Append(h)
	return c
}

// Body adds body.
func (c *Card) Body(b *Body) *Card {
	c.Append(b)
	return c
}

// Footer adds card footer.
func (c *Card) Footer(f *Footer) *Card {
	c.Append(f)
	return c
}
