//go:build js && wasm

package js

import (
	"fmt"

	"github.com/dairaga/js/v2/builtin"
)

// Template is a DocumentFragment from content property of HTMLTemplateElement.
type Template Value

var _ Appendable = Template{}

// -----------------------------------------------------------------------------

func (t Template) Travel(selector string, fn func(int, Element)) Template {
	ElementsOf(queryAll(Value(t), selector)).Foreach(fn)
	return t
}

// -----------------------------------------------------------------------------

func (t Template) JSValue() Value {
	return Value(t)
}

// -----------------------------------------------------------------------------

func (t Template) Ref() Value {
	return Value(t)
}

// -----------------------------------------------------------------------------

func (t Template) Query(selector string) Element {
	return elementOf(query(Value(t), selector))
}

// -----------------------------------------------------------------------------

func (t Template) QueryAll(selector string) Elements {
	return ElementsOf(queryAll(Value(t), selector))
}

// -----------------------------------------------------------------------------

func (t Template) Append(child Appendable, selector ...string) Template {
	//at(Value(t), selector...).Call("append", child.Ref())
	appendNode(at(Value(t), selector...), child)
	return t
}

// -----------------------------------------------------------------------------

func (t Template) Prepend(child Appendable, selector ...string) Template {
	//at(Value(t), selector...).Call("prepend", child.Ref())
	prependNode(at(Value(t), selector...), child)
	return t
}

// -----------------------------------------------------------------------------

func (t Template) Prop(selector, p string) Value {
	return query(Value(t), selector).Get(p)
}

// -----------------------------------------------------------------------------

func (t Template) SetProp(selector, p string, val any) Template {
	return t.Travel(selector, func(_ int, child Element) {
		child.SetProp(p, val)
	})
}

// -----------------------------------------------------------------------------

func (t Template) Attr(selector, a string) string {
	return attr(query(Value(t), selector), a)
}

// -----------------------------------------------------------------------------

func (t Template) SetAttr(selector, a, val string) Template {
	if a != _tattoo {
		t.Travel(selector, func(_ int, child Element) {
			setAttr(child.JSValue(), a, val)
		})
	}

	return t
}

// -----------------------------------------------------------------------------

func (t Template) Text(selector string) string {
	return query(Value(t), selector).Get("innerText").String()
}

// -----------------------------------------------------------------------------

func (t Template) SetText(selector string, txt string) Template {
	return t.Travel(selector, func(_ int, child Element) {
		child.SetText(txt)
	})
}

// -----------------------------------------------------------------------------

func (t Template) HTML(selector string) HTML {
	return HTML(query(Value(t), selector).Get("innerHTML").String())
}

// -----------------------------------------------------------------------------

func (t Template) SetHTML(selector string, content HTML) Template {
	return t.Travel(selector, func(_ int, child Element) {
		child.SetHTML(content)
	})
}

// -----------------------------------------------------------------------------

func (t Template) Add(selector, clz string) Template {
	return t.Travel(selector, func(_ int, child Element) {
		//child.Add(clz)
		addClz(child.JSValue(), clz)
	})
}

// -----------------------------------------------------------------------------

func (t Template) Remove(selector, clz string) Template {
	return t.Travel(selector, func(_ int, child Element) {
		//child.Remove(clz)
		removeClz(child.JSValue(), clz)
	})
}

// -----------------------------------------------------------------------------

func (t Template) Has(selector, clz string) bool {
	return hasClz(query(Value(t), selector), clz)
}

// -----------------------------------------------------------------------------

func (t Template) Replace(selector, old, new string) Template {
	return t.Travel(selector, func(_ int, child Element) {
		//child.Replace(old, new)
		replaceClz(child.JSValue(), old, new)
	})
}

// -----------------------------------------------------------------------------

func (t Template) Toggle(selector, clz string) Template {
	return t.Travel(selector, func(_ int, child Element) {
		//child.Toggle(clz)
		toggleClz(child.JSValue(), clz)
	})
}

// -----------------------------------------------------------------------------

func (t Template) Clone() Template {
	return Template(Value(t).Call("cloneNode", true))
}

// -----------------------------------------------------------------------------

func (t Template) Children() Elements {
	return ElementsOf(Value(t).Get("children"))
}

// -----------------------------------------------------------------------------

func (t Template) First() Element {
	return elementOf(Value(t).Get("firstElementChild"))
}

// -----------------------------------------------------------------------------

func (t Template) Last() Element {
	return elementOf(Value(t).Get("lastElementChild"))
}

// -----------------------------------------------------------------------------

func (t Template) Length() int {
	return Value(t).Get("childElementCount").Int()
}

// -----------------------------------------------------------------------------

func CreateTemplate(content HTML) Template {
	tmpl := createElement("template")

	tmpl.Set("innerHTML", content.JSValue())

	return Template(tmpl.Get("content"))
}

// -----------------------------------------------------------------------------

func CreateTemplateByID(id string) Template {
	v := query(body, id)
	if builtin.IsTemplate(v) {
		panic(fmt.Sprintf("%s is not a template", id))
	}

	return Template(v.Get("content"))
}
