//go:build js && wasm

// Package url provides datastructure and functions about javascript URL.
package url

import (
	"fmt"

	"github.com/dairaga/js/v2"
	"github.com/dairaga/js/v2/builtin"
)

type URL js.Value

func (u URL) JSValue() js.Value {
	return js.Value(u)
}

// -----------------------------------------------------------------------------

func (u URL) Hash() string {
	return js.Value(u).Get("hash").String()
}

// -----------------------------------------------------------------------------

func (u URL) SetHash(val string) {
	js.Value(u).Set("hash", val)
}

// -----------------------------------------------------------------------------

func (u URL) Host() string {
	return js.Value(u).Get("host").String()
}

// -----------------------------------------------------------------------------

func (u URL) SetHost(val string) {
	js.Value(u).Set("host", val)
}

// -----------------------------------------------------------------------------

// Hostname returns hostname property.
func (u URL) Hostname() string {
	return js.Value(u).Get("hostname").String()
}

// -----------------------------------------------------------------------------

// SetHostname sets hostname property.
func (u URL) SetHostname(val string) {
	js.Value(u).Set("hostname", val)
}

// -----------------------------------------------------------------------------

// Href returns href property.
func (u URL) Href() string {
	return js.Value(u).Get("href").String()
}

// -----------------------------------------------------------------------------

// SetHref sets href property.
func (u URL) SetHref(val string) {
	js.Value(u).Set("href", val)
}

// -----------------------------------------------------------------------------

// Origin returns origin property.
func (u URL) Origin() string {
	return js.Value(u).Get("origin").String()
}

// -----------------------------------------------------------------------------

// Password returns password property.
func (u URL) Password() string {
	return js.Value(u).Get("password").String()
}

// -----------------------------------------------------------------------------

// SetPassword sets password property.
func (u URL) SetPassword(val string) {
	js.Value(u).Set("password", val)
}

// -----------------------------------------------------------------------------

// Pathname returns pathname property.
func (u URL) Pathname() string {
	return js.Value(u).Get("pathname").String()
}

// -----------------------------------------------------------------------------

func (u URL) QueryPath() string {
	return u.Pathname()
}

// -----------------------------------------------------------------------------

// SetPathname sets pathname property.
func (u URL) SetPathname(val string) {
	js.Value(u).Set("pathname", val)
}

// -----------------------------------------------------------------------------

func (u URL) SetQueryPath(path string) {
	u.SetPathname(path)
}

// -----------------------------------------------------------------------------

// Port returns port property.
func (u URL) Port() string {
	return js.Value(u).Get("port").String()
}

// -----------------------------------------------------------------------------

// SetPort sets port property.
func (u URL) SetPort(port int) {
	js.Value(u).Set("port", fmt.Sprintf("%d", port))
}

// -----------------------------------------------------------------------------

// Protocol returns protocol property and remove ":" at end.
func (u URL) Protocol() string {
	return js.Value(u).Get("protocol").String()
}

// -----------------------------------------------------------------------------

// SetProtocol sets protocol property. Append ":" automatically if protocol is not end with ":".
func (u URL) SetProtocol(protocol string) {
	if protocol[len(protocol)-1] != ':' {
		protocol += ":"
	}

	js.Value(u).Set("protocol", protocol)
}

// -----------------------------------------------------------------------------

// Search returns search property.
func (u URL) Search() string {
	return js.Value(u).Get("search").String()
}

// -----------------------------------------------------------------------------

// Querystring returns query string in the url.
func (u URL) QueryString() string {
	return u.Search()
}

// SetSearch sets search property.
func (u URL) SetSearch(search string) {
	js.Value(u).Set("search", search)
}

// -----------------------------------------------------------------------------

// SetQuerystring sets query string in the url.
func (u URL) SetQueryString(query string) {
	u.SetSearch(query)
}

// -----------------------------------------------------------------------------

// Username returns username property.
func (u URL) Username() string {
	return js.Value(u).Get("username").String()
}

// -----------------------------------------------------------------------------

// SetUsername sets username property.
func (u URL) SetUsername(username string) {
	js.Value(u).Set("username", username)
}

// -----------------------------------------------------------------------------

// Params returns parameters in url.
func (u URL) Params() Params {
	return Params(js.Value(u).Get("searchParams"))
}

// -----------------------------------------------------------------------------

func (u URL) String() string {
	return js.Value(u).Call("toString").String()
}

// -----------------------------------------------------------------------------

func (u URL) JSON() string {
	return js.Value(u).Call("toJSON").String()
}

// -----------------------------------------------------------------------------

func New(url string) URL {
	return URL(builtin.URL.New(url))
}

// TODO: createObjectURL, revokeObjectURL()
