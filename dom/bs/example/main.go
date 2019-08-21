package main

import (
	"fmt"

	"github.com/dairaga/js/dom"

	"github.com/dairaga/js/dom/bs/table"
)

func main() {
	fmt.Println("hello")

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

	dom.AppendChild(t)
	/* table example end */

	select {}
}
