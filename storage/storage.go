package storage

import "github.com/dairaga/js"

// Storage represents javascript Storage interface.
type Storage struct {
	ref js.Value
}

// Length returns an integer representing the number of data items stored in the Storage object.
func (s *Storage) Length() int {
	return s.ref.Get("length").Int()
}

// Key returns nth key in the storage.
func (s *Storage) Key(idx int) string {
	return s.ref.Call("key", idx).String()
}

// GetItem returns the key's value.
func (s *Storage) GetItem(key string) (string, bool) {
	val := s.ref.Call("getItem", key)
	if val.Truthy() {
		return val.String(), true
	}

	return "", false
}

// SetItem set key/value into storage.
func (s *Storage) SetItem(key, val string) {
	s.ref.Call("setItem", key, val)
}

// RemoveItem remove key from storage.
func (s *Storage) RemoveItem(key string) {
	s.ref.Call("removeItem", key)
}

// Clear clears all keys in storage.
func (s *Storage) Clear() {
	s.ref.Call("clear")
}

// ----------------------------------------------------------------------------

// Local returns local storage.
func Local() *Storage {
	return &Storage{
		ref: js.Window().Get("localStorage"),
	}
}

// Session returns session storage.
func Session() *Storage {
	return &Storage{
		ref: js.Window().Get("sessionStorage"),
	}
}
