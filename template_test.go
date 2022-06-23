//go:build js && wasm

package js_test

import (
	"testing"

	"github.com/dairaga/js/v2"
	"github.com/stretchr/testify/assert"
)

func TestTemlateCreate(t *testing.T) {
	texts := [][]string{
		{"tmpl_li_1", "tmpl-1"},
		{"tmpl_li_2", "tmpl-2"},
		{"tmpl_li_3", "tmpl-3"},
	}

	clz := []string{"clz-1", "clz-2", "clz-3"}
	evenClz := []bool{true, false, true}
	oddClz := []bool{false, true, false}

	content := js.HTML(`<ul id="tmpl_ul"><li></li><li></li><li></li></ul>`)

	tmpl := js.CreateTemplate(content)

	// Initialize
	tmpl.Travel("li", func(idx int, child js.Element) {
		child.SetAttr("id", texts[idx][0])
		child.SetText(texts[idx][1])
		child.Add(clz[idx])
		if idx&0x01 == 0 {
			child.Add("clz-even")
		} else {
			child.Add("clz-odd")
		}
	})

	// Test Prepend, Append, First, and Last
	tmpl.Prepend(js.CreateElement("p").SetAttr("id", "tmpl_p_first").SetText("first"))
	tmpl.Append(js.CreateElement("p").SetAttr("id", "tmpl_p_last").SetText("last"))
	assert.Equal(t, "tmpl_p_first", tmpl.First().Attr("id"))
	assert.Equal(t, "tmpl_p_last", tmpl.Last().Attr("id"))
	assert.Equal(t, 3, tmpl.Length()) // <p> x 2 + <ul> x 1

	// Print li
	//tmpl.Travel("li", func(i int, child js.Element) {
	//	t.Log(i, child)
	//})

	// Test Query
	for i := range texts {
		assert.Equal(t, texts[i][1], tmpl.Query("#"+texts[i][0]).Text())
	}

	for i := range clz {
		assert.True(t, tmpl.Query("#"+texts[i][0]).Has(clz[i]))
	}

	for i := range evenClz {
		assert.Equal(t, evenClz[i], tmpl.Query("#"+texts[i][0]).Has("clz-even"))
		assert.Equal(t, oddClz[i], tmpl.Query("#"+texts[i][0]).Has("clz-odd"))
	}

	// Test Text
	for i := range texts {
		assert.Equal(t, texts[i][1], tmpl.Text("#"+texts[i][0]))
	}

	// Test Has
	for i := range clz {
		assert.True(t, tmpl.Has("#"+texts[i][0], clz[i]))
	}
	for i := range evenClz {
		assert.Equal(t, evenClz[i], tmpl.Has("#"+texts[i][0], "clz-even"))
	}

	// Test Remove
	for i := range clz {
		tmpl.Remove("li", clz[i])
	}
	for i := range clz {
		assert.False(t, tmpl.Has("#"+texts[i][0], clz[i]))
	}

	// Test Add
	for i := range clz {
		tmpl.Add("#"+texts[i][0], clz[i])
	}
	for i := range clz {
		assert.True(t, tmpl.Has("#"+texts[i][0], clz[i]))
	}

	// Test Toggle
	tmpl.Toggle("li", "clz-even")
	for i := range evenClz {
		assert.Equal(t, !evenClz[i], tmpl.Has("#"+texts[i][0], "clz-even"))
	}
	tmpl.Toggle("li", "clz-even")

	// Test Replace
	tmpl.Replace("li", "clz-odd", "clz-odd-r")
	for i := range texts {
		assert.False(t, tmpl.Has("#"+texts[i][0], "clz-odd"))
	}
	tmpl.Add("#"+texts[1][0], "clz-odd")

	// Test Clone
	clone := tmpl.Clone()
	assert.Equal(t, 3, clone.Length()) // <p> x 2 + <ul> x 1
	assert.Equal(t, "tmpl_p_first", clone.First().Attr("id"))
	assert.Equal(t, "tmpl_p_last", clone.Last().Attr("id"))
	clone.Travel("li", func(i int, child js.Element) {
		assert.Equal(t, texts[i][0], child.Attr("id"))
		assert.Equal(t, texts[i][1], child.Text())
	})

	// Append to body, and tmpl is empty.
	js.Append(tmpl)

	assert.Equal(t, 0, tmpl.Length())

	for i := range texts {
		assert.Equal(t, texts[i][1], js.Query("#"+texts[i][0]).Text())
	}

	for i := range clz {
		assert.True(t, js.Query("#"+texts[i][0]).Has(clz[i]))
	}

	for i := range evenClz {
		assert.Equal(t, evenClz[i], js.Query("#"+texts[i][0]).Has("clz-even"))
		assert.Equal(t, oddClz[i], js.Query("#"+texts[i][0]).Has("clz-odd"))
	}

}
