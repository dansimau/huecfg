package huev1

import (
	"encoding/json"
)

type Schedule struct {
	ID string

	// TODO: add descriptions to all the vars from: https://developers.meethuev1.com/develop/hue-api/3-schedules-api/
	Name        string
	Description string
	Command     struct {
		Address string
		Method  string
		Body    interface{}
	}
	Status     string // TODO: turn into enabled/disable enum
	AutoDelete bool
	LocalTime  string // TODO: Make this into a proper Go type
	Recycle    bool
}

func (h *Hue) GetSchedules() ([]Schedule, error) {
	respBytes, err := h.API.GetSchedules()
	if err != nil {
		return nil, err
	}

	if hueErr := parseAsHueError(respBytes); hueErr != nil {
		return nil, hueErr
	}

	var objs map[string]Schedule
	if err := json.Unmarshal(respBytes, &objs); err != nil {
		return nil, err
	}

	res := []Schedule{}
	for id, obj := range objs {
		obj.ID = id
		res = append(res, obj)
	}

	return res, nil
}

func (h *Hue) GetSchedule(id string) (Schedule, error) {
	respBytes, err := h.API.GetSchedule(id)
	if err != nil {
		return Schedule{}, err
	}

	if hueErr := parseAsHueError(respBytes); hueErr != nil {
		return Schedule{}, hueErr
	}

	var obj Schedule
	if err := json.Unmarshal(respBytes, &obj); err != nil {
		return Schedule{}, err
	}

	obj.ID = id

	return obj, nil
}
