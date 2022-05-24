//go:build js && wasm

package ajax

import (
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"testing"

	"github.com/dairaga/js/v2/xhr"
	"github.com/stretchr/testify/assert"
)

type respData struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

type category struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Parent uint64 `json:"parent"`
}

// -----------------------------------------------------------------------------

const timeout = 2000
const apiURL = `http://127.0.0.1:8080/api/v1`

// -----------------------------------------------------------------------------

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

func TestGet(t *testing.T) {
	ch := make(chan struct{})
	testURL := apiURL + "/categories"
	t.Log(testURL)
	cli, err := Get(testURL, func(resp *Response, err error) {
		assert.Nil(t, err)
		assert.Equal(t, `application/json; charset=UTF-8`, resp.Header("Content-Type"))

		data := new(respData)
		err = json.Unmarshal(resp.Body(), data)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, data.Code)

		categories := []category{}
		err = json.Unmarshal(data.Data, &categories)
		assert.Nil(t, err)
		assert.Equal(t, uint64(1), categories[0].ID)
		assert.Equal(t, `3C`, categories[0].Name)
		assert.Equal(t, uint64(0), categories[0].Parent)

		ch <- struct{}{}
	})
	assert.Nil(t, err)
	<-ch
	cli.Release()
	assert.True(t, cli.released)
	assert.Equal(t, xhr.ErrReleased, cli.Do(nil))

}
