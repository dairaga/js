//go:build js && wasm

package js

type Credential string

func (c Credential) String() string {
	return string(c)
}

const (
	Omit       Credential = "omit"
	SameOrigin Credential = "same-origin"
	Include    Credential = "include"
)
