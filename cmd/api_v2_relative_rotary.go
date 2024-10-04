package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2RelativeRotaryCmd struct {
	Get *apiV2RelativeRotaryGetCmd `command:"get"`
	Put *apiV2RelativeRotaryPutCmd `command:"put"`
}

type apiV2RelativeRotaryGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the relative rotary" optional:"true"`
	} `positional-args:"true" positional-arg-name:"relative-rotary-ID"`
}

func (c *apiV2RelativeRotaryGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID != "" {
		respBytes, err = huev2.GetRelativeRotary(c.Arguments.ID)
	} else {
		respBytes, err = huev2.GetRelativeRotaries()
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2RelativeRotaryPutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the relative rotary"`
	} `positional-args:"true" required:"true" positional-arg-name:"relative-rotary-ID"`
}

func (c *apiV2RelativeRotaryPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutRelativeRotary(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
