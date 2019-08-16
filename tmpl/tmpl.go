package tmpl

import (
	"fmt"
	"html/template"
	"strings"
)

// HTML executes golang template.
func HTML(tmpl string, data interface{}) string {
	t, err := template.New("_dairaga_js_").Parse(tmpl)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	sb := &strings.Builder{}
	t.Execute(sb, data)

	return sb.String()
}
