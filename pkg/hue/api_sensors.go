package hue

import (
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

	resp, err := api.httpGet(fmt.Sprintf("/api/%s/scenes/%s", api.username(), id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
