//go:build js && wasm

package js

import (
	"math/rand"
	"time"
	"unsafe"
)

/*
	Reference from https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
*/

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits

	_tattoo    = `data-dairaga`
	_tattooLen = 10
)

var (
	randSrc = rand.NewSource(time.Now().UnixNano())
)

// -----------------------------------------------------------------------------

func tattoo(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, randSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// -----------------------------------------------------------------------------

func tattoos(target Value) Value {
	val := target.Call("getAttribute", _tattoo)
	if val.Truthy() {
		return target
	}
	target.Call("setAttribute", _tattoo, tattoo(_tattooLen))
	return target
}
