//go:build js && wasm

package js

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreend(t *testing.T) {
	Prepend(Plain("hello_prepend"))
	assert.Equal(t, 0, strings.Index(document.Get("body").Get("innerText").String(), "hello_prepend"))
}

// -----------------------------------------------------------------------------

func TestCreateElement(t *testing.T) {
	elm := CreateElement("button")
	elm.SetAttr("id", "btn_create_element")
	elm.SetText("Test Create Element")
	Append(elm)

	assert.Equal(t, elm.JSValue(), document.Call("querySelector", "#btn_create_element"))
	assert.Equal(t, elm, Query("#btn_create_element"))
}

// -----------------------------------------------------------------------------

func TestApend(t *testing.T) {

	Append(Plain("hello_append"))
	body := document.Get("body").Get("innerText").String()
	pos := len(body) - len("hello_append")
	assert.Equal(t, pos, strings.LastIndex(body, "hello_append"))
}

// -----------------------------------------------------------------------------

func TestQueryAll(t *testing.T) {
	testData := `<div>
	<ul>
		<li>0</li>
		<li>1</li>
		<li>2</li>
		<li>3</li>
	</ul>
	</div>`

	Append(ElementOf(HTML(testData)))

	lst := QueryAll("li")
	for i, item := range lst {
		assert.True(t, len(item.Tattoo()) > 0)
		assert.Equal(t, strconv.Itoa(i), item.Text())
	}
}

// -----------------------------------------------------------------------------

func TestRemoveChild(t *testing.T) {
	testData := `<div><div id='div_1'></div><div id='div_2'></div><div id='div_3'></div><div id='div_4'></div></div>`
	testElm := ElementOf(HTML(testData))

	Append(testElm)
	assert.NotPanics(t, func() {
		_ = Query("#div_1")
		div2 := Query("#div_2")
		div3 := Query("#div_3")
		_ = Query("#div_4")

		RemoveChild("#div_1")
		RemoveChild(div2)
		RemoveChild(div3.JSValue())
	})

	assert.Panics(t, func() { Query("#div_1") })
	assert.Panics(t, func() { Query("#div_2") })
	assert.Panics(t, func() { Query("#div_3") })

	_ = Query("#div_4")
}
