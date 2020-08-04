package hue

import (
	"fmt"
)

// CreateResourceLink creates a new resourcelink
func (api *API) CreateResourceLink(data interface{}) ([]byte, error) {
	return api.post(fmt.Sprintf("/api/%s/resourcelinks", api.username()), data)
}

// DeleteResourceLink deletes a resourcelink from the bridge.
func (api *API) DeleteResourceLink(id string) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	return api.delete(fmt.Sprintf("/api/%s/resourcelinks/%s", api.username(), id))
}

// GetResourceLinks gets a list of all resourcelinks that have been added to the
// bridge.
func (api *API) GetResourceLinks() ([]byte, error) {
	return api.get(fmt.Sprintf("/api/%s/resourcelinks", api.username()))
}

// GetResourceLink gets all attributes for a resourcelink.
func (api *API) GetResourceLink(id string) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	return api.get(fmt.Sprintf("/api/%s/resourcelinks/%s", api.username(), id))
}
