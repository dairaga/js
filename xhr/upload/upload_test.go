//go:build js && wasm

package upload

import (
	"fmt"
	"os"
	"os/signal"
	"testing"

	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/form"
	"github.com/dairaga/js/v2/xhr"
)

const uploadURL = `http://127.0.0.1:8080/upload`

// -----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	headless := os.Getenv("WASM_HEADLESS")
	defaultWithCredentials = false
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

type testHander struct{}

// -----------------------------------------------------------------------------

func (h *testHander) Progress(jobID string, state ProgressState, loaded, total uint64) {
	fmt.Printf("%s: (%v) %d/%d\n", jobID, state, loaded, total)
}

// -----------------------------------------------------------------------------

func (h *testHander) Completed(resp *xhr.Response, err error) {
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(string(resp.Body()))
}

// -----------------------------------------------------------------------------

func TestUpload(t *testing.T) {
	tmpl := `
	<form id="myform" method="post" enctype="multipart/form-data" action='http://127.0.0.1:8080/upload'>
		<div>test job id:<input type='text' name='job_id' id='job_id' /></div>
		<div>name: <input type='text' name='name' id='name' /></div>
		<div><input type="file" id="file_picker" name="file_picker" /></div>
		<input type='submit' value='ok' />
	</form>
	`

	elm := js.ElementOf(js.HTML(tmpl))
	js.Append(elm)

	btn := js.CreateElement("button")

	cli := New(&testHander{})

	btn.SetText("upload")
	btn.OnClick(func(_ js.Element, _ js.Event) {
		f := form.FormDataOf("#myform")
		cli.Upload(uploadURL, "xxx", f)
	})

	js.Append(btn)
}
