package hue

import (
	"errors"
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
func (api *API) GetLight(ID string) ([]byte, error) {
	if ID == "" {
		return nil, errors.New("ID cannot be empty")
	}

	resp, err := api.httpGet(fmt.Sprintf("/api/%s/lights/%s", api.username(), ID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
