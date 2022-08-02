//go:build js && wasm

package js

// TODO: cancel these.

import (
	"encoding/json"
)

// -----------------------------------------------------------------------------

func Marshal(x any) (ret Value, err error) {
	ret = Null()
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
