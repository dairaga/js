//go:build js && wasm

package url

import (
	"fmt"

	"github.com/dairaga/js/v2"
)

type URL js.Value

func (u URL) JSValue() js.Value {
	return js.Value(u)
}

// -----------------------------------------------------------------------------

func (u URL) Hash() string {
	return u.JSValue().Get("hash").String()
}

// -----------------------------------------------------------------------------

func (u URL) SetHash(val string) {
	u.JSValue().Set("hash", val)
}

// -----------------------------------------------------------------------------

func (u URL) Host() string {
	return u.JSValue().Get("host").String()
}

// -----------------------------------------------------------------------------

func (u URL) SetHost(val string) {
	u.JSValue().Set("host", val)
}

// -----------------------------------------------------------------------------

// Hostname returns hostname property.
func (u URL) Hostname() string {
	return u.JSValue().Get("hostname").String()
}

// -----------------------------------------------------------------------------

// SetHostname sets hostname property.
func (u URL) SetHostname(val string) {
	u.JSValue().Set("hostname", val)
}

// -----------------------------------------------------------------------------

// Href returns href property.
func (u URL) Href() string {
	return u.JSValue().Get("href").String()
}

// -----------------------------------------------------------------------------

// SetHref sets href property.
func (u URL) SetHref(val string) {
	u.JSValue().Set("href", val)
}

// -----------------------------------------------------------------------------

// Origin returns origin property.
func (u URL) Origin() string {
	return u.JSValue().Get("origin").String()
}

// -----------------------------------------------------------------------------

// Password returns password property.
func (u URL) Password() string {
	return u.JSValue().Get("password").String()
}

// -----------------------------------------------------------------------------

// SetPassword sets password property.
func (u URL) SetPassword(val string) {
	u.JSValue().Set("password", val)
}

// -----------------------------------------------------------------------------

// Pathname returns pathname property.
func (u URL) Pathname() string {
	return u.JSValue().Get("pathname").String()
}

// -----------------------------------------------------------------------------

func (u URL) QueryPath() string {
	return u.Pathname()
}

// -----------------------------------------------------------------------------

// SetPathname sets pathname property.
func (u URL) SetPathname(val string) {
	u.JSValue().Set("pathname", val)
}

// -----------------------------------------------------------------------------

func (u URL) SetQueryPath(path string) {
	u.SetPathname(path)
}

// -----------------------------------------------------------------------------

// Port returns port property.
func (u URL) Port() string {
	return u.JSValue().Get("port").String()
}

// -----------------------------------------------------------------------------

// SetPort sets port property.
func (u URL) SetPort(port int) {
	u.JSValue().Set("port", fmt.Sprintf("%d", port))
}

// -----------------------------------------------------------------------------

// Protocol returns protocol property and remove ":" at end.
func (u URL) Protocol() string {
	return u.JSValue().Get("protocol").String()
}

// -----------------------------------------------------------------------------

// SetProtocol sets protocol property. Append ":" automatically if protocol is not end with ":".
func (u URL) SetProtocol(protocol string) {
	if protocol[len(protocol)-1] != ':' {
		protocol += ":"
	}

	u.JSValue().Set("protocol", protocol)
}

// -----------------------------------------------------------------------------

// Search returns search property.
func (u URL) Search() string {
	return u.JSValue().Get("search").String()
}

// -----------------------------------------------------------------------------

// Querystring returns query string in the url.
func (u URL) QueryString() string {
	return u.Search()
}

// SetSearch sets search property.
func (u URL) SetSearch(search string) {
	u.JSValue().Set("search", search)
}

// -----------------------------------------------------------------------------

// SetQuerystring sets query string in the url.
func (u URL) SetQueryString(query string) {
	u.SetSearch(query)
}

// -----------------------------------------------------------------------------

// Username returns username property.
func (u URL) Username() string {
	return u.JSValue().Get("username").String()
}

// -----------------------------------------------------------------------------

// SetUsername sets username property.
func (u URL) SetUsername(username string) {
	u.JSValue().Set("username", username)
}

// -----------------------------------------------------------------------------

// Params returns parameters in url.
func (u URL) Params() Params {
	return Params(u.JSValue().Get("searchParams"))
}

// -----------------------------------------------------------------------------

func (u URL) String() string {
	return u.JSValue().Call("toString").String()
}

// -----------------------------------------------------------------------------

func (u URL) JSON() string {
	return u.JSValue().Call("toJSON").String()
}

// -----------------------------------------------------------------------------

func New(url string) URL {
	return URL(constructor.New(url))
}

// TODO: createObjectURL, revokeObjectURL()
