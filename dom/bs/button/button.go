/*Package button wraps Bootstrap Button component. */
package button

import (
	"github.com/dairaga/js"
	"github.com/dairaga/js/dom"
	"github.com/dairaga/js/dom/bs"
)

// Button represents Bootstrap button component.
type Button struct {
	*bs.Component
}

// Attach binds a Bootstrap button on page
func Attach(id string) *Button {
	return &Button{bs.Attach(id)}
}

// New returns a button with <button> tag.
func New(style bs.Style, content ...interface{}) *Button {
	btn := &Button{bs.ComponentOf(dom.CreateElement("button"))}

	btn.AddClass("btn", "btn-"+style).Append(content...)
	return btn
}

// OnClick add click callback.
func (btn *Button) OnClick(fn func(*Button, *js.Event)) *Button {
	cb := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		fn(btn, js.EventOf(args[0]))
		return nil
	})

	btn.AddEventListener("click", cb)
	return btn
}
