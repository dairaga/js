.PHONY: clean test document xhr

GOPATH=$(shell go env GOPATH)
WASMEXEC=${GOPATH}/bin/wasmbrowsertest
WASM_HEADLESS=on

all:
	env WASM_HEADLESS=${WASM_HEADLESS} GOOS=js GOARCH=wasm go test -p 1 . -exec=${WASMEXEC} -test.v