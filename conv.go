// go:build js && wasm

package js

import (
	"encoding/json"
	"syscall/js"
)

// -----------------------------------------------------------------------------

func Marshal(x any) (ret Value, err error) {
	ret = js.Null()
	var dataBytes []byte

	if dataBytes, err = json.Marshal(x); err != nil {
		return
	}
	ret = jsjson.Call("parse", string(dataBytes))
	return
}

// -----------------------------------------------------------------------------

func Unmarshal(data Value, x any) (err error) {
	objBytes := []byte(jsjson.Call("stringify", data).String())
	err = json.Unmarshal(objBytes, x)
	return
}
