package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2GroupedMotionCmd struct {
	Get *apiV2GroupedMotionGetCmd `command:"get"`
	Put *apiV2GroupedMotionPutCmd `command:"put"`
}

type apiV2GroupedMotionGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the grouped motion" optional:"true"`
	} `positional-args:"true" positional-arg-name:"grouped-motion-ID"`
}

func (c *apiV2GroupedMotionGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID != "" {
		respBytes, err = huev2.GetGroupedMotion(c.Arguments.ID)
	} else {
		respBytes, err = huev2.GetGroupedMotions()
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2GroupedMotionPutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the grouped motion"`
	} `positional-args:"true" required:"true" positional-arg-name:"grouped-motion-ID"`
}

func (c *apiV2GroupedMotionPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutGroupedMotion(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
