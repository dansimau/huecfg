package hue

import (
	"fmt"
	"io/ioutil"
)

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
