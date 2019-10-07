package bs

import (
	"fmt"

	"github.com/dairaga/js/dom"
)

// pre-defined viewport size values.
const (
	VSExtraSmall = ""
	VSSmall      = "sm"
	VSMedium     = "md"
	VSLarge      = "lg"
	VSExtraLarge = "xl"
)

// pre-defined position values.
const (
	SideLeft   = "left"
	SideRight  = "right"
	SideTop    = "top"
	SideBottom = "bottom"

	SideCenter = "center"

	SideNone = "none"

	// margin, padding
	SideNegative = "n"
	SideBlank    = ""
	SideAll      = ""
	SideX        = "x"
	SideY        = "y"

	// vertical alignment
	SideMiddle     = "middle"
	SideBaseline   = "baseline"
	SideTextTop    = "text-top"
	SideTextBottom = "text-bottom"

	// flex justify content.
	SideStart   = "start"
	SideEnd     = "end"
	SideBetween = "between"
	SideAround  = "around"
	SideStretch = "stretch"
)

// pre-defined bootstrap size values.
const (
	// for margin and padding
	Size0 = "0"
	Size1 = "1"
	Size2 = "2"
	Size3 = "3"
	Size4 = "4"
	Size5 = "5"

	// for width and height
	Size25   = "25"
	Size50   = "50"
	Size75   = "75"
	Size100  = "100"
	SizeAuto = "auto"
)

// pre-defined values.
const (
	FGPrimary   = "text-primary"
	FGSecondary = "text-secondary"
	FGSuccess   = "text-success"
	FGDanger    = "text-danger"
	FGWarning   = "text-warning"
	FGInfo      = "text-info"
	FGLight     = "text-light"
	FGDark      = "text-dark"
	FGBody      = "text-body"
	FGMuted     = "text-muted"
	FGWhite     = "text-white"
	FGHalfBlack = "text-black-50"
	FGHalfWhite = "text-white-50"

	BGPrimary     = "bg-primary"
	BGSecondary   = "bg-secondary"
	BGSuccess     = "bg-success"
	BGDanger      = "bg-danger"
	BGWarning     = "bg-warning"
	BGInfo        = "bg-info"
	BGLight       = "bg-light"
	BGDark        = "bg-dark"
	BGWhite       = "bg-white"
	BGTransparent = "bg-transparent"
)

// ----------------------------------------------------------------------------

// Property-Side-ViewportSize.
func css1(property string, side string, vs ...string) string {
	if len(vs) > 0 {
		return fmt.Sprintf("%s-%s-%s", property, side, vs[0])
	}
	return fmt.Sprintf("%s-%s", property, side)
}

// PropertySide-Size-ViewportSize
func ccs2(property string, side string, size string, vs ...string) string {
	if side != SideBlank {
		side = side[0:1]
	}

	if len(vs) > 0 {
		return fmt.Sprintf("%s%s-%s-%s", property, side, size, vs[0])
	}
	return fmt.Sprintf("%s%s-%s", property, side, size)
}

// ----------------------------------------------------------------------------

// TextAlign add text alignment style to element.
func TextAlign(elm *dom.Element, side string, vs ...string) *dom.Element {
	return elm.AddClass(css1("text", side, vs...))
}

// TextAlign add text alignment style to component.
func (obj *Component) TextAlign(side string, vs ...string) *Component {
	TextAlign(obj.Element, side, vs...)
	return obj
}

// RemoveTextAlign removes text alignment style from element.
func RemoveTextAlign(elm *dom.Element, side string, vs ...string) *dom.Element {
	return elm.RemoveClass(css1("text", side, vs...))
}

// RemoveTextAlign removes text alignment style from component.
func (obj *Component) RemoveTextAlign(side string, vs ...string) *Component {
	RemoveTextAlign(obj.Element, side, vs...)
	return obj
}

// ----------------------------------------------------------------------------

// VerticalAlign adds vertical alignment style to inline, inline-block, inline-table, and table cell elements.
func VerticalAlign(elm *dom.Element, side string) *dom.Element {
	return elm.AddClass("align-" + side)
}

// VerticalAlign adds vertical alignment style to inline, inline-block, inline-table, and table cell elements.
func (obj *Component) VerticalAlign(side string) *Component {
	VerticalAlign(obj.Element, side)
	return obj
}

// RemoveVerticalAlign removes vertical alignment style to inline, inline-block, inline-table, and table cell elements.
func RemoveVerticalAlign(elm *dom.Element, side string) *dom.Element {
	return elm.RemoveClass("align-" + side)
}

// RemoveVerticalAlign removes vertical alignment style to inline, inline-block, inline-table, and table cell elements.
func (obj *Component) RemoveVerticalAlign(side string) *Component {
	RemoveVerticalAlign(obj.Element, side)
	return obj
}

// ----------------------------------------------------------------------------

// Float adds float style to element.
func Float(elm *dom.Element, side string, vs ...string) *dom.Element {
	return elm.AddClass(css1("float", side, vs...))
}

// Float adds float style to component.
func (obj *Component) Float(side string, vs ...string) *Component {
	Float(obj.Element, side, vs...)
	return obj
}

// RemoveFloat removes float style from element.
func RemoveFloat(elm *dom.Element, side string, vs ...string) *dom.Element {
	return elm.RemoveClass(css1("float", side, vs...))
}

// RemoveFloat removes float style from component.
func (obj *Component) RemoveFloat(side string, vs ...string) *Component {
	RemoveFloat(obj.Element, side, vs...)
	return obj
}

// ----------------------------------------------------------------------------

// Color adds foreground, backgroup or gradient colors to element.
func Color(elm *dom.Element, color ...string) *dom.Element {
	return elm.AddClass(color...)
}

// Color adds foreground, backgroup or gradient colors to component.
func (obj *Component) Color(color ...string) *Component {
	Color(obj.Element, color...)
	return obj
}

// RemoveColor removes foreground, backgroup or gradient colors from element.
func RemoveColor(elm *dom.Element, color ...string) *dom.Element {
	return elm.RemoveClass(color...)
}

// RemoveColor removes foreground, backgroup or gradient colors from component.
func (obj *Component) RemoveColor(color ...string) *Component {
	RemoveColor(obj.Element, color...)
	return obj
}

// ----------------------------------------------------------------------------

// OnShow add callback function when showing.
func (obj *Component) OnShow(cb func(*Component)) {
	obj.show = cb
}

// OnHide add callback function when hiding.
func (obj *Component) OnHide(cb func(*Component)) {
	obj.hide = cb
}

// Show add .visible and remove .invisible to show component.
func (obj *Component) Show() *Component {
	obj.RemoveClass("invisible", "d-none").AddClass("visible")
	if obj.show != nil {
		obj.show(obj)
	}
	return obj
}

// Hide add .invisible and remove .visible to hide component.
func (obj *Component) Hide() *Component {
	obj.RemoveClass("visible").AddClass("invisible", "d-none")
	if obj.hide != nil {
		obj.hide(obj)
	}
	return obj
}
