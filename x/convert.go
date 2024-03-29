//go:build js && wasm

package x

import "github.com/dairaga/js/v2"

type Valuer interface {
	FromValue(js.Value)
	ToValue() js.Value
}
