package js

// ----------------------------------------------------------------------------

// Error ...
type Error struct {
	ref Value
}

// ErrorOf ...
func ErrorOf(x Value) Error {
	return Error{ref: x}
}

// JSValue ...
func (err Error) JSValue() Value {
	return err.ref
}

func (err Error) Error() string {
	return err.ref.String()
}
