package hue

import (
	"errors"
	"fmt"
	"io/ioutil"
)

// GetSensors gets a list of all sensors that have been added to the bridge.
func (api *API) GetSensors() ([]byte, error) {
	resp, err := api.httpGet(fmt.Sprintf("/api/%s/scenes", api.username()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// GetSensor gets the sensor from the bridge with the given id.
func (api *API) GetSensor(ID string) ([]byte, error) {
	if ID == "" {
		return nil, errors.New("ID cannot be empty")
	}

	resp, err := api.httpGet(fmt.Sprintf("/api/%s/scenes/%s", api.username(), ID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
