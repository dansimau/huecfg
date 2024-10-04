package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2CameraMotionCmd struct {
	Get *apiV2CameraMotionGetCmd `command:"get"`
	Put *apiV2CameraMotionPutCmd `command:"put"`
}

type apiV2CameraMotionGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the camera motion" optional:"true"`
	} `positional-args:"true" positional-arg-name:"camera-motion-ID"`
}

func (c *apiV2CameraMotionGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID != "" {
		respBytes, err = huev2.GetCameraMotion(c.Arguments.ID)
	} else {
		respBytes, err = huev2.GetCameraMotions()
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2CameraMotionPutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the camera motion"`
	} `positional-args:"true" required:"true" positional-arg-name:"camera-motion-ID"`
}

func (c *apiV2CameraMotionPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutCameraMotion(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
