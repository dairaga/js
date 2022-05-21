//go:build js && wasm

package builtin

import "syscall/js"

var docuemnt = js.Global().Get("document")

func Query(selector string) js.Value
