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

	successMsg, err := statusMsg.ToSuccess()
	if err != nil {
		return Success{}, err
	}

	errorMsg, err := statusMsg.ToError()
	if err != nil {
		return Success{}, err
	}

	return *successMsg, errorMsg
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
