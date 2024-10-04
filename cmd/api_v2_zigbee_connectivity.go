package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2ZigbeeConnectivityCmd struct {
	Get *apiV2ZigbeeConnectivityGetCmd `command:"get" description:"Get all Zigbee connectivity resources"`
	Put *apiV2ZigbeeConnectivityPutCmd `command:"put" description:"Update a Zigbee connectivity resource"`
}

type apiV2ZigbeeConnectivityGetCmd struct {
	Arguments struct {
		ID string `positional-arg-name:"id" required:"false" description:"ID of the Zigbee connectivity resource"`
	} `positional-args:"yes"`
}

func (c *apiV2ZigbeeConnectivityGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID == "" {
		respBytes, err = huev2.GetZigbeeConnectivities()
	} else {
		respBytes, err = huev2.GetZigbeeConnectivity(c.Arguments.ID)
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2ZigbeeConnectivityPutCmd struct {
	Arguments struct {
		ID   string `positional-arg-name:"id" required:"true" description:"ID of the Zigbee connectivity resource"`
		Data string `positional-arg-name:"data" required:"true" description:"JSON data to update the Zigbee connectivity resource"`
	} `positional-args:"yes"`
}

func (c *apiV2ZigbeeConnectivityPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	data := []byte(c.Arguments.Data)
	respBytes, err := huev2.PutZigbeeConnectivity(c.Arguments.ID, data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
