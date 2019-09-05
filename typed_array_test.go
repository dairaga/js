// +build js,wasm

package js

import (
	"reflect"
	"testing"
)

func verifyTypedArray(t *testing.T, x interface{}) {
	xval := reflect.ValueOf(x)
	jsv := TypedArrayOf(x)
	if xval.Len() != jsv.Length() {
		t.Errorf("length not match, src: %d, dest: %v", xval.Len(), jsv.Length())
	}

	switch xval.Type().Elem().Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32:
		for i := 0; i < xval.Len(); i++ {
			if xval.Index(i).Int() != int64(jsv.Index(i).Int()) {
				t.Errorf("index %d must be %v but %v", i, xval.Index(i).Int(), jsv.Index(i).Int())
			}
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		for i := 0; i < xval.Len(); i++ {
			if xval.Index(i).Uint() != uint64(jsv.Index(i).Int()) {
				t.Errorf("index %d must be %v but %v", i, xval.Index(i).Uint(), jsv.Index(i).Int())
			}
		}
	case reflect.Float32, reflect.Float64:
		for i := 0; i < xval.Len(); i++ {
			if xval.Index(i).Float() != float64(jsv.Index(i).Float()) {
				t.Errorf("index %d must be %v but %v", i, xval.Index(i).Float(), jsv.Index(i).Float())
			}
		}
	default:
		t.Errorf("%v not support", xval.Type())
	}
}

func TestTypedArray(t *testing.T) {
	n8 := []int8{1, -2, 3}
	verifyTypedArray(t, n8)

	n16 := []int16{1, -2, 3}
	verifyTypedArray(t, n16)

	n32 := []int32{1, -2, 3}
	verifyTypedArray(t, n32)

	un8 := []byte{1, 2, 3}
	verifyTypedArray(t, un8)

	un16 := []uint16{1, 2, 3}
	verifyTypedArray(t, un16)

	un32 := []uint32{1, 2, 3}
	verifyTypedArray(t, un32)

	f32 := []float32{1.1, -2.2, 3.3}
	verifyTypedArray(t, f32)

	f64 := []float64{1.1, -2.2, 3.3}
	verifyTypedArray(t, f64)
}

func TestTypedArrayToGo(t *testing.T) {
	n8 := uint8Array.Call("of", 1, 2, 3)

	t.Log(n8)

	for i := 0; i < n8.Length(); i++ {
		t.Log(i, n8.Index(i))
	}

	tmp := make([]byte, 3)

	x := TypedArrayOf(tmp)

	x.Call("set", n8)

	t.Log(tmp)
}
