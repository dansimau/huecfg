package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2MotionCmd struct {
	Get *apiV2MotionGetCmd `command:"get"`
	Put *apiV2MotionPutCmd `command:"put"`
}

type apiV2MotionGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the motion" optional:"true"`
	} `positional-args:"true" positional-arg-name:"motion-ID"`
}

func (c *apiV2MotionGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID != "" {
		respBytes, err = huev2.GetMotion(c.Arguments.ID)
	} else {
		respBytes, err = huev2.GetMotions()
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2MotionPutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the motion"`
	} `positional-args:"true" required:"true" positional-arg-name:"motion-ID"`
}

func (c *apiV2MotionPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutMotion(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
