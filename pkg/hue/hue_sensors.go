package hue

import (
	"encoding/json"
)

type Sensor struct {
	ID string

	ManufacturerName string
	ModelID          string
	Name             string
	Recycle          bool
	SWVersion        string
	Type             string // TODO: string enum?
	UniqueID         string

	Config struct {
		Battery   int
		On        bool
		Reachable bool
	}

	// TODO: map out all possible state fields
	State map[string]interface{}
}

// GetSensors gets a list of all sensors that have been added to the bridge.
func (h *Hue) GetSensors() ([]Sensor, error) {
	respBytes, err := h.API.GetSensors()
	if err != nil {
		return nil, err
	}

	var objs map[string]Sensor
	if err := json.Unmarshal(respBytes, &objs); err != nil {
		return nil, err
	}

	var res = []Sensor{}
	for id, obj := range objs {
		obj.ID = id
		res = append(res, obj)
	}

	return res, nil
}

// GetSensor gets the sensor from the bridge with the given ID
func (h *Hue) GetSensor(id string) (Sensor, error) {
	respBytes, err := h.API.GetSensor(id)
	if err != nil {
		return Sensor{}, err
	}

	var obj Sensor
	if err := json.Unmarshal(respBytes, &obj); err != nil {
		return Sensor{}, err
	}

	obj.ID = id

	return obj, nil
}
