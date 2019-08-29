package io

import "github.com/dairaga/js"

// Blob represents javascript blob object.
type Blob struct {
	ref js.Value
}

// NewBlob returns a blob object.
func NewBlob(raw []byte, mineType ...string) *Blob {
	typedArr := js.TypedArrayOf(raw)
	defer typedArr.Release()

	if len(mineType) > 0 {
		return &Blob{ref: js.New("Blob", []interface{}{typedArr}, map[string]interface{}{"type": mineType})}
	}

	return &Blob{ref: js.New("Blob", []interface{}{typedArr})}
}

// ----------------------------------------------------------------------------

// JSValue ...
func (b *Blob) JSValue() js.Value {
	return b.ref
}

// Type returns blob type property.
func (b *Blob) Type() string {
	return b.ref.Get("type").String()
}

// Size returns blob size in byte.
func (b *Blob) Size() int {
	return b.ref.Get("size").Int()
}
