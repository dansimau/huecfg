package hue

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// LightsAPI is the API for lights on the Hue Bridge.
type LightsAPI struct {
	hue *Hue
}

// Light represents a light object returned by the Hue Bridge.
type Light struct {
	*responseData

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

// Lights represents a map of lights returned by the Hue Bridge.
type Lights struct {
	*responseData

	m map[string]*Light
}

func NewLights(m map[string]*Light) *Lights {
	return &Lights{
		m: m,
	}
}

// All returns a map of lights, indexed by string ID.
func (l *Lights) All() map[string]*Light {
	return l.m
}

// ByID returns a light from the map, by string ID. If the light doesn't exist,
// this method returns nil.
func (l *Lights) ByID(ID string) *Light {
	item, ok := l.m[ID]
	if !ok {
		return nil
	}
	return item
}

// GetAll gets a list of all lights that have been discovered by the bridge.
func (h *LightsAPI) GetAll() (*Lights, error) {
	resp, err := h.hue.httpGet(fmt.Sprintf("%s/api/%s/lights", h.hue.host, h.hue.requireUsername()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var obj map[string]*Light
	if err := json.Unmarshal(content, &obj); err != nil {
		return nil, err
	}

	return &Lights{
		m:            obj,
		responseData: &responseData{content},
	}, nil
}

// Get light attributes and state.
func (h *LightsAPI) Get(ID string) (*Light, error) {
	if ID == "" {
		return nil, errors.New("ID cannot be empty")
	}

	resp, err := h.hue.httpGet(fmt.Sprintf("%s/api/%s/lights/%s", h.hue.host, h.hue.requireUsername(), ID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var obj *Light
	if err := json.Unmarshal(content, obj); err != nil {
		return nil, err
	}
	obj.responseData = &responseData{content}

	return obj, nil
}
