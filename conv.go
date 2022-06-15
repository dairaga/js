// go:build js && wasm

package js

import (
	"encoding/json"
	"syscall/js"

	"github.com/dairaga/js/v2/builtin"
)

func toAny(v Value) any {
	switch v.Type() {
	case js.TypeBoolean:
		return v.Bool()
	case js.TypeNumber:
		return v.Float()
	case js.TypeString:
		return v.String()
	case js.TypeNull, js.TypeUndefined, js.TypeFunction, js.TypeSymbol:
		return nil
	}

	if builtin.IsArray(v) {
		size := v.Length()
		ret := make([]any, size)
		for i := 0; i < size; i++ {
			ret[i] = toAny(v.Index(i))
		}
		return ret
	}

	if v.Type() == js.TypeObject {
		ret := make(map[string]any)
		keys := object.Call("keys", v)
		size := keys.Length()
		for i := 0; i < size; i++ {
			prop := keys.Index(i).String()
			ret[prop] = toAny(v.Get(prop))
		}
		return ret
	}

	panic(js.ValueError{Method: "js.Unmarshal", Type: v.Type()})
}

// -----------------------------------------------------------------------------

func Marshal(x any) (ret Value, err error) {
	ret = js.Null()
	var dataBytes []byte

	if dataBytes, err = json.Marshal(x); err != nil {
		return
	}

	obj := make(map[string]any)

	if err = json.Unmarshal(dataBytes, &obj); err != nil {
		return
	}

	ret = ValueOf(obj)
	return
}

// -----------------------------------------------------------------------------

func Unmarshal(data Value, x any) (err error) {
	obj := toAny(data)
	var objBytes []byte
	if objBytes, err = json.Marshal(obj); err != nil {
		return
	}
	err = json.Unmarshal(objBytes, x)
	return
}
