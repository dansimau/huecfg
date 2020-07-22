package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// CreateResourceLink creates a new resourcelink
func (api *API) CreateResourceLink(data interface{}) ([]byte, error) {
	postJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpPost(fmt.Sprintf("/api/%s/resourcelinks", api.username()), bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// DeleteResourceLink deletes a resourcelink from the bridge.
func (api *API) DeleteResourceLink(id string) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	resp, err := api.httpDelete(fmt.Sprintf("/api/%s/resourcelinks/%s", api.username(), id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// GetResourceLinks gets a list of all resourcelinks that have been added to the
// bridge.
func (api *API) GetResourceLinks() ([]byte, error) {
	resp, err := api.httpGet(fmt.Sprintf("/api/%s/resourcelinks", api.username()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// GetResourceLink gets all attributes for a resourcelink.
func (api *API) GetResourceLink(id string) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	resp, err := api.httpGet(fmt.Sprintf("/api/%s/resourcelinks/%s", api.username(), id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
