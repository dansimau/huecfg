package hue

import (
	"encoding/json"
)

type Group struct {
	ID int

	Name   string
	Lights []string // TODO: convert this to int, or convert IDs back to string?
	Type   string

	Action struct {
		On        bool
		Bri       int
		Hue       int
		Sat       int
		Effect    string    // TODO: can be "none"
		XY        []float64 // TODO: make this proper pair type
		CT        int
		Alert     string
		ColorMode string
	}
}

func (h *Hue) GetGroups() ([]Group, error) {
	respBytes, err := h.API.GetGroups()
	if err != nil {
		return nil, err
	}

	var objs map[int]Group
	if err := json.Unmarshal(respBytes, &objs); err != nil {
		return nil, err
	}

	var res = []Group{}
	for ID, obj := range objs {
		obj.ID = ID
		res = append(res, obj)
	}

	return res, nil
}

func (h *Hue) GetGroup(ID int) (Group, error) {
	respBytes, err := h.API.GetGroup(ID)
	if err != nil {
		return Group{}, err
	}

	var obj Group
	if err := json.Unmarshal(respBytes, &obj); err != nil {
		return Group{}, err
	}

	obj.ID = ID

	return obj, nil
}
