package io

import (
	"time"

	"syscall/js"
)

// File ...
type File struct {
	ref js.Value
}

// JSValue ...
func (f File) JSValue() js.Value {
	return f.ref
}

// FileOf ...
func FileOf(x interface{}) File {
	return File{ref: js.ValueOf(x)}
}

// Size ...
func (f File) Size() int {
	return f.ref.Get("size").Int()
}

// FileType ...
func (f File) FileType() string {
	return f.ref.Get("type").String()
}

// Name ...
func (f File) Name() string {
	return f.ref.Get("name").String()
}

// LastModified ...
func (f File) LastModified() time.Time {
	m := int64(f.ref.Get("lastModified").Int())
	return time.Unix(0, m*int64(time.Millisecond))
}
