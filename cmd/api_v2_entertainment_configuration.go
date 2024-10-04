package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2EntertainmentConfigurationCmd struct {
	Get    *apiV2EntertainmentConfigurationGetCmd    `command:"get"`
	Post   *apiV2EntertainmentConfigurationPostCmd   `command:"post"`
	Put    *apiV2EntertainmentConfigurationPutCmd    `command:"put"`
	Delete *apiV2EntertainmentConfigurationDeleteCmd `command:"delete"`
}

type apiV2EntertainmentConfigurationGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the entertainment configuration" optional:"true"`
	} `positional-args:"true" positional-arg-name:"entertainment-configuration-ID"`
}

func (c *apiV2EntertainmentConfigurationGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID != "" {
		respBytes, err = huev2.GetEntertainmentConfiguration(c.Arguments.ID)
	} else {
		respBytes, err = huev2.GetEntertainmentConfigurations()
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2EntertainmentConfigurationPostCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`
}

func (c *apiV2EntertainmentConfigurationPostCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PostEntertainmentConfiguration(c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2EntertainmentConfigurationPutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the entertainment configuration"`
	} `positional-args:"true" required:"true" positional-arg-name:"entertainment-configuration-ID"`
}

func (c *apiV2EntertainmentConfigurationPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutEntertainmentConfiguration(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2EntertainmentConfigurationDeleteCmd struct {
	Arguments struct {
		ID string `description:"ID of the entertainment configuration"`
	} `positional-args:"true" required:"true" positional-arg-name:"entertainment-configuration-ID"`
}

func (c *apiV2EntertainmentConfigurationDeleteCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.DeleteEntertainmentConfiguration(c.Arguments.ID)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
