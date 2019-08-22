package main

import (
	"fmt"

	"github.com/dairaga/js/dom/bs/list"
	"github.com/dairaga/js/dom/bs/table"

	"github.com/dairaga/js/dom/bs/card"
	"github.com/dairaga/js/dom/bs/progress"

	"github.com/dairaga/js/dom"

	"github.com/dairaga/js/dom/bs"
	"github.com/dairaga/js/dom/bs/badge"
	"github.com/dairaga/js/dom/bs/spinner"
)

func main() {
	fmt.Println("hello")
	container := dom.S("#main")

	ExampleTable(container)

	ExampleBadge(container)

	ExampleSpinner(container)

	ExampleProgress(container)

	ExampleCard(container)

	ExampleListGroup(container)

	container.Prepend(badge.Link(bs.Danger, "test prepend"))
	select {}
}

// ExampleTable is an sample code about table.
func ExampleTable(container *dom.Element) {
	t := table.Attach("#test_table")
	fmt.Println("all table:", t)
	fmt.Println("header:", t.Header())
	fmt.Println("body:", t.Body())

	for i := 0; i < 4; i++ {
		fmt.Println(t.Head(i))
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println(t.Cell(i, j))
		}
	}

	t = table.New(
		[]interface{}{"#", "First", "Last", "Handle"},
		[][]interface{}{
			[]interface{}{1, "Mark", "Otto", "@mdo"},
			[]interface{}{2, "Jacob", "Thornton", "@fat"},
			[]interface{}{3, "Larry", "the Bird", "@twitter"},
		},
	).Caption("List of users")

	container.Append(t)
}

// ExampleBadge is a sample code about Badge.
func ExampleBadge(container *dom.Element) {
	b := badge.New(bs.Primary, "my test").Pill()
	container.Append(b)
	b = badge.Link(bs.Danger, "my test link").Pill()
	container.Append(b)
}

// ExampleSpinner is a sample code about Spinner.
func ExampleSpinner(container *dom.Element) {
	sp := spinner.Border(bs.FGPrimary)
	container.Append(sp)

	sp = spinner.Grow(bs.FGDanger).Smaller()
	container.Append(sp)
}

// ExampleProgress is a sample code about progress bar.
func ExampleProgress(container *dom.Element) {
	pb := progress.New(bs.BGInfo, 0, 100, 41)
	pb.Bar(0).Stripped().Animate().Show("41%")
	pb.Add(progress.NewBar(bs.BGPrimary, 0, 100, 32).Stripped().Animate().Show("32%"))

	container.Append(pb)
}

// ExampleCard is a sample code about card.
func ExampleCard(container *dom.Element) {
	crd := card.New()
	crd.AddClass("border-success")
	crdheader := card.NewHeader("Header", badge.New(bs.Primary, "abc"))
	crdheader.AddClass("border-success")
	crd.Header(crdheader)

	crdbody := card.NewBody(nil)

	crdbody.Color(bs.FGSuccess)
	crdbody.Title("Success card title")
	crdbody.SubTitle("Success card subtitle")
	crdbody.Text(`Some quick example text to build on the card title and make up the bulk of the card's content.`)
	crdbody.Link("#", "link 1")
	crdbody.Link("#", "link 2")
	crd.Append(crdbody)

	crdfooter := card.NewFooter("Footer")
	crdfooter.AddClass("border-success")
	crd.Footer(crdfooter)

	crd.Width(bs.Size25)
	container.Append(crd)
}

// ExampleListGroup is an sample code about list group.
func ExampleListGroup(container *dom.Element) {
	g := list.New([][]interface{}{
		[]interface{}{`Cras justo odio`, badge.New(bs.Warning, "14").Pill()},
		[]interface{}{`Dapibus ac facilisis in`, badge.New(bs.Warning, "2").Pill()},
		[]interface{}{`Morbi leo risus`, badge.New(bs.Warning, "1").Pill()},
	})

	g.Add(`test add item`, badge.New(bs.Warning, "100").Pill())

	g.Foreach(func(_ int, it *list.Item) {
		it.AddClass("d-flex", "justify-content-between", "align-items-center")
		it.Action()
	})
	g.Active(1, true)
	g.Disable(2, true)
	g.Width(bs.Size25)
	container.Append(g)

	g = list.Button([][]interface{}{
		[]interface{}{`Cras justo odio`, badge.New(bs.Warning, "14").Pill()},
		[]interface{}{`Dapibus ac facilisis in`, badge.New(bs.Warning, "2").Pill()},
		[]interface{}{`Morbi leo risus`, badge.New(bs.Warning, "1").Pill()},
	})

	g.Foreach(func(_ int, it *list.Item) {
		it.AddClass("d-flex", "justify-content-between", "align-items-center")
		it.Action()
	})
	g.Active(1, true)
	g.Disable(2, true)
	g.Width(bs.Size25)
	container.Append(g)

	g = list.Link([][]interface{}{
		[]interface{}{`#1`, `Cras justo odio`, badge.New(bs.Warning, "14").Pill()},
		[]interface{}{`#2`, `Dapibus ac facilisis in`, badge.New(bs.Warning, "2").Pill()},
		[]interface{}{`#3`, `Morbi leo risus`, badge.New(bs.Warning, "1").Pill()},
	})

	g.Foreach(func(_ int, it *list.Item) {
		it.AddClass("d-flex", "justify-content-between", "align-items-center")
		it.Action()
	})
	g.Active(1, true)
	g.Disable(2, true)
	g.Width(bs.Size25)
	container.Append(g)
}
