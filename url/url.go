//go:build js && wasm

// Package url provides common data structures and functions about javascript URL.
package url

import (
	"fmt"
	"strings"

	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

// URL is javascript URL.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL.
type URL js.Value

// JSValue returns javascript value.
func (u URL) JSValue() js.Value {
	return js.Value(u)
}

// -----------------------------------------------------------------------------

// Hash returns hash string containing a '#' in the url.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/hash.
func (u URL) Hash() string {
	return js.Value(u).Get("hash").String()
}

// -----------------------------------------------------------------------------

// SetHash sets a new hash string to the url. Given val must prefix '#'.
func (u URL) SetHash(val string) {
	js.Value(u).Set("hash", val)
}

// -----------------------------------------------------------------------------

// Host returns host string in the url. Host contains hostname and port.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/host.
func (u URL) Host() string {
	return js.Value(u).Get("host").String()
}

// -----------------------------------------------------------------------------

// SetHost sets a new host string to the url.
func (u URL) SetHost(val string) {
	js.Value(u).Set("host", val)
}

// -----------------------------------------------------------------------------

// Hostname returns the domain name in the url.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/hostname.
func (u URL) Hostname() string {
	return js.Value(u).Get("hostname").String()
}

// -----------------------------------------------------------------------------

// SetHostname sets new domain name to the url.
func (u URL) SetHostname(val string) {
	js.Value(u).Set("hostname", val)
}

// -----------------------------------------------------------------------------

// Href returns whole string in the url.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/href.
func (u URL) Href() string {
	return js.Value(u).Get("href").String()
}

// -----------------------------------------------------------------------------

// SetHref sets a new whole url string to the url.
func (u URL) SetHref(val string) {
	js.Value(u).Set("href", val)
}

// -----------------------------------------------------------------------------

// Origin returns returns a string containing the Unicode serialization of the origin of the url.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/origin.
func (u URL) Origin() string {
	return js.Value(u).Get("origin").String()
}

// -----------------------------------------------------------------------------

// Password returns the password string in the url.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/password.
func (u URL) Password() string {
	return js.Value(u).Get("password").String()
}

// -----------------------------------------------------------------------------

// SetPassword sets a new password to the url.
func (u URL) SetPassword(val string) {
	js.Value(u).Set("password", val)
}

// -----------------------------------------------------------------------------

// Pathname returns pathname string in the url.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/pathname.
//func (u URL) Pathname() string {
//	return js.Value(u).Get("pathname").String()
//}

// -----------------------------------------------------------------------------

// QueryPath returns query path string has prefix '/' in the url.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/pathname.
func (u URL) QueryPath() string {
	return js.Value(u).Get("pathname").String()
}

// -----------------------------------------------------------------------------

// SetPathname sets a new pathname to url.
//func (u URL) SetPathname(val string) {
//	js.Value(u).Set("pathname", val)
//}

// -----------------------------------------------------------------------------

// SetQueryPath sets a new query path to the url.
func (u URL) SetQueryPath(path string) {
	js.Value(u).Set("pathname", path)
}

// -----------------------------------------------------------------------------

// Port returns port (in string) in the url.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/port.
func (u URL) Port() string {
	return js.Value(u).Get("port").String()
}

// -----------------------------------------------------------------------------

// SetPort sets new port to the url.
func (u URL) SetPort(port int) {
	js.Value(u).Set("port", fmt.Sprintf("%d", port))
}

// -----------------------------------------------------------------------------

// Protocol returns protocol string has suffix ':' in the url.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/protocol.
func (u URL) Protocol() string {
	return js.Value(u).Get("protocol").String()
}

// -----------------------------------------------------------------------------

// SetProtocol sets new protocol suffix ':' automatically to the url.
func (u URL) SetProtocol(protocol string) {
	if protocol[len(protocol)-1] != ':' {
		protocol += ":"
	}

	js.Value(u).Set("protocol", protocol)
}

// -----------------------------------------------------------------------------

// Search returns search property.
//func (u URL) Search() string {
//	return js.Value(u).Get("search").String()
//}

// -----------------------------------------------------------------------------

// Querystring returns query string has prefix '?' in the url.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/search.
func (u URL) QueryString() string {
	return js.Value(u).Get("search").String()
}

// SetSearch sets search property.
//func (u URL) SetSearch(search string) {
//	js.Value(u).Set("search", search)
//}

// -----------------------------------------------------------------------------

// SetQuerystring sets query string prefix '?' automatically to the url.
func (u URL) SetQueryString(query string) {
	if !strings.HasPrefix(query, "?") {
		query = "?" + query
	}
	js.Value(u).Set("search", query)
}

// -----------------------------------------------------------------------------

// Username returns user name string in the url.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/username.
func (u URL) Username() string {
	return js.Value(u).Get("username").String()
}

// -----------------------------------------------------------------------------

// SetUsername sets new user name string to the url.
func (u URL) SetUsername(username string) {
	js.Value(u).Set("username", username)
}

// -----------------------------------------------------------------------------

// Params returns all parameters in the url. The resulted is a javascript URLSearchParams.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/searchParams.
func (u URL) Params() Params {
	return Params(js.Value(u).Get("searchParams"))
}

// -----------------------------------------------------------------------------

// String invokes javascript URL.toString method and returns whole url string like Href.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/toString.
func (u URL) String() string {
	return js.Value(u).Call("toString").String()
}

// -----------------------------------------------------------------------------

// JSON invokes javascript URL.toJSON method and returns a serialized string of the url.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/toJSON.
func (u URL) JSON() string {
	return js.Value(u).Call("toJSON").String()
}

// -----------------------------------------------------------------------------

// New returns a new URL object with given string url.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/URL.
func New(url string) URL {
	return URL(builtin.URL.New(url))
}

// -----------------------------------------------------------------------------

// CreateObjectURL creates a url string prepresenting the given object x.
// The given object x must be Blob, File, or MediaSource.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/createObjectURL.
func CreateObjectURL(x any) string {
	var resulted js.Value

	switch v := x.(type) {
	case js.Wrapper:
		resulted = v.JSValue()
	case js.Value:
		resulted = v
	default:
		panic(fmt.Sprintf("unsupported type %T", x))
	}

	if !builtin.IsMediaSource(resulted) && !builtin.IsBlob(resulted) && !builtin.IsFile(resulted) {
		panic(fmt.Sprintf("unsupported type %v", resulted.Type()))
	}

	return builtin.URL.JSValue().Call("createObjectURL", resulted).String()
}

// -----------------------------------------------------------------------------

// RevokeObjectURL revokes a URL previously created by CreateObjectURL.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/URL/revokeObjectURL.
func RevokeObjectURL(url string) {
	builtin.URL.JSValue().Call("revokeObjectURL", url)
}
