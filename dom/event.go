package dom

// Event ...
type Event struct {
	Value
}

// EventOf ...
func EventOf(x interface{}) Event {
	return Event{ValueOf(x)}
}

// PreventDefault ...
func (e Event) PreventDefault() {
	e.call("preventDefault")
}

// Type ...
func (e Event) Type() string {
	return e.Get("type").String()
}

// Target ...
func (e Event) Target() Element {
	return ElementOf(e.Get("target"))
}
