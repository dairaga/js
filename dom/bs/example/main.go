package main

import (
	"fmt"

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

	select {}
}
