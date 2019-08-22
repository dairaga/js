package bs

import "fmt"

// BootStrap utilities

// ViewportSize represents bootstrap viewport size value.
type ViewportSize = string

// pre-defined viewport size values.
const (
	ExtraSmall ViewportSize = ""
	Small      ViewportSize = "sm"
	Medium     ViewportSize = "md"
	Large      ViewportSize = "lg"
	ExtraLarge ViewportSize = "xl"
)

// Side represents css position for alignment, border and etc.
type Side = string

// pre-defined position values.
const (
	Left   Side = "left"
	Right  Side = "right"
	Center Side = "center"

	Top    Side = "top"
	Bottom Side = "bottom"

	None Side = "none"

	// margin, padding
	Negative      = "n"
	Blank    Side = ""
	X        Side = "x"
	Y        Side = "y"

	// vertical alignment
	Middle     = "middle"
	Baseline   = "baseline"
	TextTop    = "text-top"
	TextBottom = "text-bottom"
)

// Size represents bootstrap size values.
type Size = string

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

// Color represents bootstrap color value.
type Color = string

// pre-defined color values.
const (
	FGPrimary   Color = "text-primary"
	FGSecondary Color = "text-secondary"
	FGSuccess   Color = "text-success"
	FGDanger    Color = "text-danger"
	FGWarning   Color = "text-warning"
	FGInfo      Color = "text-info"
	FGLight     Color = "text-light"
	FGDark      Color = "text-dark"
	FGBody      Color = "text-body"
	FGMuted     Color = "text-muted"
	FGWhite     Color = "text-white"
	FGHalfBlack Color = "text-black-50"
	FGHalfWhite Color = "text-white-50"

	BGPrimary     Color = "bg-primary"
	BGSecondary   Color = "bg-secondary"
	BGSuccess     Color = "bg-success"
	BGDanger      Color = "bg-danger"
	BGWarning     Color = "bg-warning"
	BGInfo        Color = "bg-info"
	BGLight       Color = "bg-light"
	BGDark        Color = "bg-dark"
	BGWhite       Color = "bg-white"
	BGTransparent Color = "bg-transparent"
)

// ----------------------------------------------------------------------------

func css1(property string, side Side, vs ...ViewportSize) string {
	if len(vs) > 0 {
		return fmt.Sprintf("%s-%s-%s", property, side, vs[0])
	}
	return fmt.Sprintf("%s-%s", property, side)
}

func ccs2(property string, side Side, size Size, vs ...ViewportSize) string {
	if side != Blank {
		side = side[0:1]
	}

	if len(vs) > 0 {
		return fmt.Sprintf("%s%s-%s-%s", property, side, size, vs[0])
	}
	return fmt.Sprintf("%s%s-%s", property, side, size)
}

// ----------------------------------------------------------------------------

func (obj *Component) wh(property string, size Size) *Component {
	obj.AddClass(property + size)
	return obj
}

func (obj *Component) rmwh(property string) *Component {
	obj.RemoveClass(property+Size25, property+Size50, property+Size75, property+Size100, property+SizeAuto)
	return obj
}

// Width adds width style to component.
func (obj *Component) Width(size Size) *Component {
	return obj.wh("w-", size)
}

// RemoveWidth removes width style from component.
func (obj *Component) RemoveWidth() *Component {
	return obj.rmwh("w-").rmwh("mw-")
}

// Height adds height style to component.
func (obj *Component) Height(size Size) *Component {
	return obj.wh("h-", size)
}

// RemoveHeight removes height style from component.
func (obj *Component) RemoveHeight() *Component {
	return obj.rmwh("h-").rmwh("mh-")
}

// ----------------------------------------------------------------------------

// Margin adds margin style to component.
func (obj *Component) Margin(side Side, size Size, vs ...ViewportSize) *Component {
	obj.AddClass(ccs2("m", side, size, vs...))
	return obj
}

// RemoveMargin removes margin style from component.
func (obj *Component) RemoveMargin(side Side, size Size, vs ...ViewportSize) *Component {
	obj.RemoveClass(ccs2("m", side, size, vs...))
	return obj
}

// NegativeMargin add negative margin style to component.
func (obj *Component) NegativeMargin(side Side, size Size, vs ...ViewportSize) *Component {
	return obj.Margin(side, Negative+size, vs...)
}

// RemoveNegativeMargin removes negative margin style from component.
func (obj *Component) RemoveNegativeMargin(side Side, size Size, vs ...ViewportSize) *Component {
	return obj.RemoveMargin(side, Negative+size, vs...)
}

// Center centers fixed-width component horizontally.
func (obj *Component) Center() *Component {
	obj.AddClass("mx-auto")
	return obj
}

// RemoveCenter removes horizontally centering style from component.
func (obj *Component) RemoveCenter() *Component {
	obj.RemoveClass("mx-auto")
	return obj
}

// ----------------------------------------------------------------------------

// Padding add padding style to component.
func (obj *Component) Padding(side Side, size Size, vs ...ViewportSize) *Component {
	obj.AddClass(ccs2("p", side, size, vs...))
	return obj
}

// RemovePadding removes paddind style from component.
func (obj *Component) RemovePadding(side Side, size Size, vs ...ViewportSize) *Component {
	obj.RemoveClass(ccs2("p", side, size, vs...))
	return obj
}

// ----------------------------------------------------------------------------

// TextAlign add text alignment style to component.
func (obj *Component) TextAlign(side Side, vs ...ViewportSize) *Component {
	obj.AddClass(css1("text", side, vs...))
	return obj
}

// RemoveTextAlign removes text alignment style from component.
func (obj *Component) RemoveTextAlign(side Side, vs ...ViewportSize) *Component {
	obj.RemoveClass(css1("text", side, vs...))
	return obj
}

// ----------------------------------------------------------------------------

// VerticalAlign adds vertical alignment style to inline, inline-block, inline-table, and table cell elements.
func (obj *Component) VerticalAlign(side Side) *Component {
	obj.AddClass("align-" + side)
	return obj
}

// RemoveVerticalAlign removes vertical alignment style to inline, inline-block, inline-table, and table cell elements.
func (obj *Component) RemoveVerticalAlign(side Side) *Component {
	obj.RemoveClass("align-" + side)
	return obj
}

// ----------------------------------------------------------------------------

// Float adds float style to component.
func (obj *Component) Float(side Side, vs ...ViewportSize) *Component {
	obj.AddClass(css1("float", side, vs...))
	return obj
}

// RemoveFloat removes float style from component.
func (obj *Component) RemoveFloat(side Side, vs ...ViewportSize) *Component {
	obj.RemoveClass(css1("float", side, vs...))
	return obj
}

// ----------------------------------------------------------------------------

// Color adds foreground, backgroup or gradient colors to component.
func (obj *Component) Color(color ...Color) *Component {
	obj.AddClass(color...)
	return obj
}

// RemoveColor removes foreground, backgroup or gradient colors from component.
func (obj *Component) RemoveColor(color ...Color) *Component {
	obj.RemoveClass(color...)
	return obj
}

// ----------------------------------------------------------------------------

// Show add .visible and remove .invisible to show component.
func (obj *Component) Show() *Component {
	obj.RemoveClass("invisible").AddClass("visible")
	return obj
}

// Hide add .invisible and remove .visible to hide component.
func (obj *Component) Hide() *Component {
	obj.RemoveClass("visible").AddClass("invisible")
	return obj
}
