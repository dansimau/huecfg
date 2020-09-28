package fixtures

import (
	"github.com/dansimau/huecfg/pkg/hue"
)

var (
	Default = []byte(`
		{
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

	DefaultHueColorLamp7 = hue.Light{
		ID:               "1",
		ModelID:          "LCT007",
		ManufacturerName: "Philips",
		Name:             "Hue color lamp 7",
		ProductName:      "Hue color lamp",
		SWVersion:        "5.105.0.21169",
		Type:             "Extended color light",
		UniqueID:         "00:17:88:01:00:bd:c7:b9-0b",
		Capabilities: hue.LightCapabilities{
			Certified: true,
			Control: hue.LightControlCapabilities{
				MinDimLevel:    5000,
				MaxLumen:       600,
				ColorGamutType: "B",
				ColorGamut: [][]float64{
					[]float64{0.675, 0.322},
					[]float64{0.409, 0.518},
					[]float64{0.167, 0.04},
				},
				CT: hue.LightColorTemperature{
					Min: 153,
					Max: 500,
				},
			},
			Streaming: hue.LightStreamingCapabilities{
				Renderer: true,
				Proxy:    false,
			},
		},
		Config: hue.LightConfig{
			Archetype: "sultanbulb",
			Function:  "mixed",
			Direction: "omnidirectional",
		},
		State: hue.LightState{
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
		SWUpdate: hue.LightSWUpdate{
			State: "noupdates",
			LastInstall: &hue.AbsoluteTime{
				Time: mustParseTime(hue.AbsoluteTimeFormat, "2018-01-02T19:24:20"),
			},
		},
	}
)
