package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2ZoneCmd struct {
	Get    *apiV2ZoneGetCmd    `command:"get"`
	Post   *apiV2ZonePostCmd   `command:"post"`
	Put    *apiV2ZonePutCmd    `command:"put"`
	Delete *apiV2ZoneDeleteCmd `command:"delete"`
}

type apiV2ZoneGetCmd struct{}

func (c *apiV2ZoneGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.GetZones()
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2ZonePostCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`
}

func (c *apiV2ZonePostCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PostZone(c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2ZonePutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the zone"`
	} `positional-args:"true" required:"true" positional-arg-name:"zone-ID"`
}

func (c *apiV2ZonePutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutZone(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2ZoneDeleteCmd struct {
	Arguments struct {
		ID string `description:"ID of the zone"`
	} `positional-args:"true" required:"true" positional-arg-name:"zone-ID"`
}

func (c *apiV2ZoneDeleteCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.DeleteZone(c.Arguments.ID)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
