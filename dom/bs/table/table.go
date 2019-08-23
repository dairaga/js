package table

import (
	"fmt"

	"github.com/dairaga/js/dom"
	"github.com/dairaga/js/dom/bs"
)

// Table represents Bootstrap table content.
type Table struct {
	*bs.Component
	head *bs.Component
	body *bs.Component
}

// Attach binds a Bootstrap table on page.
func Attach(id string) *Table {
	t := &Table{bs.Attach(id), nil, nil}
	t.head = bs.ComponentOf(dom.S("thead"))
	t.body = bs.ComponentOf(dom.S("tbody"))

	return t
}

// New returns a table with headers and data.
func New(head []interface{}, data [][]interface{}) *Table {
	table := &Table{bs.ComponentOf(dom.CreateElement("table")), nil, nil}
	table.AddClass("table")
	table.head = bs.ComponentOf(dom.CreateElement("thead"))
	table.Append(table.head)

	if len(head) > 0 {
		tr := dom.CreateElement("tr")
		tr.Append(dom.CreateElement("th").SetAttr("scope", "col").Append("#"))

		for _, x := range head {
			th := dom.CreateElement("th").SetAttr("scope", "col").Append(x)
			tr.Append(th)
		}

		table.head.Append(tr)
	}

	table.body = bs.ComponentOf(dom.CreateElement("tbody"))
	table.Append(table.body)
	for i, x := range data {
		table.Add(i, x)
	}

	return table
}

// ----------------------------------------------------------------------------

// Caption add or replace caption.
func (t *Table) Caption(x interface{}) *Table {
	elm := t.S("caption")
	if elm.Truthy() {
		t.RemoveChild(elm)
	}
	elm = dom.CreateElement("caption")
	t.Append(elm)
	elm.Append(x)
	return t
}

/*
// Header returns theader.
func (t *Table) Header() *bs.Component {
	elm := t.S("thead")
	return bs.ComponentOf(elm)
}
*/

// Head return n-th head cell. idx is 0-index.
func (t *Table) Head(idx int) *bs.Component {
	selector := fmt.Sprintf("thead tr th:nth-child(%d)", idx+1)
	elm := t.S(selector)
	return bs.ComponentOf(elm)
}

/*
// Body returns tbody.
func (t *Table) Body() *bs.Component {
	elm := t.S("tbody")
	return bs.ComponentOf(elm)
}
*/

// Cell return (row, col) cell in tbody. row and col are 0-index.
func (t *Table) Cell(row, col int) *bs.Component {
	selector := ""
	if col <= 0 {
		selector = fmt.Sprintf("tbody tr:nth-child(%d) th:nth-child(1)", row+1)
	} else {
		selector = fmt.Sprintf("tbody tr:nth-child(%d) td:nth-child(%d)", row+1, col+1)
	}

	elm := t.S(selector)
	return bs.ComponentOf(elm)
}

// Add adds data to table.
func (t *Table) Add(idx interface{}, data []interface{}) *Table {
	tr := dom.CreateElement("tr")
	th := dom.CreateElement("th")
	th.SetAttr("scope", "row").Append(idx)
	tr.Append(th)
	for _, x := range data {
		tr.Append(dom.CreateElement("td").Append(x))
	}

	t.body.Append(tr)

	return t
}
