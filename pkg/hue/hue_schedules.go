package hue

import (
	"encoding/json"
)

type Schedule struct {
	ID int

	// TODO: add descriptions to all the vars from: https://developers.meethue.com/develop/hue-api/3-schedules-api/
	Name        string
	Description string
	Command     struct {
		Address string
		Method  string
		Body    string // TODO: check format here with api command; is it in fact a JSON string?
	}
	Status     string // TODO: turn into enabled/disable enum
	AutoDelete bool
	LocalTime  AbsoluteTime
	Recycle    bool
}

func (h *Hue) GetSchedules() ([]Schedule, error) {
	respBytes, err := h.API.GetSchedules()
	if err != nil {
		return nil, err
	}

	var objs map[int]Schedule
	if err := json.Unmarshal(respBytes, &objs); err != nil {
		return nil, err
	}

	var res = []Schedule{}
	for ID, obj := range objs {
		obj.ID = ID
		res = append(res, obj)
	}

	return res, nil
}

func (h *Hue) GetSchedule(ID int) (Schedule, error) {
	respBytes, err := h.API.GetSchedule(ID)
	if err != nil {
		return Schedule{}, err
	}

	var obj Schedule
	if err := json.Unmarshal(respBytes, &obj); err != nil {
		return Schedule{}, err
	}

	obj.ID = ID

	return obj, nil
}
