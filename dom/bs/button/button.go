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

	btn.Component.EventTarget.On("click", cb)
	return btn
}
