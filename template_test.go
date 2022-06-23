//go:build js && wasm

package js_test

import (
	"strconv"
	"testing"

	"github.com/dairaga/js/v2"
)

func TestTemlateCreate(t *testing.T) {
	content := js.HTML(`<ul><li></li><li></li><li></li></ul>`)

	tmpl := js.CreateTemplate(content)

	tmpl.Set("li", func(idx int, child js.Element) {
		child.SetText("tmpl-" + strconv.Itoa(idx+1))
		//addClz(child.JSValue(), "clz-"+strconv.Itoa(idx+1))
		child.Add("clz-" + strconv.Itoa(idx+1))
	})

	js.Append(tmpl)

}
