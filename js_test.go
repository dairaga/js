//go:build js && wasm

package js_test

import (
	"os"
	"os/signal"
	"testing"
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
