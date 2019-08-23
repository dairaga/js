package dom

import "github.com/dairaga/js"

// ElementList is a collection of elements.
type ElementList []*Element

var emptyNodeList = []*Element{}

// ElementListOf returns a element list.
func ElementListOf(v js.Value) ElementList {
	size := v.Length()
	if size <= 0 {
		return emptyNodeList
	}
	lst := make([]*Element, size)

	for i := 0; i < size; i++ {
		lst[i] = ElementOf(v.Index(i))
	}
	return lst
}

// Length returns the size of list.
func (nl ElementList) Length() int {
	return len(nl)
}

// Foreach apply function on each element.
func (nl ElementList) Foreach(fn func(idx int, v *Element)) {
	size := nl.Length()
	for i := 0; i < size; i++ {
		fn(i, nl[i])
	}
}

// Append add a new element.
func (nl ElementList) Append(x ...*Element) ElementList {
	return append(nl, x...)
}

// AppendList appends other list.
func (nl ElementList) AppendList(other ElementList) ElementList {
	return append(nl, other...)
}

// Attr returns some attribute values.
func (nl ElementList) Attr(name string) []string {
	var result []string

	nl.Foreach(func(_ int, e *Element) {
		result = append(result, e.Attr(name))
	})

	return result
}

// SetAttr sets attribute to all elements.
func (nl ElementList) SetAttr(name, value string) ElementList {
	nl.Foreach(func(_ int, e *Element) {
		e.SetAttr(name, value)
	})

	return nl
}

// Prop returns some property values.
func (nl ElementList) Prop(name string) []js.Value {
	var result []js.Value
	nl.Foreach(func(_ int, e *Element) {
		result = append(result, e.Prop(name))
	})
	return result
}

// SetProp sets property to all elements.
func (nl ElementList) SetProp(name string, flag bool) ElementList {
	nl.Foreach(func(_ int, e *Element) {
		e.SetProp(name, flag)
	})
	return nl
}

// AddClass adds class to all elements.
func (nl ElementList) AddClass(names ...string) ElementList {
	nl.Foreach(func(_ int, e *Element) {
		e.AddClass(names...)
	})
	return nl
}

// RemoveClass remove some class from all elements.
func (nl ElementList) RemoveClass(names ...string) ElementList {
	nl.Foreach(func(_ int, e *Element) {
		e.RemoveClass(names...)
	})
	return nl
}

// ToggleClass toggles some class on all elements.
func (nl ElementList) ToggleClass(name string) ElementList {
	nl.Foreach(func(_ int, e *Element) {
		e.ToggleClass(name)
	})

	return nl
}

// ReplaceClass replace some class with new one on all elements.
func (nl ElementList) ReplaceClass(oldName, newName string) ElementList {
	nl.Foreach(func(_ int, e *Element) {
		e.ReplaceClass(oldName, newName)
	})

	return nl
}

// HasClass returns elements which has specific class.
func (nl ElementList) HasClass(name string) ElementList {
	var result []*Element

	nl.Foreach(func(_ int, e *Element) {
		if e.HasClass(name) {
			result = append(result, e)
		}
	})

	return result
}

// TagName returns all tag names.
func (nl ElementList) TagName() []string {
	var result []string

	nl.Foreach(func(_ int, e *Element) {
		result = append(result, e.TagName())
	})

	return result
}

// Val returns all values.
func (nl ElementList) Val() []string {
	var result []string

	nl.Foreach(func(_ int, e *Element) {
		result = append(result, e.Val())
	})

	return result
}

// SetVal set value to all elements.
func (nl ElementList) SetVal(val interface{}) ElementList {

	nl.Foreach(func(_ int, e *Element) {
		nl.SetVal(val)
	})
	return nl
}
