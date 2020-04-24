package hue

import (
	"fmt"
	"io/ioutil"
	"strconv"
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
func (api *API) GetSensor(ID int) ([]byte, error) {
	// if ID == "" {
	// 	return nil, errors.New("ID cannot be empty")
	// }

	strID := strconv.Itoa(ID)

	resp, err := api.httpGet(fmt.Sprintf("/api/%s/scenes/%s", api.username(), strID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
