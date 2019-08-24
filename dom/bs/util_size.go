package bs

import "github.com/dairaga/js/dom"

func wh(elm *dom.Element, property string, size string) *dom.Element {
	return elm.AddClass(property + size)
}

func rmwh(elm *dom.Element, property string) *dom.Element {
	return elm.RemoveClass(property+Size25, property+Size50, property+Size75, property+Size100, property+SizeAuto)
}

// Width adds width style to element.
func Width(elm *dom.Element, size string) *dom.Element {
	return wh(elm, "w-", size)
}

// Width adds width style to component.
func (obj *Component) Width(size string) *Component {
	Width(obj.Element, size)
	return obj
}

// RemoveWidth removes width style from element.
func RemoveWidth(elm *dom.Element) *dom.Element {
	rmwh(elm, "w-")
	return rmwh(elm, "mw-")
}

// RemoveWidth removes width style from component.
func (obj *Component) RemoveWidth() *Component {
	RemoveWidth(obj.Element)
	return obj
}

// Height adds height style to element.
func Height(elm *dom.Element, size string) *dom.Element {
	return wh(elm, "h-", size)
}

// Height adds height style to component.
func (obj *Component) Height(size string) *Component {
	Height(obj.Element, size)
	return obj
}

// RemoveHeight removes height style.
func RemoveHeight(elm *dom.Element) *dom.Element {
	rmwh(elm, "h-")
	return rmwh(elm, "mh-")
}

// RemoveHeight removes height style from component.
func (obj *Component) RemoveHeight() *Component {
	RemoveHeight(obj.Element)
	return obj
}
