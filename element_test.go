//go:build js && wasm

package js

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElementOfByID(t *testing.T) {
	btn := document.Call("createElement", "button")
	btn.Call("setAttribute", "id", "btn_create_element_by_id")
	btn.Set("innerText", "Test Create Element By ID")
	document.Get("body").Call("append", btn)

	elm := ElementOf("#btn_create_element_by_id")
	assert.Equal(t, "btn_create_element_by_id", elm.Attr("id"))
	assert.Equal(t, document.Call("querySelector", "#btn_create_element_by_id"), elm.JSValue())
	assert.Equal(t, elm, Query("#btn_create_element_by_id"))
}

// -----------------------------------------------------------------------------

func TestElementOfByHTML(t *testing.T) {

	html := `<button id='create_element_by_html'>Create Element By HTML</button>`
	elm := ElementOf(HTML(html))
	Append(elm)

	assert.Equal(t, "create_element_by_html", elm.Attr("id"))
	assert.Equal(t, document.Call("querySelector", "#create_element_by_html"), elm.JSValue())
	assert.Equal(t, elm, Query("#create_element_by_html"))
	elm.SetText("Change Content")
	assert.Equal(t, Query("#create_element_by_html").Text(), elm.Text())
}
