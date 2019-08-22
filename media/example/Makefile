.PHONY: test

all:
	env GOOS=js GOARCH=wasm go build -o test.wasm
	goexec 'http.ListenAndServeTLS(":8080","web.crt", "web.key", http.FileServer(http.Dir(".")))'