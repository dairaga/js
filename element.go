//go:build js && wasm

package js

import (
	"fmt"

	"github.com/dairaga/js/v2/builtin"
	"github.com/dairaga/js/v2/mvvm"
)

type HTML string

func (h HTML) JSValue() Value {
	return ValueOf(string(h))
}

// -----------------------------------------------------------------------------

type Plain string

func (p Plain) JSValue() Value {
	return ValueOf(string(p))
}

// -----------------------------------------------------------------------------

func (p Plain) Ref() Value {
	return p.JSValue()
}

// -----------------------------------------------------------------------------

type Element interface {
	Appendable

	Tattoo() string

	Query(selector string) Element
	QueryAll(selector string) Elements

	Append(child Appendable, at ...string) Element
	Prepend(child Appendable, at ...string) Element

	Prop(p string, at ...string) Value
	SetProp(p string, val any, at ...string) Element

	Attr(a string, at ...string) string
	SetAttr(a, value string, at ...string) Element

	Text() string
	SetText(txt string) Element

	HTML() HTML
	SetHTML(html HTML) Element

	Value() string
	SetValue(val string) Element
	Files() []File

	Add(clz string, at ...string) Element
	Remove(clz string, at ...string) Element
	Has(clz string, at ...string) bool
	Replace(oldClz, newClz string, at ...string) Element
	Toggle(clz string, at ...string) Element

	On(typ string, fn func(sender Element), at ...string) Element
	OnClick(fn func(sender Element), at ...string) Element
	OnChange(fn func(sender Element), at ...string) Element

	Click(at ...string) Element
	Foucs(at ...string) Element
	Blur(at ...string) Element

	Empty() Element
	Relese()

	Trigger(name string, at ...string)
	Bind(name string, val *string, cb func(string, string))
}

// -----------------------------------------------------------------------------

type element Value

func (e element) JSValue() Value {
	return Value(e)
}

// -----------------------------------------------------------------------------

func (e element) Ref() Value {
	return e.JSValue()
}

// -----------------------------------------------------------------------------

func (e element) tattoo() element {
	val := Value(e).Call("getAttribute", _tattoo)
	if val.Truthy() {
		return e
	}
	Value(e).Call("setAttribute", _tattoo, tattoo(10))
	return e
}

// -----------------------------------------------------------------------------

func (e element) Tattoo() string {
	return e.Attr(_tattoo)
}

// -----------------------------------------------------------------------------

func (e element) at(a ...string) Value {
	if len(a) > 0 {
		v := Value(e).Call("querySelector", a[0])
		element(v).tattoo()
		return v
	}
	return Value(e)
}

// -----------------------------------------------------------------------------

func (e element) Query(selector string) Element {
	return elementOf(query(Value(e), selector))
}

// -----------------------------------------------------------------------------

func (e element) QueryAll(selector string) Elements {
	return ElementsOf(queryAll(Value(e), selector))
}

// -----------------------------------------------------------------------------

func (e element) Append(child Appendable, at ...string) Element {
	e.at(at...).Call("append", child.Ref())
	return e
}

// -----------------------------------------------------------------------------

func (e element) Prepend(child Appendable, at ...string) Element {
	e.at(at...).Call("prepend", child.Ref())
	return e
}

// -----------------------------------------------------------------------------

func (e element) Prop(p string, at ...string) Value {
	return e.at(at...).Get(p)
}

// -----------------------------------------------------------------------------

func (e element) SetProp(p string, val any, at ...string) Element {
	e.at(at...).Set(p, val)
	return e
}

// -----------------------------------------------------------------------------

func (e element) Attr(a string, at ...string) string {
	val := e.at(at...).Call("getAttribute", a)

	if val.Truthy() {
		return val.String()
	}
	return ""
}

// -----------------------------------------------------------------------------

func (e element) SetAttr(a, val string, at ...string) Element {
	if _tattoo != a {
		e.at(at...).Call("setAttribute", a, val)
	}
	return e
}

// -----------------------------------------------------------------------------

func (e element) Text() string {
	return Value(e).Get("innerText").String()
}

// -----------------------------------------------------------------------------

func (e element) SetText(txt string) Element {
	Value(e).Set("innerText", txt)
	return e
}

// -----------------------------------------------------------------------------

func (e element) HTML() HTML {
	return HTML(Value(e).Get("innerHTML").String())
}

// -----------------------------------------------------------------------------

func (e element) SetHTML(html HTML) Element {
	Value(e).Set("innerHTML", html.JSValue())
	return e
}

// -----------------------------------------------------------------------------

func (e element) Value() string {
	if builtin.HasValueProperty(Value(e)) {
		return Value(e).Get("value").String()
	}
	return ""
}

// -----------------------------------------------------------------------------

func (e element) SetValue(val string) Element {
	if builtin.HasValueProperty(Value(e)) {
		Value(e).Set("value", val)
	}
	return e
}

// -----------------------------------------------------------------------------

func (e element) Files() []File {
	if builtin.IsInputElement(Value(e)) && e.Attr("type") == "file" {
		lst := Value(e).Get("files")
		size := lst.Length()
		ret := make([]File, size)
		for i := 0; i < size; i++ {
			ret[i] = FileOf(lst.Index(i))
		}
	}
	return []File{}
}

// -----------------------------------------------------------------------------

func (e element) Add(clz string, at ...string) Element {
	e.Prop("classList", at...).Call("add", clz)
	return e
}

// -----------------------------------------------------------------------------

func (e element) Remove(clz string, at ...string) Element {
	e.Prop("classList", at...).Call("remove", clz)
	return e
}

// -----------------------------------------------------------------------------

func (e element) Has(clz string, at ...string) bool {
	return e.Prop("classList", at...).Call("contains", clz).Bool()
}

// -----------------------------------------------------------------------------

func (e element) Replace(oldClz, newClz string, at ...string) Element {
	e.Prop("classList", at...).Call("replace", oldClz, newClz)
	return e
}

// -----------------------------------------------------------------------------

func (e element) Toggle(clz string, at ...string) Element {
	e.Prop("classList", at...).Call("toggle", clz)
	return e
}

// -----------------------------------------------------------------------------

func (e element) On(typ string, fn func(sender Element), at ...string) Element {
	cb := FuncOf(func(_this Value, args []Value) any {
		elm := elementOf(args[0].Get("target"))
		fn(elm)
		return nil
	})

	e.at(at...).Call("addEventListener", typ, cb)
	return e
}

// -----------------------------------------------------------------------------

func (e element) OnClick(fn func(sender Element), at ...string) Element {
	return e.On("click", fn, at...)
}

// -----------------------------------------------------------------------------

func (e element) OnChange(fn func(sender Element), at ...string) Element {
	return e.On("change", fn, at...)
}

// -----------------------------------------------------------------------------

func (e element) Click(at ...string) Element {
	e.at(at...).Call("click")
	return e
}

// -----------------------------------------------------------------------------

func (e element) Foucs(at ...string) Element {
	e.at(at...).Call("focus")
	return e
}

// -----------------------------------------------------------------------------

func (e element) Blur(at ...string) Element {
	e.at(at...).Call("blur")
	return e
}

// -----------------------------------------------------------------------------

func (e element) Empty() Element {
	children := Value(e).Get("children")
	size := children.Length()

	for i := 0; i < size; i++ {
		children.Index(i).Call("remove")
	}
	return e
}

// -----------------------------------------------------------------------------

func (e element) Relese() {
	Value(e).Call("remove")
}

// -----------------------------------------------------------------------------

func (e element) Trigger(name string, at ...string) {
	v := e.at(at...)
	mvvm.Trigger(v.Call("getAttribute", _tattoo).String(), name)
}

// -----------------------------------------------------------------------------

func (e element) Bind(name string, val *string, cb func(string, string)) {
	mvvm.Add(name, val)
	mvvm.Watch(name, cb)
	e.OnChange(func(_ Element) {
		*val = e.Value()
		e.Trigger(name)
	})
}

// -----------------------------------------------------------------------------

func elementOf(v Value) element {
	if builtin.IsElement(v) {
		return element(v).tattoo()
	}
	panic(fmt.Sprintf("%s is not an Element", v.Type().String()))
}

// -----------------------------------------------------------------------------

func ElementOf(x any) Element {
	switch v := x.(type) {
	case HTML:
		tmpl := createElement("template")
		tmpl.Set("innerHTML", v.JSValue())
		return elementOf(fragment(tmpl.Get("content")))
	case string:
		return Query(v)
	case Wrapper:
		return elementOf(v.JSValue())
	case Value:
		return elementOf(v)
	}
	panic(fmt.Sprintf("unsupport type %T", x))
}

// -----------------------------------------------------------------------------

func BindValue(elm any, name string, val *string, cb func(string, string)) {
	e := ElementOf(elm)
	e.Bind(name, val, cb)
}
