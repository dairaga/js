//go:build js && wasm

package url

import "github.com/dairaga/js/v2"

// Params is javascript URLSearchParams.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URLSearchParams.
type Params js.Value

// JSValue returns javascript value.
func (p Params) JSValue() js.Value {
	return js.Value(p)
}

// -----------------------------------------------------------------------------

// String returns query string.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URLSearchParams/toString.
func (p Params) String() string {
	return js.Value(p).Call("toString").String()
}

// -----------------------------------------------------------------------------

// Has returns true if given name is in Params.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URLSearchParams/has.
func (p Params) Has(name string) bool {
	return js.Value(p).Call("has", name).Bool()
}

// -----------------------------------------------------------------------------

// Get returns the first value of the given name.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URLSearchParams/get.
func (p Params) Get(name string) (val string, ok bool) {
	x := js.Value(p).Call("get", name)
	if ok = !x.IsNull(); ok {
		val = x.String()
	}
	return
}

// -----------------------------------------------------------------------------

// GetAll returns all values associated with a given name.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URLSearchParams/getAll.
func (p Params) GetAll(name string) []string {
	x := js.Value(p).Call("getAll", name)
	if !x.Truthy() {
		return []string{}
	}

	size := x.Length()
	ret := make([]string, size)
	for i := 0; i < size; i++ {
		ret[i] = x.Index(i).String()
	}
	return ret
}

// -----------------------------------------------------------------------------

// Set set value associated with a given name to given value val.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URLSearchParams/set.
func (p Params) Set(name, val string) {
	js.Value(p).Call("set", name, val)
}

// -----------------------------------------------------------------------------

// Append appends a given key/value pair name and val as a new parameter.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URLSearchParams/append.
func (p Params) Append(name, val string) {
	js.Value(p).Call("append", name, val)
}

// -----------------------------------------------------------------------------

// Delete remvoes all values associated with a given name.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URLSearchParams/delete.
func (p Params) Delete(name string) {
	js.Value(p).Call("delete", name)
}

// -----------------------------------------------------------------------------

// Foreach travels all values in parameters.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URLSearchParams/forEach.
func (p Params) Foreach(fn func(key, val string)) {
	cb := js.FuncOf(func(_this js.Value, args []js.Value) any {
		fn(args[1].String(), args[0].String())
		return nil
	})

	js.Value(p).Call("forEach", cb)
	cb.Release()
}
