//go:build js && wasm

package js

type Elements []Element

func (lst Elements) Length() int {
	return len(lst)
}

// -----------------------------------------------------------------------------

func (lst Elements) Foreach(fn func(int, Element)) {
	for i, elm := range lst {
		fn(i, elm)
	}
}

// -----------------------------------------------------------------------------

func (lst Elements) Filter(p func(Element) bool) (ret Elements) {
	ret = Elements{}

	for i := range lst {
		if p(lst[i]) {
			ret = append(ret, lst[i])
		}
	}

	return
}

// -----------------------------------------------------------------------------

func (lst Elements) FilterNot(p func(Element) bool) Elements {
	return lst.Filter(func(elm Element) bool {
		return !p(elm)
	})
}

// -----------------------------------------------------------------------------

// Attr returns some attribute values.
func (lst Elements) Attr(a string) (ret []string) {

	lst.Foreach(func(_ int, elm Element) {
		ret = append(ret, elm.Attr(a))
	})

	return
}

// -----------------------------------------------------------------------------

// SetAttr sets attribute to all elements.
func (e Elements) SetAttr(a, val string) Elements {
	e.Foreach(func(_ int, elm Element) {
		elm.SetAttr(a, val)
	})

	return e
}

// -----------------------------------------------------------------------------

// Prop returns some property values.
func (e Elements) Prop(p string) (ret []Value) {
	e.Foreach(func(_ int, elm Element) {
		ret = append(ret, elm.Prop(p))
	})

	return
}

// -----------------------------------------------------------------------------

// SetProp sets property to all elements.
func (e Elements) SetProp(a string, val Value) Elements {
	e.Foreach(func(_ int, elm Element) {
		elm.SetProp(a, val)
	})

	return e
}

// -----------------------------------------------------------------------------

func (e Elements) Add(clz string) Elements {
	e.Foreach(func(_ int, elm Element) {
		elm.Add(clz)
	})
	return e
}

// -----------------------------------------------------------------------------

func (e Elements) Remove(clz string) Elements {
	e.Foreach(func(_ int, elm Element) {
		elm.Remove(clz)
	})
	return e
}

// -----------------------------------------------------------------------------

func (e Elements) Toggle(clz string) Elements {
	e.Foreach(func(_ int, elm Element) {
		elm.Toggle(clz)
	})
	return e
}

// -----------------------------------------------------------------------------

func (e Elements) Replace(oldClz, newClz string) Elements {
	e.Foreach(func(_ int, elm Element) {
		elm.Replace(oldClz, newClz)
	})
	return e
}

// -----------------------------------------------------------------------------

func (e Elements) Has(clz string) Elements {
	return e.Filter(func(elm Element) bool {
		return elm.Has(clz)
	})
}

// -----------------------------------------------------------------------------

func (e Elements) Value() (ret []string) {
	e.Foreach(func(_ int, elm Element) {
		ret = append(ret, elm.Value())
	})
	return
}

// -----------------------------------------------------------------------------

func (e Elements) SetValue(val string) Elements {
	e.Foreach(func(_ int, elm Element) {
		e.SetValue(val)
	})
	return e
}

// -----------------------------------------------------------------------------

func ElementsOf(v Value) (ret Elements) {
	size := v.Length()
	ret = make([]Element, size)

	for i := 0; i < size; i++ {
		ret[i] = elementOf(v.Index(i))
	}

	return
}
