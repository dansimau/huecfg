package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// CreateSchedule creates a new schedule
func (api *API) CreateSchedule(data interface{}) ([]byte, error) {
	postJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpPost(fmt.Sprintf("/api/%s/schedules", api.username()), bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// DeleteSchedule deletes a schedule from the bridge.
func (api *API) DeleteSchedule(id string) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	resp, err := api.httpDelete(fmt.Sprintf("/api/%s/schedules/%s", api.username(), id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// GetSchedules gets a list of all schedules that have been added to the
// bridge.
func (api *API) GetSchedules() ([]byte, error) {
	resp, err := api.httpGet(fmt.Sprintf("/api/%s/schedules", api.username()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// GetSchedule gets all attributes for a schedule.
func (api *API) GetSchedule(id string) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	resp, err := api.httpGet(fmt.Sprintf("/api/%s/schedules/%s", api.username(), id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
