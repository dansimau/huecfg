package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2TamperCmd struct {
	Get *apiV2TamperGetCmd `command:"get"`
	Put *apiV2TamperPutCmd `command:"put"`
}

type apiV2TamperGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the tamper" optional:"true"`
	} `positional-args:"true" positional-arg-name:"tamper-ID"`
}

func (c *apiV2TamperGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID != "" {
		respBytes, err = huev2.GetTamper(c.Arguments.ID)
	} else {
		respBytes, err = huev2.GetTampers()
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2TamperPutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the tamper"`
	} `positional-args:"true" required:"true" positional-arg-name:"tamper-ID"`
}

func (c *apiV2TamperPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutTamper(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
