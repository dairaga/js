package bs

// Style represents Bootstrap component style like primary, secondary and etc.
type Style = string

// pre-defined style.
const (
	Primary   Style = "primary"
	Secondary Style = "secondary"
	Success   Style = "success"
	Danger    Style = "danger"
	Warning   Style = "warning"
	Info      Style = "info"
	Light     Style = "light"
	Dark      Style = "dark"

	// for button
	link Style = "link"
)
