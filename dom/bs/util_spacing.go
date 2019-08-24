package bs

import "github.com/dairaga/js/dom"

// Margin adds margin style to element.
func Margin(elm *dom.Element, side string, size string, vs ...string) *dom.Element {
	return elm.AddClass(ccs2("m", side, size, vs...))
}

// Margin adds margin style to component.
func (obj *Component) Margin(side string, size string, vs ...string) *Component {
	Margin(obj.Element, side, size, vs...)
	return obj
}

// RemoveMargin removes margin style from element.
func RemoveMargin(elm *dom.Element, side string, size string, vs ...string) *dom.Element {
	return elm.RemoveClass(ccs2("m", side, size, vs...))
}

// RemoveMargin removes margin style from component.
func (obj *Component) RemoveMargin(side string, size string, vs ...string) *Component {
	RemoveMargin(obj.Element, side, size, vs...)
	return obj
}

// NegativeMargin add negative margin style to element.
func NegativeMargin(elm *dom.Element, side string, size string, vs ...string) *dom.Element {
	return Margin(elm, side, SideNegative+size, vs...)
}

// NegativeMargin add negative margin style to component.
func (obj *Component) NegativeMargin(side string, size string, vs ...string) *Component {
	NegativeMargin(obj.Element, side, size, vs...)
	return obj
}

// RemoveNegativeMargin removes negative margin style from element.
func RemoveNegativeMargin(elm *dom.Element, side string, size string, vs ...string) *dom.Element {
	return RemoveMargin(elm, side, SideNegative+size, vs...)
}

// RemoveNegativeMargin removes negative margin style from component.
func (obj *Component) RemoveNegativeMargin(side string, size string, vs ...string) *Component {
	RemoveNegativeMargin(obj.Element, side, size, vs...)
	return obj
}

// Center centers fixed-width element horizontally.
func Center(elm *dom.Element) *dom.Element {
	return elm.AddClass("mx-auto")
}

// Center centers fixed-width component horizontally.
func (obj *Component) Center() *Component {
	Center(obj.Element)
	return obj
}

// RemoveCenter removes horizontally centering style from element.
func RemoveCenter(elm *dom.Element) *dom.Element {
	return elm.RemoveClass("mx-auto")
}

// RemoveCenter removes horizontally centering style from component.
func (obj *Component) RemoveCenter() *Component {
	RemoveCenter(obj.Element)
	return obj
}

// ----------------------------------------------------------------------------

// Padding add padding style to element.
func Padding(elm *dom.Element, side string, size string, vs ...string) *dom.Element {
	return elm.AddClass(ccs2("p", side, size, vs...))
}

// Padding add padding style to component.
func (obj *Component) Padding(side string, size string, vs ...string) *Component {
	Padding(obj.Element, side, size, vs...)
	return obj
}

// RemovePadding removes paddind style from element.
func RemovePadding(elm *dom.Element, side string, size string, vs ...string) *dom.Element {
	return elm.RemoveClass(ccs2("p", side, size, vs...))
}

// RemovePadding removes paddind style from component.
func (obj *Component) RemovePadding(side string, size string, vs ...string) *Component {
	obj.RemoveClass(ccs2("p", side, size, vs...))
	return obj
}
