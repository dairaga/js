package io

import (
	"time"

	"github.com/dairaga/js"
)

// File represents javascript File. https://developer.mozilla.org/en-US/docs/Web/API/File
type File struct {
	ref js.Value
}

// JSValue ...
func (f *File) JSValue() js.Value {
	return f.ref
}

// FileOf returns a File object.
func FileOf(x interface{}) *File {
	return &File{ref: js.ValueOf(x)}
}

// Size https://developer.mozilla.org/en-US/docs/Web/API/Blob/size
func (f *File) Size() int {
	return f.ref.Get("size").Int()
}

// FileType https://developer.mozilla.org/en-US/docs/Web/API/File/type
func (f *File) FileType() string {
	return f.ref.Get("type").String()
}

// Name https://developer.mozilla.org/en-US/docs/Web/API/File/name
func (f *File) Name() string {
	return f.ref.Get("name").String()
}

// LastModified https://developer.mozilla.org/en-US/docs/Web/API/File/lastModified
func (f *File) LastModified() time.Time {
	m := int64(f.ref.Get("lastModified").Int())
	return time.Unix(0, m*int64(time.Millisecond))
}
