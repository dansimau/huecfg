package hue

import (
	"fmt"
)

// CreateSchedule creates a new schedule
func (api *API) CreateSchedule(data interface{}) ([]byte, error) {
	return api.post(fmt.Sprintf("/api/%s/schedules", api.username()), data)
}

// DeleteSchedule deletes a schedule from the bridge.
func (api *API) DeleteSchedule(id string) ([]byte, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	return api.delete(fmt.Sprintf("/api/%s/schedules/%s", api.username(), id))
}

// GetSchedules gets a list of all schedules that have been added to the
// bridge.
func (api *API) GetSchedules() ([]byte, error) {
	return api.get(fmt.Sprintf("/api/%s/schedules", api.username()))
}

// GetSchedule gets all attributes for a schedule.
func (api *API) GetSchedule(id string) ([]byte, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	return api.get(fmt.Sprintf("/api/%s/schedules/%s", api.username(), id))
}
