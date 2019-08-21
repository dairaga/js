package table

import (
	"fmt"

	"github.com/dairaga/js/dom"
	"github.com/dairaga/js/dom/bs"
)

// Table represents Bootstrap table content.
type Table struct {
	*bs.Component
}

// Attach binds a Bootstrap table on page
func Attach(id string) *Table {
	return &Table{bs.Attach(id)}
}

// New returns a table with headers and data.
func New(header []interface{}, data [][]interface{}) *Table {
	table := &Table{bs.ComponentOf(dom.CreateElement("table"), "")}
	table.AddClass("table")
	head := dom.CreateElement("thead")
	if len(header) > 0 {
		tr := dom.CreateElement("tr")

		for _, x := range header {
			th := dom.CreateElement("th").SetAttr("scope", "col").Append(x)
			tr.Append(th)
		}

		head.Append(tr)
	}
	table.Append(head)

	body := dom.CreateElement("tbody")

	for _, x := range data {
		tr := dom.CreateElement("tr")

		for i, y := range x {
			if i == 0 {
				th := dom.CreateElement("th")
				th.SetAttr("scope", "row").Append(y)
				tr.Append(th)
				continue
			}

			tr.Append(dom.CreateElement("td").Append(y))

		}

		body.Append(tr)
	}

	table.Append(body)

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

// Header returns theader.
func (t *Table) Header() *bs.Component {
	elm := t.S("thead")
	return bs.ComponentOf(elm, "")
}

// Head return n-th header cell. idx is 0-index.
func (t *Table) Head(idx int) *bs.Component {
	selector := fmt.Sprintf("thead tr th:nth-child(%d)", idx+1)
	elm := t.S(selector)
	return bs.ComponentOf(elm, "")
}

// Body returns tbody.
func (t *Table) Body() *bs.Component {
	elm := t.S("tbody")
	return bs.ComponentOf(elm, "")
}

// Cell return (row, col) cell in tbody. row and col are 0-index.
func (t *Table) Cell(row, col int) *bs.Component {
	selector := ""
	if col <= 0 {
		selector = fmt.Sprintf("tbody tr:nth-child(%d) th:nth-child(1)", row+1)
	} else {
		selector = fmt.Sprintf("tbody tr:nth-child(%d) td:nth-child(%d)", row+1, col+1)
	}

	elm := t.S(selector)
	return bs.ComponentOf(elm, "")
}
