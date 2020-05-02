package hue

import (
	"fmt"
	"io/ioutil"
)

// GetLights gets a list of all lights that have been discovered by the bridge.
func (api *API) GetLights() ([]byte, error) {
	resp, err := api.httpGet(fmt.Sprintf("/api/%s/lights", api.username()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// GetLight light attributes and state.
func (api *API) GetLight(id string) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	resp, err := api.httpGet(fmt.Sprintf("/api/%s/lights/%s", api.username(), id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
