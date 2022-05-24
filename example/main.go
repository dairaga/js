//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/dairaga/js/v2/app"
	"github.com/dairaga/js/v2/url"
)

type MyData struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

// -----------------------------------------------------------------------------

type serv struct {
	cur string
	old string
}

// -----------------------------------------------------------------------------

func (s *serv) Serve(curURL url.URL, curHash, oldHash string) {
	s.old = oldHash
	s.cur = curHash
}

// -----------------------------------------------------------------------------

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

	time.Now()
	app.Init(&serv{})
	app.Start()
}
