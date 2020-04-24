package hue

import (
	"fmt"
	"io/ioutil"
	"strconv"
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
func (api *API) GetSchedule(ID int) ([]byte, error) {
	// if ID == "" {
	// 	return nil, errors.New("ID cannot be empty")
	// }

	strID := strconv.Itoa(ID)

	resp, err := api.httpGet(fmt.Sprintf("/api/%s/schedules/%s", api.username(), strID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
