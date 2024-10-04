package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2ZigbeeDeviceDiscoveryCmd struct {
	Get *apiV2ZigbeeDeviceDiscoveryGetCmd `command:"get" description:"Get all Zigbee device discovery resources"`
	Put *apiV2ZigbeeDeviceDiscoveryPutCmd `command:"put" description:"Update a Zigbee device discovery resource"`
}

type apiV2ZigbeeDeviceDiscoveryGetCmd struct {
	Arguments struct {
		ID string `positional-arg-name:"id" required:"false" description:"ID of the Zigbee device discovery resource"`
	} `positional-args:"yes"`
}

func (c *apiV2ZigbeeDeviceDiscoveryGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID == "" {
		respBytes, err = huev2.GetZigbeeDeviceDiscoveries()
	} else {
		respBytes, err = huev2.GetZigbeeDeviceDiscovery(c.Arguments.ID)
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2ZigbeeDeviceDiscoveryPutCmd struct {
	Arguments struct {
		ID   string `positional-arg-name:"id" required:"true" description:"ID of the Zigbee device discovery resource"`
		Data string `positional-arg-name:"data" required:"true" description:"JSON data to update the Zigbee device discovery resource"`
	} `positional-args:"yes"`
}

func (c *apiV2ZigbeeDeviceDiscoveryPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	data := []byte(c.Arguments.Data)
	respBytes, err := huev2.PutZigbeeDeviceDiscovery(c.Arguments.ID, data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
