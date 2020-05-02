package hue

import (
	"encoding/json"
)

type CapabilitiesResourceUsage struct {
	Available int
	Total     int
}

type Capabilities struct {
	Groups        CapabilitiesResourceUsage
	Lights        CapabilitiesResourceUsage
	ResourceLinks CapabilitiesResourceUsage
	Rules         struct {
		CapabilitiesResourceUsage
		Actions    CapabilitiesResourceUsage
		Conditions CapabilitiesResourceUsage
	}
	Scenes struct {
		CapabilitiesResourceUsage
		LightStates CapabilitiesResourceUsage
	}
	Schedules CapabilitiesResourceUsage
	Sensors   struct {
		CapabilitiesResourceUsage
		Clip CapabilitiesResourceUsage
		Zgb  CapabilitiesResourceUsage // TODO: verify letter case of "zgb"
		Zll  CapabilitiesResourceUsage // TODO: verify letter case of "zll"
	}
	Streaming struct {
		CapabilitiesResourceUsage
		Channels int
	}
	TimeZones struct {
		Values []string
	}
}

func (h *Hue) GetCapabilities() (Capabilities, error) {
	respBytes, err := h.API.GetCapabilities()
	if err != nil {
		return Capabilities{}, err
	}

	var obj Capabilities
	if err := json.Unmarshal(respBytes, &obj); err != nil {
		return Capabilities{}, err
	}

	return obj, nil
}
