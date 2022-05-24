//go:build js && wasm

package app

import (
	"fmt"
	"reflect"
)

var models map[string]reflect.Value
var triggers map[string][]reflect.Value

// -----------------------------------------------------------------------------

func initMVVM() {
	models = make(map[string]reflect.Value)
	triggers = make(map[string][]reflect.Value)
}

// -----------------------------------------------------------------------------

func Add(name string, x any) {
	v := reflect.ValueOf(x)
	if reflect.Ptr != v.Kind() {
		panic(fmt.Sprintf("x must be ptr, but %v", v.Kind()))
	}

	_, ok := models[name]
	if ok {
		panic(fmt.Sprintf("%s existed", name))
	}
	models[name] = v
}

// -----------------------------------------------------------------------------

func Remove(name string) {
	delete(models, name)
}

// -----------------------------------------------------------------------------

func Trigger(sender, name string) {
	val, ok := models[name]
	if !ok {
		return
	}

	callbacks := triggers[name]
	size := len(callbacks)
	if size > 0 {
		args := []reflect.Value{reflect.ValueOf(sender), val.Elem()}
		for i := 0; i < size; i++ {
			callbacks[i].Call(args)
		}
	}
}
