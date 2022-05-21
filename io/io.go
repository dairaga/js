// go:build js && wasm

package io

import "syscall/js"

func convToJSValue(x any) js.Value {
	// TODO: move code BlobOf to here
	return js.Value{}
}
