package hue

import (
	"encoding/json"
)

// Light represents a light object returned by the Hue Bridge.
type Light struct {
	ID string
	// The hardware model of the light.
	ModelID          string
	ManufacturerName string
	// A unique, editable name given to the light.
	Name        string
	ProductName string
	SWVersion   string
	// A fixed name describing the type of light e.g. "Extended color light".
	Type string
	// Unique id of the device. The MAC address of the device with a unique endpoint id in the form: AA:BB:CC:DD:EE:FF:00:11-XX
	UniqueID string

	Capabilities struct {
		Certified bool
		Control   struct {
			MinDimLevel    int
			MaxLumen       int
			ColorGamutType string
			ColorGamut     [][]float64 // TODO: proper type
			CT             struct {    // TODO: verify camelcase
				Min int
				Max int
			}
		}
		Streaming struct {
			Renderer bool
			Proxy    bool
		}
	}

	Config struct { // TODO: verify these types
		Archetype string
		Function  string
		Direction string
	}

	State struct { // TODO: verify all of these field types
		On        bool
		Bri       int
		Hue       int
		Sat       int
		Effect    string    // TODO: can be "none"
		XY        []float64 // TODO: make this proper pair type
		CT        int
		Alert     string
		ColorMode string
		Mode      string
		Reachable bool
	}

	SWUpdate struct {
		State       string
		LastInstall *AbsoluteTime
	}
}

// GetLights gets a list of all lights that have been discovered by the bridge.
func (h *Hue) GetLights() ([]Light, error) {
	respBytes, err := h.API.GetLights()
	if err != nil {
		return nil, err
	}

	var objs map[string]Light
	if err := json.Unmarshal(respBytes, &objs); err != nil {
		return nil, err
	}

	var res = []Light{}
	for id, obj := range objs {
		obj.ID = id
		res = append(res, obj)
	}

	return res, nil
}

// GetLight gets light attributes and state.
func (h *Hue) GetLight(id string) (Light, error) {
	respBytes, err := h.API.GetLight(id)
	if err != nil {
		return Light{}, err
	}

	var obj Light
	if err := json.Unmarshal(respBytes, &obj); err != nil {
		return Light{}, err
	}

	obj.ID = id

	return obj, nil
}
