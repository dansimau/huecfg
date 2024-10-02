package huev1

import (
	"encoding/json"
)

// Light represents a light object returned by the Hue Bridge.
type Light struct {
	ID               string
	ModelID          string // The hardware model of the light.
	ManufacturerName string
	Name             string // A unique, editable name given to the light.
	ProductName      string
	SWVersion        string
	Type             string // A fixed name describing the type of light e.g. "Extended color light".
	UniqueID         string // Unique id of the device. The MAC address of the device with a unique endpoint id in the form: AA:BB:CC:DD:EE:FF:00:11-XX

	Capabilities LightCapabilities
	Config       LightConfig
	State        LightState
	SWUpdate     LightSWUpdate
}

type LightCapabilities struct {
	Certified bool
	Control   LightControlCapabilities
	Streaming LightStreamingCapabilities
}

type LightColorTemperature struct {
	Min int
	Max int
}

type LightConfig struct { // TODO: verify these types
	Archetype string
	Function  string
	Direction string
}

type LightControlCapabilities struct {
	MinDimLevel    int
	MaxLumen       int
	ColorGamutType string
	ColorGamut     [][]float64 // TODO: proper type
	CT             LightColorTemperature
}

type LightState struct { // TODO: verify all of these field types
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

type LightStreamingCapabilities struct {
	Renderer bool
	Proxy    bool
}

type LightSWUpdate struct {
	State       string
	LastInstall *AbsoluteTime
}

type SetLightAttributeParams struct {
	Name *string `json:"name,omitempty"`
}

type SetLightStateParams struct {
	On             *bool `json:"on,omitempty"`
	Bri            *int  `json:"bri,omitempty"`
	TransitionTime *int  `json:"transitiontime,omitempty"`
}

// DeleteLight deletes a light from the bridge.
func (h *Hue) DeleteLight(id string) (Success, error) {
	respBytes, err := h.API.DeleteLight(id)
	if err != nil {
		return Success{}, err
	}

	var statusMsg Status
	if err := json.Unmarshal(respBytes, &statusMsg); err != nil {
		return Success{}, err
	}

	return *statusMsg.Success, statusMsg.Error
}

// GetLights gets a list of all lights that have been discovered by the bridge.
func (h *Hue) GetLights() ([]Light, error) {
	respBytes, err := h.API.GetLights()
	if err != nil {
		return nil, err
	}

	if hueErr := parseAsHueError(respBytes); hueErr != nil {
		return nil, hueErr
	}

	var objs map[string]Light
	if err := json.Unmarshal(respBytes, &objs); err != nil {
		return nil, err
	}

	res := []Light{}
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

	if hueErr := parseAsHueError(respBytes); hueErr != nil {
		return Light{}, hueErr
	}

	var obj Light
	if err := json.Unmarshal(respBytes, &obj); err != nil {
		return Light{}, err
	}

	obj.ID = id

	return obj, nil
}

// SearchForNewLights gets light attributes and state.
func (h *Hue) SearchForNewLights(deviceIds ...string) (SuccessMessages, error) {
	respBytes, err := h.API.SearchForNewLights(deviceIds...)
	if err != nil {
		return nil, err
	}

	var statusMsgs StatusResponse
	if err := json.Unmarshal(respBytes, &statusMsgs); err != nil {
		return nil, err
	}

	if errs := statusMsgs.Errors(); errs != nil {
		return nil, errs
	}

	return statusMsgs.SuccessMessages(), nil
}

func (h *Hue) SetLightAttributes(id string, attrs SetLightAttributeParams) (StatusResponse, error) {
	respBytes, err := h.API.SetLightAttributes(id, attrs)
	if err != nil {
		return nil, err
	}

	var statusMsgs StatusResponse
	if err := json.Unmarshal(respBytes, &statusMsgs); err != nil {
		return nil, err
	}

	if errs := statusMsgs.Errors(); errs != nil {
		return nil, errs
	}

	return statusMsgs, nil
}

func (h *Hue) SetLightState(id string, state SetLightStateParams) (StatusResponse, error) {
	respBytes, err := h.API.SetLightState(id, state)
	if err != nil {
		return nil, err
	}

	var statusMsgs StatusResponse
	if err := json.Unmarshal(respBytes, &statusMsgs); err != nil {
		return nil, err
	}

	if errs := statusMsgs.Errors(); errs != nil {
		return nil, errs
	}

	return statusMsgs, nil
}
