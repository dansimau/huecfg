package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2HomekitCmd struct {
	Get *apiV2HomekitGetCmd `command:"get"`
	Put *apiV2HomekitPutCmd `command:"put"`
}

type apiV2HomekitGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the homekit" optional:"true"`
	} `positional-args:"true" positional-arg-name:"homekit-ID"`
}

func (c *apiV2HomekitGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID != "" {
		respBytes, err = huev2.GetHomekit(c.Arguments.ID)
	} else {
		respBytes, err = huev2.GetHomekits()
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2HomekitPutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the homekit"`
	} `positional-args:"true" required:"true" positional-arg-name:"homekit-ID"`
}

func (c *apiV2HomekitPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutHomekit(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
