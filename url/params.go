//go:build js && wasm

package url

import "github.com/dairaga/js/v2"

type Params js.Value

func (p Params) JSValue() js.Value {
	return js.Value(p)
}

// -----------------------------------------------------------------------------

func (p Params) String() string {
	return js.Value(p).Call("toString").String()
}

// -----------------------------------------------------------------------------

func (p Params) Has(name string) bool {
	return js.Value(p).Call("has", name).Bool()
}

// -----------------------------------------------------------------------------

func (p Params) Get(name string) (val string, ok bool) {
	x := js.Value(p).Call("get", name)
	if ok = !x.IsNull(); ok {
		val = x.String()
	}
	return
}

// -----------------------------------------------------------------------------

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

func (p Params) Set(name, val string) {
	js.Value(p).Call("set", name, val)
}

// -----------------------------------------------------------------------------

func (p Params) Append(name, val string) {
	js.Value(p).Call("append", name, val)
}

// -----------------------------------------------------------------------------

func (p Params) Delete(name string) {
	js.Value(p).Call("delete", name)
}

// -----------------------------------------------------------------------------

func (p Params) Foreach(fn func(key, val string)) {
	cb := js.FuncOf(func(_this js.Value, args []js.Value) any {
		fn(args[1].String(), args[0].String())
		return nil
	})

	js.Value(p).Call("forEach", cb)
	cb.Release()
}
