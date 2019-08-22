package list

import (
	"github.com/dairaga/js/dom"
	"github.com/dairaga/js/dom/bs"
)

// Item represents Bootstrap list group item.
type Item struct {
	*bs.Component
}

func generateItem(tag string, content ...interface{}) *Item {
	it := &Item{bs.ComponentOf(dom.CreateElement(tag))}
	it.AddClass("list-group-item").Append(content...)
	return it
}

// Action adds action style.
func (it *Item) Action() *Item {
	it.AddClass("list-group-item-action")
	return it
}

// Active actives the item.
func (it *Item) Active(flag bool) *Item {
	if flag {
		it.AddClass("active")
	} else {
		it.RemoveClass("active")
	}
	return it
}

// Actived returns boolean indicates whether or not the item is active.
func (it *Item) Actived() bool {
	return it.HasClass("active")
}

// Disable disables the item.
func (it *Item) Disable(flag bool) *Item {
	if flag {
		it.AddClass("disabled").SetAttr("aria-disabled", "true")
	} else {
		it.RemoveClass("disabled").RemoveAttr("aria-disabled")
	}
	return it
}

// Disabled returns boolean indicates whether or not the item is disabled.
func (it *Item) Disabled() bool {
	return it.HasClass("disabled") || it.Prop("disabled").Bool()
}

// Style adds bootstrap pre-defined style like primary, secondary and etc.
func (it *Item) Style(style bs.Style) *Item {
	it.AddClass("list-group-item-" + style)
	return it
}

// ----------------------------------------------------------------------------

// Group represents Bootstrap list group container.
type Group struct {
	*bs.Component
	childTag string
	items    []*Item
}

func generate(tag, itemTag string, data [][]interface{}) *Group {
	group := &Group{bs.ComponentOf(dom.CreateElement(tag)), itemTag, nil}

	for _, content := range data {
		it := generateItem(itemTag, content...)
		group.items = append(group.items, it)
		group.Append(it)
	}
	group.AddClass("list-group")

	return group
}

// New returns a list group with "<ul>"
func New(data [][]interface{}) *Group {
	return generate("ul", "li", data)
}

// Button returns a list group with <div> and <button> item.
func Button(data [][]interface{}) *Group {
	return generate("div", "button", data)
}

// Link returns a list group with <div> and <a> item.
// The dataset must be slice of []interface{}{link, ...}
func Link(data [][]interface{}) *Group {
	g := generate("div", "a", data[1:])
	for i, x := range g.items {
		x.SetAttr("href", data[i][0].(string))
	}
	return g
}

// Flush removes some borders and rounded corners to render list group items edge-to-edge in a parent container (e.g., cards).
func (g *Group) Flush() *Group {
	g.AddClass("list-group-flush")
	return g
}

// Horizontalize returns a horizontal list group.
func (g *Group) Horizontalize(vs ...bs.ViewportSize) *Group {
	clz := "list-group-horizontal"
	if len(vs) > 0 {
		clz = clz + "-" + vs[0]
	}

	g.AddClass(clz)
	return g
}

// Length returns length of items.
func (g *Group) Length() int {
	return len(g.items)
}

// Item returns nth item.
func (g *Group) Item(idx int) *Item {
	if idx < 0 || idx >= len(g.items) {
		return nil
	}

	return g.items[idx]
}

// Foreach applies function fn on each item.
func (g *Group) Foreach(fn func(int, *Item)) *Group {
	for i, it := range g.items {
		fn(i, it)
	}
	return g
}

// Disable disables nth item.
func (g *Group) Disable(idx int, flag bool) *Group {
	it := g.Item(idx)
	if it != nil {
		it.Disable(flag)
	}
	return g
}

// Active actives nth item.
func (g *Group) Active(idx int, flag bool) *Group {
	it := g.Item(idx)
	if it != nil {
		it.Active(flag)
	}

	return g
}

// InactiveItems resets all items to inactive.
func (g *Group) InactiveItems() *Group {
	return g.Foreach(func(_ int, it *Item) {
		it.Active(false)
	})
}

// EnableItems reset all items to enabled.
func (g *Group) EnableItems() *Group {
	return g.Foreach(func(_ int, it *Item) {
		it.Disable(false)
	})
}

// Reset reset all items to enabled and inactive.
func (g *Group) Reset() *Group {
	return g.InactiveItems().EnableItems()
}

// Add adds a new item with content and returns the new item.
func (g *Group) Add(content ...interface{}) *Item {
	it := generateItem(g.childTag, content...)
	g.items = append(g.items, it)
	g.Append(it)
	return it
}

// Action adds action style to all items.
func (g *Group) Action() *Group {
	g.Foreach(func(_ int, it *Item) {
		it.Action()
	})

	return g
}
