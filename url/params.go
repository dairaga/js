//go:build js && wasm

package url

import "github.com/dairaga/js/v2"

type Params js.Value

func (p Params) JSValue() js.Value {
	return js.Value(p)
}

// -----------------------------------------------------------------------------

func (p Params) String() string {
	return p.JSValue().Call("toString").String()
}

// -----------------------------------------------------------------------------

func (p Params) Has(name string) bool {
	return p.JSValue().Call("has", name).Bool()
}

// -----------------------------------------------------------------------------

func (p Params) Get(name string) (val string, ok bool) {
	x := p.JSValue().Call("get", name)
	if ok = !x.IsNull(); ok {
		val = x.String()
	}
	return
}

// -----------------------------------------------------------------------------

func (p Params) GetAll(name string) []string {
	x := p.JSValue().Call("getAll", name)
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
	p.JSValue().Call("set", name, val)
}

// -----------------------------------------------------------------------------

func (p Params) Append(name, val string) {
	p.JSValue().Call("append", name, val)
}

// -----------------------------------------------------------------------------

func (p Params) Delete(name string) {
	p.JSValue().Call("delete", name)
}

// -----------------------------------------------------------------------------

func (p Params) Foreach(fn func(key, val string)) {
	cb := js.FuncOf(func(_this js.Value, args []js.Value) any {
		fn(args[1].String(), args[0].String())
		return nil
	})

	p.JSValue().Call("forEach", cb)
	cb.Release()
}
