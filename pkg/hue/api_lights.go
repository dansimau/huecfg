package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// DeleteLight deletes a light from the bridge.
func (api *API) DeleteLight(id string) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	resp, err := api.httpDelete(fmt.Sprintf("/api/%s/lights/%s", api.username(), id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

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

// SearchForNewLights starts searching for new lights.
//
// The bridge will open the network for 40s. The overall search might take
// longer since the configuration of (multiple) new devices can take longer. If
// many devices are found the command will have to be issued a second time after
// discovery time has elapsed. If the command is received again during search
// the search will continue for at least an additional 40s.
//
// When the search has finished, new lights will be available using the get new
// lights command. In addition, the new lights will now be available by calling
// get all lights or by calling get group attributes on group 0. Group 0 is a
// special group that cannot be deleted and will always contain all lights known
// by the bridge.
func (api *API) SearchForNewLights(deviceIds ...string) ([]byte, error) {
	params := struct {
		DeviceIDs []string `json:"deviceid,omitempty"`
	}{deviceIds}

	postJSON, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpPost(fmt.Sprintf("/api/%s/lights", api.username()), bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// SetLightAttributes is used to rename lights. A light can have its name
// changed when in any state, including when it is unreachable or off.
func (api *API) SetLightAttributes(id string, attrs interface{}) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	postJSON, err := json.Marshal(attrs)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpPut(fmt.Sprintf("/api/%s/lights/%s", api.username(), id), bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// SetLightState allows you to turn the light on and off, modify the hue and
// effects.
func (api *API) SetLightState(id string, state interface{}) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	postJSON, err := json.Marshal(state)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpPut(fmt.Sprintf("/api/%s/lights/%s/state", api.username(), id), bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
