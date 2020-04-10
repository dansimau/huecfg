package hue

// Hue represents a Hue Bridge
type Hue struct {
	API *API
}

// NewConn creates a connection to a Hue Bridge.
func NewConn(host string, username string) *Hue {
	return &Hue{
		API: &API{
			Host:     host,
			Username: username,
		},
	}
}
