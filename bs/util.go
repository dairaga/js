//go:build js && wasm

package bs

import "github.com/dairaga/js/v3"

func Show(elm js.Element, show bool) js.Element {
	if show {
		elm.Remove("d-none")
	} else {
		elm.Add("d-none")
	}
	return elm
}

// -----------------------------------------------------------------------------

func Hide(elm js.Element) js.Element {
	elm.Add("d-none")
	return elm
}

// -----------------------------------------------------------------------------

func Visible(elm js.Element, v bool) js.Element {
	if v {
		elm.Remove("invisible")
	} else {
		elm.Add("invisible")
	}
	return elm
}

// -----------------------------------------------------------------------------

func Invisible(elm js.Element) js.Element {
	elm.Add("invisible")
	return elm
}
