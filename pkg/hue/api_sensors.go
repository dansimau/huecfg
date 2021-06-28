package hue

import (
	"fmt"
)

// CreateSensor creates a new sensor
func (api *API) CreateSensor(data interface{}) ([]byte, error) {
	return api.post(fmt.Sprintf("/api/%s/sensors", api.username()), data)
}

// DeleteSensor deletes a sensor from the bridge.
func (api *API) DeleteSensor(id string) ([]byte, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	return api.delete(fmt.Sprintf("/api/%s/sensors/%s", api.username(), id))
}

// GetSensors gets a list of all sensors that have been added to the bridge.
func (api *API) GetSensors() ([]byte, error) {
	return api.get(fmt.Sprintf("/api/%s/sensors", api.username()))
}

// GetSensor gets the sensor from the bridge with the given id.
func (api *API) GetSensor(id string) ([]byte, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	return api.get(fmt.Sprintf("/api/%s/sensors/%s", api.username(), id))
}

// SetSensorAttributes is used to rename sensors. A sensor can have its name
// changed when in any state, including when it is unreachable or off.
func (api *API) SetSensorAttributes(id string, attrs interface{}) ([]byte, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	return api.put(fmt.Sprintf("/api/%s/sensors/%s", api.username(), id), attrs)
}

// SetSensorConfig changes sensor configuration parameters. The allowed
// configuration parameters depend on the sensor type.
func (api *API) SetSensorConfig(id string, config interface{}) ([]byte, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	return api.put(fmt.Sprintf("/api/%s/sensors/%s/config", api.username(), id), config)
}

// SetSensorState sets the state of a CLIP sensor.
func (api *API) SetSensorState(id string, state interface{}) ([]byte, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	return api.put(fmt.Sprintf("/api/%s/sensors/%s/state", api.username(), id), state)
}
