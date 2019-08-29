// +build js,wasm

/*Package url wraps javascript url object and funcions. */
package url

import (
	"fmt"

	"github.com/dairaga/js"
)

var (
	_url = js.Value{}
)

func init() {
	if x := js.Window().Get("URL"); x.Truthy() {
		_url = x
	} else if x := js.Window().Get("webkitURL"); x.Truthy() {
		_url = x
	} else {
		panic("Window.URL not supported")
	}
}

//-----------------------------------------------------------------------------

// Params represents javascript URLSearchParams. https://developer.mozilla.org/en-US/docs/Web/API/URLSearchParams
type Params struct {
	ref js.Value
}

// JSValue ...
func (p *Params) JSValue() js.Value {
	return p.ref
}

func (p *Params) String() string {
	return p.ref.Call("toString").String()
}

// Has returns boolean indicates whether or not params has the key.
func (p *Params) Has(name string) bool {
	return p.ref.Call("has", name).Bool()
}

// Get returns the first value of the key.
func (p *Params) Get(name string) (string, bool) {
	if x := p.ref.Call("get", name); x.Truthy() {
		return x.String(), true
	}

	return "", false
}

// GetAll returns all values of the key.
func (p *Params) GetAll(name string) []string {
	if x := p.ref.Call("getAll", name); x.Truthy() {
		if size := x.Length(); size > 0 {
			result := make([]string, size)
			for i := 0; i < size; i++ {
				result[i] = x.Index(i).String()
			}
		}
	}

	return nil
}

// Set sets value for the key.
func (p *Params) Set(name, value string) *Params {
	p.ref.Call("set", name, value)
	return p
}

// Append appends a value to the key.
func (p *Params) Append(name, value string) *Params {
	p.ref.Call("append", name, value)
	return p
}

// Delete removes the key from parameters.
func (p *Params) Delete(name string) *Params {
	p.ref.Call("delete", name)
	return p
}

// Foreach applies fn on each key/value pair.
func (p *Params) Foreach(fn func(key, value string)) {
	cb := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		fn(args[1].String(), args[0].String())
		return nil
	})

	p.ref.Call("forEach", cb)
	cb.Release()
}

//-----------------------------------------------------------------------------

// URL represents jaavascript url. https://developer.mozilla.org/en-US/docs/Web/API/URL
type URL struct {
	ref js.Value
}

// New returns a URL object.
func New(link string) *URL {
	return &URL{ref: _url.New(link)}
}

// JSValue ...
func (u *URL) JSValue() js.Value {
	return u.ref
}

// Hash returns hash property.
func (u *URL) Hash() string {
	return u.ref.Get("hash").String()
}

// SetHash set hash property.
func (u *URL) SetHash(hash string) *URL {
	u.ref.Set("hash", hash)
	return u
}

// Host returns host property.
func (u *URL) Host() string {
	return u.ref.Get("host").String()
}

// SetHost sets host property.
func (u *URL) SetHost(host string) *URL {
	u.ref.Set("host", host)
	return u
}

// Hostname returns hostname property.
func (u *URL) Hostname() string {
	return u.ref.Get("hostname").String()
}

// SetHostname sets hostname property.
func (u *URL) SetHostname(hostname string) *URL {
	u.ref.Set("hostname", hostname)
	return u
}

// Href returns href property.
func (u *URL) Href() string {
	return u.ref.Get("href").String()
}

// SetHref sets href property.
func (u *URL) SetHref(href string) *URL {
	u.ref.Set("href", href)
	return u
}

// Origin returns origin property.
func (u *URL) Origin() string {
	return u.ref.Get("origin").String()
}

// Password returns password property.
func (u *URL) Password() string {
	return u.ref.Get("password").String()
}

// SetPassword sets password property.
func (u *URL) SetPassword(password string) *URL {
	u.ref.Set("password", password)
	return u
}

// Pathname returns pathname property.
func (u *URL) Pathname() string {
	return u.ref.Get("pathname").String()
}

// SetPathname sets pathname property.
func (u *URL) SetPathname(pathname string) *URL {
	u.ref.Set("pathname", pathname)
	return u
}

// Port returns port property.
func (u *URL) Port() string {
	return u.ref.Get("port").String()
}

// SetPort sets port property.
func (u *URL) SetPort(port int) *URL {
	u.ref.Set("port", fmt.Sprintf("%d", port))
	return u
}

// Protocol returns protocol property and remove ":" at end.
func (u *URL) Protocol() string {
	if x := u.ref.Get("protocol"); x.Truthy() {
		if tmp := x.String(); tmp != "" {
			return tmp[0 : len(tmp)-1]
		}
	}

	return ""
}

// SetProtocol sets protocol property. Append ":" automatically if protocol is not end with ":".
func (u *URL) SetProtocol(protocol string) *URL {
	if protocol != "" {
		if protocol[len(protocol)-1] == ':' {
			u.ref.Set("protocol", protocol)
		} else {
			u.ref.Set("protocol", protocol+":")
		}
	}
	return u
}

// Search returns search property.
func (u *URL) Search() string {
	return u.Querystring()
}

// Querystring returns query string in the url.
func (u *URL) Querystring() string {
	return u.ref.Get("search").String()
}

// SetSearch sets search property.
func (u *URL) SetSearch(search string) *URL {
	return u.SetQuerystring(search)
}

// SetQuerystring sets query string in the url.
func (u *URL) SetQuerystring(query string) *URL {
	u.ref.Set("search", query)
	return u
}

// Username returns username property.
func (u *URL) Username() string {
	return u.ref.Get("username").String()
}

// SetUsername sets username property.
func (u *URL) SetUsername(username string) *URL {
	u.ref.Set("username", username)
	return u
}

// Params returns parameters in url.
func (u *URL) Params() *Params {
	return &Params{ref: u.ref.Get("searchParams")}
}

func (u *URL) String() string {

	if x := u.ref.Get("toString"); x.Truthy() {
		fmt.Println("xxx")
		return u.ref.Call("toString").String()
	}

	return fmt.Sprintf("%s://%s%s%s%s", u.Protocol(), u.Host(), u.Pathname(), u.Querystring(), u.Hash())
}

// ----------------------------------------------------------------------------

// Create returns a URL representing the obj.
func Create(obj interface{}) string {
	return _url.Call("createObjectURL", obj).String()
}

// Revoke releases an existing object URL. https://developer.mozilla.org/en-US/docs/Web/API/URL/revokeObjectURL
func Revoke(link string) {
	_url.Call("revokeObjectURL", link)
}
