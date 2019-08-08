package dom

import (
	"fmt"
	"html/template"
	"strings"
	"syscall/js"
)

var window = js.Global().Get("window")
var document = ElementOf(js.Global().Get("document"))
var undefined = ValueOf(js.Undefined())
var null = ValueOf(js.Null())

// ----------------------------------------------------------------------------

// New ...
func New(constructor string, args ...interface{}) Value {
	return ValueOf(js.Global().Get(constructor).New(args...))
}

// Call ...
func Call(fn string, args ...interface{}) Value {
	return ValueOf(js.Global().Call(fn, args...))
}

// Get ...
func Get(name string) Value {
	return ValueOf(js.Global().Get(name))
}

// ----------------------------------------------------------------------------

// Func ...
type Func func(Value, []Value) interface{}

// RegisterFunc ...
func RegisterFunc(name string, fn Func) {
	fx := js.FuncOf(func(_this js.Value, args []js.Value) interface{} {
		argx := make([]Value, len(args))
		for i, x := range args {
			argx[i] = ValueOf(x)
		}

		return fn(ValueOf(_this), argx)
	})

	js.Global().Set(name, fx)
}

// ----------------------------------------------------------------------------

// window

// Alert ...
func Alert(a ...interface{}) {
	window.Call("alert", fmt.Sprint(a...))
}

// Alertf ...
func Alertf(format string, a ...interface{}) {
	window.Call("alert", fmt.Sprintf(format, a...))
}

// Confirm ...
func Confirm(a ...interface{}) bool {
	return window.Call("confirm", fmt.Sprint(a...)).Bool()
}

// Confirmf ...
func Confirmf(format string, a ...interface{}) bool {
	return window.Call("confirm", fmt.Sprintf(format, a...)).Bool()
}

// ----------------------------------------------------------------------------

// document

// CreateElement ...
func CreateElement(tag string) Element {
	return ElementOf(document.ref.Call("createElement", tag))
}

// AppendChild ...
func AppendChild(child Element) {
	document.ref.Call("appendChild", child.ref)
}

// RemoveChild ...
func RemoveChild(child Element) {
	document.ref.Call("removeChild", child.ref)
}

// S ...
func S(selector string) Element {
	return document.S(selector)
}

// SS ...
func SS(selector string) NodeList {
	return document.SS(selector)
}

// ----------------------------------------------------------------------------

// HTML executes golang template.
func HTML(tmpl string, data interface{}) string {
	t, err := template.New("_dairaga_js_").Parse(tmpl)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	sb := &strings.Builder{}
	t.Execute(sb, data)

	return sb.String()
}
