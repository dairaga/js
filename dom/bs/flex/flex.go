package flex

import (
	"fmt"

	"github.com/dairaga/js/dom"
	"github.com/dairaga/js/dom/bs"
)

func reversecss(property string, reverse bool, vs ...string) string {
	v := ""
	if len(vs) > 0 && vs[0] != bs.VSExtraSmall {
		v = "-" + vs[0]
	}
	dir := ""
	if reverse {
		dir = "-reverse"
	}

	return fmt.Sprintf("flex%s-%s%s", v, property, dir)

}

func rowcolumncss(row, reverse bool, vs ...string) string {
	property := "row"
	if !row {
		property = "column"
	}

	return reversecss(property, reverse, vs...)
}

func wrapcss(wrap, reverse bool, vs ...string) string {
	property := "wrap"
	if !wrap {
		property = "nowrap"
	}

	return reversecss(property, reverse, vs...)
}

func flexcss(property string, side string, vs ...string) string {
	v := ""
	if len(vs) > 0 || vs[0] != bs.VSExtraSmall {
		v = "-" + vs[0]
	}

	return fmt.Sprintf("%s%s-%s", property, v, side)
}

// ----------------------------------------------------------------------------

func applyNoVS(elm *dom.Element, box bool) *dom.Element {
	if box {
		elm.AddClass("d-flex")
	} else {
		elm.AddClass("d-inline-flex")
	}
	return elm
}

func apply(elm *dom.Element, box bool, vs ...string) *dom.Element {
	if len(vs) <= 0 {
		return applyNoVS(elm, box)
	}

	s := vs[0]
	if s == bs.VSExtraSmall {
		return applyNoVS(elm, box)
	}

	if box {
		elm.AddClass("d-" + s + "-flex")
	} else {
		elm.AddClass("d-" + s + "-inline-flex")
	}

	return elm
}

// Apply applies Bootstrap d-flex-*.
func Apply(elm *dom.Element, vs ...string) *dom.Element {
	return apply(elm, true, vs...)
}

// ApplyInline applies Bootstrap d-inline-flex-*.
func ApplyInline(elm *dom.Element, vs ...string) *dom.Element {
	return apply(elm, false, vs...)
}

// Column applies flex column style on element.
func Column(elm *dom.Element, reverse bool, vs ...string) *dom.Element {
	return elm.AddClass(rowcolumncss(false, reverse, vs...))
}

// Reverse applies flex reverse style on element.
func Reverse(elm *dom.Element, row bool, vs ...string) *dom.Element {
	return elm.AddClass(rowcolumncss(row, true, vs...))
}

// Justify applies flex justifing content style on element.
func Justify(elm *dom.Element, side string, vs ...string) *dom.Element {
	return elm.AddClass(flexcss("justify-content", side, vs...))
}

// AlignItems applies aligning items style on element.
func AlignItems(elm *dom.Element, side string, vs ...string) *dom.Element {
	return elm.AddClass(flexcss("align-items", side, vs...))
}

// Wrap add wrap style to element.
func Wrap(elm *dom.Element, reverse bool, vs ...string) *dom.Element {
	return elm.AddClass("d-flex", reversecss("wrap", reverse, vs...))
}

// NoWrap add nowrap style to element.
func NoWrap(elm *dom.Element, reverse bool, vs ...string) *dom.Element {
	return elm.AddClass("d-flex", reversecss("nowrap", reverse, vs...))
}

// AlignContent add align content style to element.
func AlignContent(elm *dom.Element, side string, vs ...string) *dom.Element {
	return elm.AddClass(flexcss("align-content", side, vs...))
}

// ChildAlign applies aligning self styel on child element in flexbox.
func ChildAlign(child *dom.Element, side string, vs ...string) *dom.Element {
	return child.AddClass(flexcss("align-self", side, vs...))
}

// ChildFill applies fill style on child element in flexbox.
func ChildFill(child *dom.Element, vs ...string) *dom.Element {
	return child.AddClass(flexcss("flex", "fill", vs...))
}

// ----------------------------------------------------------------------------

// New returns flex box.
func New(vs ...string) *bs.Component {
	f := bs.ComponentOf(dom.CreateElement("div"))
	Apply(f.Element, vs...)

	return f
}

// NewInline returns flex inline box.
func NewInline(vs ...string) *bs.Component {
	f := bs.ComponentOf(dom.CreateElement("div"))
	ApplyInline(f.Element, vs...)
	return f
}
