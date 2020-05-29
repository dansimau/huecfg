package hue

import (
	"encoding/json"
)

type Group struct {
	ID string

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

	if hueErr := parseAsHueError(respBytes); hueErr != nil {
		return nil, hueErr
	}

	var objs map[string]Group
	if err := json.Unmarshal(respBytes, &objs); err != nil {
		return nil, err
	}

	var res = []Group{}
	for id, obj := range objs {
		obj.ID = id
		res = append(res, obj)
	}

	return res, nil
}

func (h *Hue) GetGroup(id string) (Group, error) {
	respBytes, err := h.API.GetGroup(id)
	if err != nil {
		return Group{}, err
	}

	if hueErr := parseAsHueError(respBytes); hueErr != nil {
		return Group{}, hueErr
	}

	var obj Group
	if err := json.Unmarshal(respBytes, &obj); err != nil {
		return Group{}, err
	}

	obj.ID = id

	return obj, nil
}
