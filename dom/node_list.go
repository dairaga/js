package dom

import "github.com/dairaga/js"

// NodeList ...
type NodeList []Element

var emptyNodeList = []Element{}

// NodeListOf ...
func NodeListOf(v js.Value) NodeList {
	size := v.Length()
	if size <= 0 {
		return emptyNodeList
	}
	lst := make([]Element, size)

	for i := 0; i < size; i++ {
		lst[i] = ElementOf(v.Index(i))
	}
	return lst
}

// Length ...
func (nl NodeList) Length() int {
	return len(nl)
}

// Foreach ...
func (nl NodeList) Foreach(fn func(v Element)) {
	size := nl.Length()
	for i := 0; i < size; i++ {
		fn(nl[i])
	}
}

// Append ...
func (nl NodeList) Append(x ...Element) NodeList {
	return append(nl, x...)
}

// AppendList ...
func (nl NodeList) AppendList(other NodeList) NodeList {
	return append(nl, other...)
}

// Attr ...
func (nl NodeList) Attr(name string) []string {
	var result []string

	nl.Foreach(func(e Element) {
		result = append(result, e.Attr(name))
	})

	return result
}

// SetAttr ...
func (nl NodeList) SetAttr(name, value string) NodeList {
	nl.Foreach(func(e Element) {
		e.SetAttr(name, value)
	})

	return nl
}

// Prop ...
func (nl NodeList) Prop(name string) []js.Value {
	var result []js.Value
	nl.Foreach(func(e Element) {
		result = append(result, e.Prop(name))
	})
	return result
}

// SetProp ...
func (nl NodeList) SetProp(name string, flag bool) NodeList {
	nl.Foreach(func(e Element) {
		e.SetProp(name, flag)
	})
	return nl
}

// AddClass ...
func (nl NodeList) AddClass(names ...string) NodeList {
	nl.Foreach(func(e Element) {
		e.AddClass(names...)
	})
	return nl
}

// RemoveClass ...
func (nl NodeList) RemoveClass(names ...string) NodeList {
	nl.Foreach(func(e Element) {
		e.RemoveClass(names...)
	})
	return nl
}

// ToggleClass ...
func (nl NodeList) ToggleClass(name string) NodeList {
	nl.Foreach(func(e Element) {
		e.ToggleClass(name)
	})

	return nl
}

// ReplaceClass ...
func (nl NodeList) ReplaceClass(oldName, newName string) NodeList {
	nl.Foreach(func(e Element) {
		e.ReplaceClass(oldName, newName)
	})

	return nl
}

// HasClass ...
func (nl NodeList) HasClass(name string) []bool {
	var result []bool

	nl.Foreach(func(e Element) {
		result = append(result, e.HasClass(name))
	})

	return result
}

// TagName ...
func (nl NodeList) TagName() []string {
	var result []string

	nl.Foreach(func(e Element) {
		result = append(result, e.TagName())
	})

	return result
}

// Val ...
func (nl NodeList) Val() []string {
	var result []string

	nl.Foreach(func(e Element) {
		result = append(result, e.Val())
	})

	return result
}

// SetVal ...
func (nl NodeList) SetVal(val interface{}) NodeList {

	nl.Foreach(func(e Element) {
		nl.SetVal(val)
	})
	return nl
}
