//go:build js && wasm

package js

import (
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

	jsa, err := Marshal(a)
	if err != nil {
		t.Error(err)
		return
	}

	b := new(myTest)

	if err := Unmarshal(jsa, b); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, a, b)
}
