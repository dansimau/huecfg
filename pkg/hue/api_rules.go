package hue

import (
	"fmt"
)

// CreateRule creates a new rule
func (api *API) CreateRule(data interface{}) ([]byte, error) {
	return api.post(fmt.Sprintf("/api/%s/rules", api.username()), data)
}

// DeleteRule deletes a rule from the bridge.
func (api *API) DeleteRule(id string) ([]byte, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	return api.delete(fmt.Sprintf("/api/%s/rules/%s", api.username(), id))
}

// GetRules gets a list of all rules that are in the bridge.
func (api *API) GetRules() ([]byte, error) {
	return api.get(fmt.Sprintf("/api/%s/rules", api.username()))
}

// GetRule returns a rule object with id matching <id> or an error if <id> is
// not available.
func (api *API) GetRule(id string) ([]byte, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	return api.get(fmt.Sprintf("/api/%s/rules/%s", api.username(), id))
}
