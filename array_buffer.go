// +build js,wasm

package js

// Bytes convert javascript array buffer to Go byte slice.
func Bytes(arrayBuffer Value) []byte {
	srcArray := New("Uint8Array", arrayBuffer)
	return ToGoBytes(srcArray)
}
