package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// GetSensors gets a list of all sensors that have been added to the bridge.
func (api *API) GetSensors() ([]byte, error) {
	resp, err := api.httpGet(fmt.Sprintf("/api/%s/sensors", api.username()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// GetSensor gets the sensor from the bridge with the given id.
func (api *API) GetSensor(id string) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	resp, err := api.httpGet(fmt.Sprintf("/api/%s/sensors/%s", api.username(), id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// SetSensorAttributes is used to rename sensors. A sensor can have its name
// changed when in any state, including when it is unreachable or off.
func (api *API) SetSensorAttributes(id string, attrs interface{}) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	postJSON, err := json.Marshal(attrs)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpPut(fmt.Sprintf("/api/%s/sensors/%s", api.username(), id), bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// SetSensorConfig changes sensor configuration parameters. The allowed
// configuration parameters depend on the sensor type.
func (api *API) SetSensorConfig(id string, state interface{}) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	postJSON, err := json.Marshal(state)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpPut(fmt.Sprintf("/api/%s/sensors/%s/config", api.username(), id), bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// SetSensorState sets the state of a CLIP sensor.
func (api *API) SetSensorState(id string, state interface{}) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	postJSON, err := json.Marshal(state)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpPut(fmt.Sprintf("/api/%s/sensors/%s/state", api.username(), id), bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
