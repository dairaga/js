//go:build js && wasm

// Package 主要目的，是使用 Golang 來開發前端功能，達到前/後端分離。
// 封裝常用的 DOM 功能，Javascript 常用的 Class 與 Function；並實作 MVVM 架構。
package js

import (
	"syscall/js"

	"github.com/dairaga/js/v3/builtin"
)

const (
	TypeUndefined = js.TypeUndefined
	TypeNull      = js.TypeNull
	TypeBoolean   = js.TypeBoolean
	TypeNumber    = js.TypeNumber
	TypeString    = js.TypeString
	TypeSymbol    = js.TypeSymbol
	TypeObject    = js.TypeObject
	TypeFunction  = js.TypeFunction
)

// -----------------------------------------------------------------------------

var (
	global    = js.Global()            // global 物件，也是 Javascript 的 window 物件。
	document  = global.Get("document") // Document 物件。
	body      = document.Get("body")   // Document 下的 Body。
	null      = js.Null()              // Javascript null。
	undefined = js.Undefined()         // Javascript undefined。
	jsjson    = global.Get("JSON")     // Javascript JSON.
)

// -----------------------------------------------------------------------------

// alias syscall/js 的物件，讓 import 比較方便
type (
	Value      = js.Value
	Type       = js.Type
	Func       = js.Func
	JSFunc     = func(Value, []Value) any
	ValueError = js.ValueError
	Error      = js.Error

	Obj = map[string]any
)

// -----------------------------------------------------------------------------

// Wrapper 原 golang 的 Wrapper interface，覺得很好用，保留。
type Wrapper interface {
	// JSValue returns Javascript value.
	JSValue() Value
}

// -----------------------------------------------------------------------------

// ValueOf 仿 syscall/js 的 ValueOf。x 可以是 Wrapper 或是原生 js.Value，否則則以 syscall/js 的定義為準。
func ValueOf(x any) Value {
	switch v := x.(type) {
	case Appendable:
		return v.Ref()
	case Wrapper:
		return v.JSValue()
	case Value:
		return v
	default:
		return js.ValueOf(x)
	}
}

// -----------------------------------------------------------------------------

// GoBytes 將 Javascript 的 Uint8Array 轉成 golang 的 []byte。src 必須是 Javascript 的 Uint8Array。
func GoBytes(src Value) []byte {
	if !builtin.Uint8Array.Is(src) {
		panic("src is not an Uint8Array")
	}

	size := src.Get("byteLength").Int()
	dst := make([]byte, size)
	js.CopyBytesToGo(dst, src)
	return dst
}

// -----------------------------------------------------------------------------

// Uint8Array 將 golang 的 []byte 轉成 Javascript 的 Uint8Array。
func Uint8Array(src []byte) Value {
	dst := builtin.Uint8Array.New(len(src))
	js.CopyBytesToJS(dst, src)
	return dst
}

// -----------------------------------------------------------------------------

// ArrayBufferToBytes 將 Javascript 的 ArrayBuffer 轉成 golang []byte。src 必須是 Javascript 的 ArrayBuffer。
func ArrayBufferToBytes(src Value) []byte {
	if !builtin.ArrayBuffer.Is(src) {
		panic("src is not an ArrayBuffer")
	}

	return GoBytes(builtin.Uint8Array.New(src))
}

// -----------------------------------------------------------------------------

func Float32Array(src js.Value) []float32 {
	if !builtin.Float32Array.Is(src) {
		panic(ValueError{
			Method: "FromFloat32Array",
			Type:   src.Type(),
		})
	}

	size := src.Length()
	ret := make([]float32, size)
	for i := 0; i < size; i++ {
		ret[i] = float32(src.Index(i).Float())
	}
	return ret
}

// -----------------------------------------------------------------------------

// Null 回傳 Javascript 的 null。
func Null() Value {
	return null
}

// -----------------------------------------------------------------------------

// Undefined 回傳 Javascript 的 Undefined。
func Undefined() Value {
	return undefined
}

// -----------------------------------------------------------------------------

func EncodeURI(src string) string {
	return global.Call("encodeURIComponent", src).String()
}

// -----------------------------------------------------------------------------

func DecodeURI(src string) string {
	return global.Call("decodeURIComponent", src).String()
}

// -----------------------------------------------------------------------------

func ParseInt(src string, base int) (int, bool) {
	result := global.Call("parseInt", src, base)

	if result.Truthy() {
		return result.Int(), true
	} else {
		return 0, false
	}
}

// -----------------------------------------------------------------------------

func ParseFloat(src string) (float64, bool) {
	result := global.Call("parseFloat", src)

	if result.Truthy() {
		return result.Float(), true
	} else {
		return 0, false
	}
}
