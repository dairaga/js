package flex

import (
	"github.com/dairaga/js/dom"
	"github.com/dairaga/js/dom/bs"
)

func applyNoVS(elm *dom.Element, box bool) *dom.Element {
	if box {
		elm.AddClass("d-flex")
	} else {
		elm.AddClass("d-inline-flex")
	}
	return elm
}

func apply(elm *dom.Element, box bool, vs ...bs.ViewportSize) *dom.Element {
	if len(vs) <= 0 {
		return applyNoVS(elm, box)
	}

	s := vs[0]
	if s == bs.ExtraSmall {
		return applyNoVS(elm, box)
	}

	if box {
		elm.AddClass("d-" + s + "-flex")
	} else {
		elm.AddClass("d-" + s + "-inline-flex")
	}

	return elm
}

// Apply applies Bootstrap d-flex.
func Apply(elm *dom.Element, vs ...bs.ViewportSize) *dom.Element {
	return apply(elm, true, vs...)
}
