//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type MyData struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

func main() {
	fmt.Println("hello world")

	data := &MyData{
		ID:      1,
		Name:    "test",
		Created: time.Now(),
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(dataBytes))
	}

	_ = reflect.ValueOf(100)
}
