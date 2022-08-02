//go:build js && wasm

package x

import (
	"testing"

	"github.com/dairaga/js/v3"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) FromValue(v js.Value) *Person {
	p.Name = v.Get("name").String()
	p.Age = v.Get("age").Int()
	return p
}

func (p *Person) ToValue() js.Value {
	return js.ValueOf(map[string]any{
		"name": p.Name,
		"age":  p.Age,
	})
}

func TestConvert(t *testing.T) {

}
