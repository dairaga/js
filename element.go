//go:build js && wasm

package js

import (
	"fmt"

	"github.com/dairaga/js/v2/builtin"
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

type HandlerFunc func(Element, Event)

// -----------------------------------------------------------------------------

type Element interface {
	Appendable

	Parent() Element
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
	Var(val *string) Element

	Files() []File

	Add(clz string, at ...string) Element
	Remove(clz string, at ...string) Element
	Has(clz string, at ...string) bool
	Replace(oldClz, newClz string, at ...string) Element
	Toggle(clz string, at ...string) Element

	On(typ string, fn HandlerFunc, at ...string) Element
	OnClick(fn HandlerFunc, at ...string) Element
	OnChange(fn HandlerFunc, at ...string) Element

	Click(at ...string) Element
	Foucs(at ...string) Element
	Blur(at ...string) Element

	Empty() Element
	Relese()
}

// -----------------------------------------------------------------------------

type element Value

var _ Appendable = element{}

// -----------------------------------------------------------------------------

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

func (e element) String() string {
	return fmt.Sprintf(`{"tag": %q, "tattoo": %q, "id": %q, "class": %q}`, Value(e).Get("tagName").String(), e.Tattoo(), e.Attr("id"), e.Prop("classList").Get("value").String())
}

// -----------------------------------------------------------------------------

func (e element) Parent() Element {
	p := e.Prop("parentElement")
	if p.Truthy() && builtin.IsElement(p) {
		return elementOf(p)
	}
	return nil
}

// -----------------------------------------------------------------------------

func (e element) Tattoo() string {
	return e.Attr(_tattoo)
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

func (e element) Append(child Appendable, selector ...string) Element {
	//at(Value(e), selector...).Call("append", child.Ref())
	appendNode(at(Value(e), selector...), child)
	return e
}

// -----------------------------------------------------------------------------

func (e element) Prepend(child Appendable, selector ...string) Element {
	//at(Value(e), selector...).Call("prepend", child.Ref())
	prependNode(at(Value(e), selector...), child)
	return e
}

// -----------------------------------------------------------------------------

func (e element) Prop(p string, selector ...string) Value {
	return at(Value(e), selector...).Get(p)
}

// -----------------------------------------------------------------------------

func (e element) SetProp(p string, val any, selector ...string) Element {
	at(Value(e), selector...).Set(p, val)
	return e
}

// -----------------------------------------------------------------------------

func (e element) Attr(a string, selector ...string) string {
	return attr(at(Value(e), selector...), a)
}

// -----------------------------------------------------------------------------

func (e element) SetAttr(a, val string, selector ...string) Element {
	if _tattoo != a {
		setAttr(at(Value(e), selector...), a, val)
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

func (e element) Add(clz string, selector ...string) Element {
	addClz(at(Value(e), selector...), clz)
	//e.Prop("classList", at...).Call("add", clz)
	return e
}

// -----------------------------------------------------------------------------

func (e element) Remove(clz string, selector ...string) Element {
	//e.Prop("classList", at...).Call("remove", clz)
	removeClz(at(Value(e), selector...), clz)
	return e
}

// -----------------------------------------------------------------------------

func (e element) Has(clz string, selector ...string) bool {
	//return e.Prop("classList", at...).Call("contains", clz).Bool()
	return hasClz(at(Value(e), selector...), clz)
}

// -----------------------------------------------------------------------------

func (e element) Replace(old, new string, selector ...string) Element {
	//e.Prop("classList", at...).Call("replace", oldClz, newClz)
	replaceClz(at(Value(e), selector...), old, new)
	return e
}

// -----------------------------------------------------------------------------

func (e element) Toggle(clz string, selector ...string) Element {
	//e.Prop("classList", at...).Call("toggle", clz)
	toggleClz(at(Value(e), selector...), clz)
	return e
}

// -----------------------------------------------------------------------------

func (e element) On(typ string, fn HandlerFunc, selector ...string) Element {
	cb := FuncOf(func(_this Value, args []Value) any {
		evt := event(args[0])
		//elm := elementOf(evt.Get("target"))
		elm := elementOf(_this)
		fn(elm, evt)
		return nil
	})

	at(Value(e), selector...).Call("addEventListener", typ, cb)
	return e
}

// -----------------------------------------------------------------------------

func (e element) OnClick(fn HandlerFunc, at ...string) Element {
	return e.On("click", fn, at...)
}

// -----------------------------------------------------------------------------

func (e element) OnChange(fn HandlerFunc, at ...string) Element {
	return e.On("change", fn, at...)
}

// -----------------------------------------------------------------------------

func (e element) Click(selector ...string) Element {
	at(Value(e), selector...).Call("click")
	return e
}

// -----------------------------------------------------------------------------

func (e element) Foucs(selector ...string) Element {
	at(Value(e), selector...).Call("focus")
	return e
}

// -----------------------------------------------------------------------------

func (e element) Blur(selector ...string) Element {
	at(Value(e), selector...).Call("blur")
	return e
}

// -----------------------------------------------------------------------------

func (e element) Empty() Element {
	v := Value(e)
	for v.Get("firstChild").Truthy() {
		v.Get("firstChild").Call("remove")
	}
	return e
}

// -----------------------------------------------------------------------------

func (e element) Relese() {
	Value(e).Call("remove")
}

// -----------------------------------------------------------------------------

func (e element) Var(val *string) Element {
	if e.Attr("type") == "checkbox" && !e.Prop("checked").Bool() {
		*val = ""
	} else {
		*val = e.Value()
	}
	return e
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
	case Element:
		return v
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

func Var(elm any, val *string) {
	e := ElementOf(elm)
	e.Var(val)
}
