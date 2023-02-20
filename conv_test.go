//go:build js && wasm

package js

import (
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConv(t *testing.T) {
	type myTest struct {
		Name string
		Age  int
		OK   bool
		X    float64
		Y    int64
		Data *myTest
		Z    []*myTest
	}

	a := &myTest{
		Name: "abc",
		Age:  100,
		OK:   true,
		X:    123.456,
		Y:    123456789,
		Data: &myTest{
			Name: "def",
			Age:  10,
			OK:   false,
			X:    -123.456,
			Y:    -123456789,
			Data: nil,
			Z:    nil,
		},
		Z: []*myTest{
			{
				Name: "1",
				Age:  1,
				OK:   true,
			},
			{
				Name: "2",
				Age:  2,
				OK:   false,
			},
			{
				Name: "3",
				Age:  3,
				OK:   true,
			},
		},
	}

	testConv := func(t *testing.T, name string, src, tmp, ans any) {
		jsv, err := Marshal(src)
		if err != nil {
			t.Error(name, err)
			return
		}

		if err := Unmarshal(jsv, tmp); err != nil {
			t.Error(name, err)
			return
		}

		assert.Equal(t, tmp, ans, name)
	}

	testConv(t, "struct", a, new(myTest), a)

	boola := true
	boolb := false
	testConv(t, "boolean", boola, &boolb, &boola)

	stra := "ABC"
	strb := ""
	testConv(t, "string", stra, &strb, &stra)

	c := new(myTest)
	if err := Unmarshal(js.Null(), c); err != nil {
		t.Log("nil", err)
	} else {
		t.Log("test nil", c)
	}
}
