package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// CreateGroup creates a new group
func (api *API) CreateGroup(data interface{}) ([]byte, error) {
	postJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpPost(fmt.Sprintf("/api/%s/groups", api.username()), bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// DeleteGroup deletes a group from the bridge.
func (api *API) DeleteGroup(id string) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	resp, err := api.httpDelete(fmt.Sprintf("/api/%s/groups/%s", api.username(), id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// GetGroups gets a list of all groups that have been added to the bridge. A
// group is a list of lights that can be created, modified and deleted by a
// user.
func (api *API) GetGroups() ([]byte, error) {
	resp, err := api.httpGet(fmt.Sprintf("/api/%s/groups", api.username()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// GetGroup gets the group attributes, e.g. name, light membership and last
// command for a given group.
func (api *API) GetGroup(id string) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	resp, err := api.httpGet(fmt.Sprintf("/api/%s/groups/%s", api.username(), id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (api *API) SetGroupAttributes(id string, attrs interface{}) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	postJSON, err := json.Marshal(attrs)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpPut(fmt.Sprintf("/api/%s/groups/%s", api.username(), id), bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (api *API) SetGroupState(id string, state interface{}) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	postJSON, err := json.Marshal(state)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpPut(fmt.Sprintf("/api/%s/groups/%s/state", api.username(), id), bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
