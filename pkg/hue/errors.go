package hue

// Error represents an error from the Hue Bridge API
type Error struct {
	Address     string
	Description string
	Type        int
}

// Error is the description of the error return from the Hue Bridge API.
func (e *Error) Error() string {
	return e.Description
}
