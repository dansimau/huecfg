package hue

import (
	"fmt"
)

// CreateGroup creates a new group
func (api *API) CreateGroup(data interface{}) ([]byte, error) {
	return api.post(fmt.Sprintf("/api/%s/groups", api.username()), data)
}

// DeleteGroup deletes a group from the bridge.
func (api *API) DeleteGroup(id string) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	return api.delete(fmt.Sprintf("/api/%s/groups/%s", api.username(), id))
}

// GetGroups gets a list of all groups that have been added to the bridge. A
// group is a list of lights that can be created, modified and deleted by a
// user.
func (api *API) GetGroups() ([]byte, error) {
	return api.get(fmt.Sprintf("/api/%s/groups", api.username()))
}

// GetGroup gets the group attributes, e.g. name, light membership and last
// command for a given group.
func (api *API) GetGroup(id string) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	return api.get(fmt.Sprintf("/api/%s/groups/%s", api.username(), id))
}

func (api *API) SetGroupAttributes(id string, attrs interface{}) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	return api.put(fmt.Sprintf("/api/%s/groups/%s", api.username(), id), attrs)
}

func (api *API) SetGroupState(id string, state interface{}) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	return api.put(fmt.Sprintf("/api/%s/groups/%s/state", api.username(), id), state)
}
