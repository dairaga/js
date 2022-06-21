.PHONY: clean test app builtin url xhr xhr/ajax example mvvm xhr/upload

GOPATH=$(shell go env GOPATH)
WASMEXEC=${GOPATH}/bin/wasmbrowsertest
WASM_HEADLESS=on

all:
	env WASM_HEADLESS=${WASM_HEADLESS} GOOS=js GOARCH=wasm go test -p 1 -timeout 0 . -exec=${WASMEXEC} -test.v

app:
	env WASM_HEADLESS=${WASM_HEADLESS} GOOS=js GOARCH=wasm go test -p 1 -timeout 0 github.com/dairaga/js/v2/app -exec=${WASMEXEC} -test.v

mvvm:
	env WASM_HEADLESS=${WASM_HEADLESS} GOOS=js GOARCH=wasm go test -p 1 -timeout 0 github.com/dairaga/js/v2/mvvm -exec=${WASMEXEC} -test.v

builtin:
	env WASM_HEADLESS=${WASM_HEADLESS} GOOS=js GOARCH=wasm go test -p 1 -timeout 0 github.com/dairaga/js/v2/builtin -exec=${WASMEXEC} -test.v

url:
	env WASM_HEADLESS=${WASM_HEADLESS} GOOS=js GOARCH=wasm go test -p 1 -timeout 0 github.com/dairaga/js/v2/url -exec=${WASMEXEC} -test.v

xhr/ajax:
	env WASM_HEADLESS=${WASM_HEADLESS} GOOS=js GOARCH=wasm go test -p 1 -timeout 0 github.com/dairaga/js/v2/xhr/ajax -exec=${WASMEXEC} -test.v

xhr/upload:
	env WASM_HEADLESS=off GOOS=js GOARCH=wasm go test -p 1 -timeout 0 github.com/dairaga/js/v2/xhr/upload -exec=${WASMEXEC} -test.v

example:
	env GOOS=js GOARCH=wasm go build -o example/wasm.wasm github.com/dairaga/js/v2/example
	tinygo build -o example/tiny.wasm -target wasm github.com/dairaga/js/v2/example

doc:
	env GOOS=js GOARCH=wasm godoc -http=:6060