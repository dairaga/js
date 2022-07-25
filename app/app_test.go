//go:build js && wasm

package app_test

import (
	"os"
	"os/signal"
	"strings"
	"testing"
	"time"

	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/app"
	"github.com/dairaga/js/v2/json"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	headless := os.Getenv("WASM_HEADLESS")
	exitVal := m.Run()

	if headless == "off" {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		// Block until a signal is received.
		<-c
	}

	os.Exit(exitVal)
}

// -----------------------------------------------------------------------------

type testState struct {
	URL   string
	Name  string
	Index int
}

// -----------------------------------------------------------------------------

func (s *testState) MarshalValue() (js.Value, error) {
	return js.ValueOf(map[string]any{
		"url":   s.URL,
		"name":  s.Name,
		"index": s.Index,
	}), nil
}

// -----------------------------------------------------------------------------

func (s *testState) UnmarshalValue(v js.Value) error {
	if err := json.ValidValue(v); err != nil {
		return err
	}
	s.URL = v.Get("url").String()
	s.Name = v.Get("name").String()
	s.Index = v.Get("index").Int()
	return nil
}

// -----------------------------------------------------------------------------

func TestMVVM(t *testing.T) {

	/*triggerA := false
	triggerB := false

	a := false
	b := []string{}

	app.Var(&a, "a", func(sender string, v bool) {
		triggerA = true
		assert.Equal(t, a, v)
	})

	app.Var(&b, "b", func(sender string, v []string) {
		triggerB = true
		assert.Equal(t, b, v)
	})

	a = true
	app.Trigger("_test_", "a")
	assert.True(t, triggerA)

	b = append(b, "A", "B", "C")
	app.Trigger("_test_", "b")
	assert.True(t, triggerB)*/
}

// -----------------------------------------------------------------------------

func TestState(t *testing.T) {
	curURL := app.URL().String()
	assert.Equal(t, js.Window().Get("location").Get("href").String(), curURL)

	data := []*testState{
		{
			URL:   "?q=1",
			Name:  "a1",
			Index: 0,
		},
		{
			URL:   "?q=2",
			Name:  "a2",
			Index: 1,
		},
		{
			URL:   "?q=3",
			Name:  "a3",
			Index: 2,
		},
	}

	for _, x := range data {
		assert.NoError(t, app.Push(x, x.URL))
		assert.Equal(t, js.Window().Get("location").Get("href").String(), app.URL().String())
		assert.True(t, strings.HasSuffix(app.URL().String(), x.URL))
		x1 := new(testState)
		assert.NoError(t, app.State(x1))
		assert.Equal(t, x, x1)
	}

	app.Go(-len(data))
	time.Sleep(500 * time.Millisecond) // need to wait browser changes url.
	assert.Equal(t, curURL, app.URL().String())

	for _, x := range data {
		app.Forward()
		time.Sleep(500 * time.Millisecond)
		assert.Equal(t, js.Window().Get("location").Get("href").String(), app.URL().String())
		assert.True(t, strings.HasSuffix(app.URL().String(), x.URL))
		x1 := new(testState)
		assert.NoError(t, app.State(x1))
		assert.Equal(t, x, x1)
	}

}

/*
import (
	"testing"
	"time"

	"github.com/dairaga/js/v2/url"
	"github.com/stretchr/testify/assert"
)

type myState struct {
	Name string
	Age  int
}

type serv struct {
	t   *testing.T
	ch  chan struct{}
	cur string
	old string
}

func (s *serv) Serve(oldURL, newURL url.URL) {
	s.old = oldURL.Hash()
	s.cur = newURL.Hash()
	s.ch <- struct{}{}
}

func TestHash(t *testing.T) {
	serv := &serv{
		t:  t,
		ch: make(chan struct{}, 1),
	}

	//curURL := url.New(js.Global().Get("location").Get("href").String())
	Init(serv)
	_app.handler.Serve(_app.currentURL, _app.currentURL)
	<-serv.ch
	assert.Equal(t, "", serv.old)
	assert.Equal(t, "", serv.cur)

	ChangeHash("#a100")
	<-serv.ch
	assert.Equal(t, "", serv.old)
	assert.Equal(t, "#a100", serv.cur)

	ChangeHash("#b100")
	<-serv.ch
	assert.Equal(t, "#a100", serv.old)
	assert.Equal(t, "#b100", serv.cur)

	ChangeHash("#b100")
	select {
	case <-time.After(1 * time.Second):
	case <-serv.ch:
	}
	// the hashchange event is not triggered.
	assert.Equal(t, "#a100", serv.old)
	assert.Equal(t, "#b100", serv.cur)

	s := &myState{
		Name: "abc",
		Age:  10,
	}

	s1 := new(myState)

	Push("/test_state#1", s)
	<-serv.ch
	assert.Equal(t, "#b100", serv.old)
	assert.Equal(t, "#1", serv.cur)
	assert.NoError(t, State(s1))
	assert.Equal(t, s, s1)

	s.Name = "def"
	s.Age = 20
	Push("/test_state#2", s)
	<-serv.ch
	assert.Equal(t, "#1", serv.old)
	assert.Equal(t, "#2", serv.cur)
	assert.NoError(t, State(s1))
	assert.Equal(t, s, s1)
}
*/
