package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2GroupedLightCmd struct {
	Get *apiV2GroupedLightGetCmd `command:"get"`
	Put *apiV2GroupedLightPutCmd `command:"put"`
}

type apiV2GroupedLightGetCmd struct{}

func (c *apiV2GroupedLightGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.GetGroupedLights()
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2GroupedLightPutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the grouped light"`
	} `positional-args:"true" required:"true" positional-arg-name:"grouped-light-ID"`
}

func (c *apiV2GroupedLightPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutGroupedLight(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
