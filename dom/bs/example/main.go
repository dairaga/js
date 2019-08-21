package main

import (
	"fmt"

	"github.com/dairaga/js/dom/bs/card"
	"github.com/dairaga/js/dom/bs/progress"

	"github.com/dairaga/js/dom"

	"github.com/dairaga/js/dom/bs"
	"github.com/dairaga/js/dom/bs/badge"
	"github.com/dairaga/js/dom/bs/spinner"
	"github.com/dairaga/js/dom/bs/table"
)

func main() {
	fmt.Println("hello")
	container := dom.S("#main")
	/* table example start */
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
	/* table example end */

	/* badge start */
	b := badge.New(bs.Primary, "my test").Pill()
	container.Append(b)
	b = badge.Link(bs.Danger, "my test link").Pill()
	container.Append(b)
	/* badge end */

	/* spinner start */
	sp := spinner.Border(bs.FGPrimary)
	container.Append(sp)

	sp = spinner.Grow(bs.FGDanger).Smaller()
	container.Append(sp)
	/* spinner end */

	/* progress start */
	pb := progress.New(bs.BGInfo, 0, 100, 41)
	pb.Bar(0).Stripped().Animate().SetText("41")
	container.Append(pb)
	/* progress end */

	/* card start */
	crd := card.New()
	crd.AddClass("border-success")
	crdheader := card.NewHeader("Header")
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
	/* card end */

	select {}
}
