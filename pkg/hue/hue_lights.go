package hue

import (
	"encoding/json"
)

// Light represents a light object returned by the Hue Bridge.
type Light struct {
	ID int
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
			ColorGamut     [][]float64
			CT             struct {
				Min int
				Max int
			}
		}
		Streaming struct {
			Renderer bool
			Proxy    bool
		}
	}

	Config struct {
		Archetype string
		Function  string
		Direction string
	}

	State struct {
		On        bool
		Bri       int
		Hue       int
		Sat       int
		Effect    string
		XY        []float64
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
func (h *Hue) GetLights() (map[int]*Light, error) {
	respBytes, err := h.API.GetLights()
	if err != nil {
		return nil, err
	}

	var obj map[int]*Light
	if err := json.Unmarshal(respBytes, &obj); err != nil {
		return nil, err
	}

	for ID, light := range obj {
		light.ID = ID
	}

	return obj, nil
}

// GetLight gets light attributes and state.
func (h *Hue) GetLight(ID int) (*Light, error) {
	respBytes, err := h.API.GetLight(ID)
	if err != nil {
		return nil, err
	}

	var obj *Light
	if err := json.Unmarshal(respBytes, obj); err != nil {
		return nil, err
	}

	obj.ID = ID

	return obj, nil
}
