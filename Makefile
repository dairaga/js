.PHONY: clean test app builtin url

GOPATH=$(shell go env GOPATH)
WASMEXEC=${GOPATH}/bin/wasmbrowsertest
WASM_HEADLESS=on

all:
	env WASM_HEADLESS=${WASM_HEADLESS} GOOS=js GOARCH=wasm go test -p 1 . -exec=${WASMEXEC} -test.v

app:
	env WASM_HEADLESS=${WASM_HEADLESS} GOOS=js GOARCH=wasm go test -p 1 github.com/dairaga/js/v2/app -exec=${WASMEXEC} -test.v

builtin:
	env WASM_HEADLESS=${WASM_HEADLESS} GOOS=js GOARCH=wasm go test -p 1 github.com/dairaga/js/v2/builtin -exec=${WASMEXEC} -test.v

url:
	env WASM_HEADLESS=${WASM_HEADLESS} GOOS=js GOARCH=wasm go test -p 1 github.com/dairaga/js/v2/url -exec=${WASMEXEC} -test.v

