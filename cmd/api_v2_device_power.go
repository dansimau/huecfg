package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2DevicePowerCmd struct {
	Get *apiV2DevicePowerGetCmd `command:"get"`
	Put *apiV2DevicePowerPutCmd `command:"put"`
}

type apiV2DevicePowerGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the device power" optional:"true"`
	} `positional-args:"true" positional-arg-name:"device-power-ID"`
}

func (c *apiV2DevicePowerGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID != "" {
		respBytes, err = huev2.GetDevicePower(c.Arguments.ID)
	} else {
		respBytes, err = huev2.GetDevicePowers()
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2DevicePowerPutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the device power"`
	} `positional-args:"true" required:"true" positional-arg-name:"device-power-ID"`
}

func (c *apiV2DevicePowerPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutDevicePower(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
