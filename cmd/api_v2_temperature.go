package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2TemperatureCmd struct {
	Get *apiV2TemperatureGetCmd `command:"get" description:"Get all temperature resources"`
	Put *apiV2TemperaturePutCmd `command:"put" description:"Update a temperature resource"`
}

type apiV2TemperatureGetCmd struct {
	Arguments struct {
		ID string `positional-arg-name:"id" required:"false" description:"ID of the temperature resource"`
	} `positional-args:"yes"`
}

func (c *apiV2TemperatureGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID == "" {
		respBytes, err = huev2.GetTemperatures()
	} else {
		respBytes, err = huev2.GetTemperature(c.Arguments.ID)
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2TemperaturePutCmd struct {
	Arguments struct {
		ID   string `positional-arg-name:"id" required:"true" description:"ID of the temperature resource"`
		Data string `positional-arg-name:"data" required:"true" description:"JSON data to update the temperature resource"`
	} `positional-args:"yes"`
}

func (c *apiV2TemperaturePutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	data := []byte(c.Arguments.Data)
	respBytes, err := huev2.PutTemperature(c.Arguments.ID, data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
