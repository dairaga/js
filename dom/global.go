package dom

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
	gojs "syscall/js"
)

var window = gojs.Global().Get("window")
var document = ElementOf(gojs.Global().Get("document"))
var undefined = ValueOf(gojs.Undefined())
var null = ValueOf(gojs.Null())

// ----------------------------------------------------------------------------

// New ...
func New(constructor string, args ...interface{}) Value {
	return ValueOf(gojs.Global().Get(constructor).New(args...))
}

// Call ...
func Call(fn string, args ...interface{}) Value {
	return ValueOf(gojs.Global().Call(fn, args...))
}

// Get ...
func Get(name string) Value {
	return ValueOf(gojs.Global().Get(name))
}

// ----------------------------------------------------------------------------

// Func ...
type Func func(Value, []Value) interface{}

// RegisterFunc ...
func RegisterFunc(name string, fn Func) {
	fx := gojs.FuncOf(func(_this gojs.Value, args []gojs.Value) interface{} {
		argx := make([]Value, len(args))
		for i, x := range args {
			argx[i] = ValueOf(x)
		}

		return fn(ValueOf(_this), argx)
	})

	gojs.Global().Set(name, fx)
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

// URL ...
func URL() (*url.URL, error) {
	return url.Parse(window.Get("location").Get("href").String())
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

	return strings.TrimSpace(sb.String())
}

// ----------------------------------------------------------------------------

// Cookies returns all cookies.
func Cookies() []*http.Cookie {
	rawcookie := document.Get("cookie").String()
	if rawcookie == "" {
		return nil
	}

	pairs := strings.Split(rawcookie, ";")
	if len(pairs) <= 0 {
		return nil
	}

	var result []*http.Cookie
	for _, x := range pairs {
		pair := strings.Split(x, "=")
		switch len(pair) {
		case 2:
			result = append(result, &http.Cookie{Name: strings.TrimSpace(pair[0]), Value: pair[1]})
		case 1:
			result = append(result, &http.Cookie{Name: strings.TrimSpace(pair[0]), Value: ""})
		default:
			continue
		}
	}

	return result
}

// Cookie returns cookie for name.
func Cookie(name string) *http.Cookie {
	rawcookie := document.Get("cookie").String()
	key := name + "="
	pos := strings.Index(rawcookie, key)
	if pos < 0 {
		return nil
	}

	pos2 := strings.Index(rawcookie[pos:], ";")
	if pos2 >= pos {
		return &http.Cookie{
			Name:  name,
			Value: rawcookie[pos+len(key) : pos2],
		}
	}
	return &http.Cookie{
		Name:  name,
		Value: rawcookie[pos+len(key):],
	}
}

// SetCookie set a new cookie.
func SetCookie(cookie *http.Cookie) {
	document.Set("cookie", cookie.String())
}
