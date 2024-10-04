package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2EntertainmentCmd struct {
	Get *apiV2EntertainmentGetCmd `command:"get"`
	Put *apiV2EntertainmentPutCmd `command:"put"`
}

type apiV2EntertainmentGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the entertainment" optional:"true"`
	} `positional-args:"true" positional-arg-name:"entertainment-ID"`
}

func (c *apiV2EntertainmentGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID != "" {
		respBytes, err = huev2.GetEntertainment(c.Arguments.ID)
	} else {
		respBytes, err = huev2.GetEntertainments()
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2EntertainmentPutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the entertainment"`
	} `positional-args:"true" required:"true" positional-arg-name:"entertainment-ID"`
}

func (c *apiV2EntertainmentPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutEntertainment(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
