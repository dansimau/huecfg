package huev1

import (
	"encoding/json"
)

type ResourceUsage struct {
	Available int
	Total     int
}

type Capabilities struct {
	Groups        GroupsCapabilities
	Lights        LightsCapabilities
	ResourceLinks ResourceLinksCapabilities
	Rules         RulesCapabilities
	Scenes        ScenesCapabilities
	Schedules     SchedulesCapabilities
	Sensors       SensorsCapabilities
	Streaming     StreamingCapabilities
	TimeZones     TimeZoneCapabilities
}

type GroupsCapabilities struct {
	ResourceUsage
}

type LightsCapabilities struct {
	ResourceUsage
}

type ResourceLinksCapabilities struct {
	ResourceUsage
}

type RulesCapabilities struct {
	ResourceUsage
	Actions    ResourceUsage
	Conditions ResourceUsage
}

type ScenesCapabilities struct {
	ResourceUsage
	LightStates ResourceUsage
}

type SchedulesCapabilities struct {
	ResourceUsage
}

type SensorsCapabilities struct {
	ResourceUsage
	Clip ResourceUsage
	Zgb  ResourceUsage // TODO: verify letter case of "zgb"
	Zll  ResourceUsage // TODO: verify letter case of "zll"
}

type StreamingCapabilities struct {
	ResourceUsage
	Channels int
}

type TimeZoneCapabilities struct {
	Values []string
}

func (h *Hue) GetCapabilities() (Capabilities, error) {
	respBytes, err := h.API.GetCapabilities()
	if err != nil {
		return Capabilities{}, err
	}

	if hueErr := parseAsHueError(respBytes); hueErr != nil {
		return Capabilities{}, hueErr
	}

	var obj Capabilities
	if err := json.Unmarshal(respBytes, &obj); err != nil {
		return Capabilities{}, err
	}

	return obj, nil
}
