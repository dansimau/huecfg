package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2MatterCmd struct {
	Get *apiV2MatterGetCmd `command:"get"`
	Put *apiV2MatterPutCmd `command:"put"`
}

type apiV2MatterGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the matter" optional:"true"`
	} `positional-args:"true" positional-arg-name:"matter-ID"`
}

func (c *apiV2MatterGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID != "" {
		respBytes, err = huev2.GetMatter(c.Arguments.ID)
	} else {
		respBytes, err = huev2.GetMatters()
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2MatterPutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the matter"`
	} `positional-args:"true" required:"true" positional-arg-name:"matter-ID"`
}

func (c *apiV2MatterPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutMatter(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
