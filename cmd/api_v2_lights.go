package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2LightCmd struct {
	Get *apiV2LightGetCmd `command:"get"`
	Put *apiV2LightPutCmd `command:"put"`
}

type apiV2LightGetCmd struct{}

func (c *apiV2LightGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.GetLights()
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2LightPutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the light"`
	} `positional-args:"true" required:"true" positional-arg-name:"light-ID"`
}

func (c *apiV2LightPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutLight(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
