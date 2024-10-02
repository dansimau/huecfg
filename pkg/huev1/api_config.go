package huev1

import (
	"fmt"
)

// CreateUser creates a new user. The link button on the bridge must be pressed and this command executed within 30
// seconds. Once a new user has been created, the user key is added to a 'whitelist', allowing access to API commands
// that require a whitelisted user. At present, all other API commands require a whitelisted user.
func (api *API) CreateUser(deviceType string, generateClientKey bool) ([]byte, error) {
	params := struct {
		DeviceType        string `json:"devicetype"`
		GenerateClientKey *bool  `json:"generateclientkey,omitempty"`
	}{
		DeviceType: deviceType,
	}

	// An oddity of the hue bridge API: in testing, it accepted
	// generateclientkey: true but generateclientkey: false returned an error.
	// The field is marked as optional so it can be omitted.
	if generateClientKey {
		params.GenerateClientKey = &generateClientKey
	}

	return api.post("/api", &params)
}

// GetConfig returns list of all configuration elements in the bridge. Note all times are stored in UTC.
func (api *API) GetConfig() ([]byte, error) {
	return api.get(fmt.Sprintf("/api/%s/config", api.username()))
}

// GetFullState is used to fetch the entire datastore from the device,
// including settings and state information for lights, groups, schedules and
// configuration. It should only be used sparingly as it is resource intensive
// for the bridge, but is supplied e.g. for synchronization purposes.
func (api *API) GetFullState() ([]byte, error) {
	return api.get(fmt.Sprintf("/api/%s", api.username()))
}
