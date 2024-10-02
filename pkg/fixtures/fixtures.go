package fixtures

import "github.com/dansimau/huecfg/pkg/huev1"

var Data = []byte(`
		{
			"capabilities": {
				"groups": {
					"available": 50,
					"total": 64
				},
				"lights": {
					"available": 31,
					"total": 63
				},
				"resourcelinks": {
					"available": 38,
					"total": 64
				},
				"rules": {
					"actions": {
					"available": 544,
					"total": 1000
					},
					"available": 50,
					"conditions": {
					"available": 867,
					"total": 1500
					},
					"total": 250
				},
				"scenes": {
					"available": 74,
					"lightstates": {
					"available": 3004,
					"total": 12600
					},
					"total": 200
				},
				"schedules": {
					"available": 78,
					"total": 100
				},
				"sensors": {
					"available": 181,
					"clip": {
					"available": 181,
					"total": 250
					},
					"total": 250,
					"zgp": {
					"available": 48,
					"total": 64
					},
					"zll": {
					"available": 48,
					"total": 64
					}
				},
				"streaming": {
					"available": 0,
					"channels": 20,
					"total": 1
				},
				"timezones": {
					"values": [
					"CET",
					"CST6CDT",
					"EET",
					"EST",
					"EST5EDT",
					"HST",
					"MET",
					"MST",
					"MST7MDT",
					"PST8PDT",
					"WET"
					]
				}
			},
			"lights": {
				"1": {
					"state": {
						"on": false,
						"bri": 1,
						"hue": 33761,
						"sat": 254,
						"effect": "none",
						"xy": [
							0.3171,
							0.3366
						],
						"ct": 159,
						"alert": "none",
						"colormode": "xy",
						"mode": "homeautomation",
						"reachable": true
					},
					"swupdate": {
						"state": "noupdates",
						"lastinstall": "2018-01-02T19:24:20"
					},
					"type": "Extended color light",
					"name": "Hue color lamp 7",
					"modelid": "LCT007",
					"manufacturername": "Philips",
					"productname": "Hue color lamp",
					"capabilities": {
						"certified": true,
						"control": {
							"mindimlevel": 5000,
							"maxlumen": 600,
							"colorgamuttype": "B",
							"colorgamut": [
								[
									0.675,
									0.322
								],
								[
									0.409,
									0.518
								],
								[
									0.167,
									0.04
								]
							],
							"ct": {
								"min": 153,
								"max": 500
							}
						},
						"streaming": {
							"renderer": true,
							"proxy": false
						}
					},
					"config": {
						"archetype": "sultanbulb",
						"function": "mixed",
						"direction": "omnidirectional"
					},
					"uniqueid": "00:17:88:01:00:bd:c7:b9-0b",
					"swversion": "5.105.0.21169"
				}
			}
		}
	`)

var (
	Capabilities = huev1.Capabilities{
		Groups: huev1.GroupsCapabilities{
			ResourceUsage: huev1.ResourceUsage{Available: 50, Total: 64},
		},
		Lights: huev1.LightsCapabilities{
			ResourceUsage: huev1.ResourceUsage{Available: 31, Total: 63},
		},
		ResourceLinks: huev1.ResourceLinksCapabilities{
			ResourceUsage: huev1.ResourceUsage{Available: 38, Total: 64},
		},
		Rules: huev1.RulesCapabilities{
			ResourceUsage: huev1.ResourceUsage{Available: 50, Total: 250},
			Actions:       huev1.ResourceUsage{Available: 544, Total: 1000},
			Conditions:    huev1.ResourceUsage{Available: 867, Total: 1500},
		},
		Scenes: huev1.ScenesCapabilities{
			ResourceUsage: huev1.ResourceUsage{Available: 74, Total: 200},
			LightStates:   huev1.ResourceUsage{Available: 3004, Total: 12600},
		},
		Schedules: huev1.SchedulesCapabilities{
			ResourceUsage: huev1.ResourceUsage{Available: 78, Total: 100},
		},
		Sensors: huev1.SensorsCapabilities{
			ResourceUsage: huev1.ResourceUsage{Available: 181, Total: 250},
			Clip:          huev1.ResourceUsage{Available: 181, Total: 250},
			Zgb:           huev1.ResourceUsage{Available: 0, Total: 0},
			Zll:           huev1.ResourceUsage{Available: 48, Total: 64},
		},
		Streaming: huev1.StreamingCapabilities{
			ResourceUsage: huev1.ResourceUsage{Available: 0, Total: 1},
			Channels:      20,
		},
		TimeZones: huev1.TimeZoneCapabilities{
			Values: []string{"CET", "CST6CDT", "EET", "EST", "EST5EDT", "HST", "MET", "MST", "MST7MDT", "PST8PDT", "WET"},
		},
	}

	HueColorLamp7 = huev1.Light{
		ID:               "1",
		ModelID:          "LCT007",
		ManufacturerName: "Philips",
		Name:             "Hue color lamp 7",
		ProductName:      "Hue color lamp",
		SWVersion:        "5.105.0.21169",
		Type:             "Extended color light",
		UniqueID:         "00:17:88:01:00:bd:c7:b9-0b",
		Capabilities: huev1.LightCapabilities{
			Certified: true,
			Control: huev1.LightControlCapabilities{
				MinDimLevel:    5000,
				MaxLumen:       600,
				ColorGamutType: "B",
				ColorGamut: [][]float64{
					{0.675, 0.322},
					{0.409, 0.518},
					{0.167, 0.04},
				},
				CT: huev1.LightColorTemperature{
					Min: 153,
					Max: 500,
				},
			},
			Streaming: huev1.LightStreamingCapabilities{
				Renderer: true,
				Proxy:    false,
			},
		},
		Config: huev1.LightConfig{
			Archetype: "sultanbulb",
			Function:  "mixed",
			Direction: "omnidirectional",
		},
		State: huev1.LightState{
			On:        false,
			Bri:       1,
			Hue:       33761,
			Sat:       254,
			Effect:    "none",
			XY:        []float64{0.3171, 0.3366},
			CT:        159,
			Alert:     "none",
			ColorMode: "xy",
			Mode:      "homeautomation",
			Reachable: true,
		},
		SWUpdate: huev1.LightSWUpdate{
			State: "noupdates",
			LastInstall: &huev1.AbsoluteTime{
				Time: mustParseTime(huev1.AbsoluteTimeFormat, "2018-01-02T19:24:20"),
			},
		},
	}
)
