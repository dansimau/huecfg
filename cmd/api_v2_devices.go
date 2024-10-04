package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2DeviceCmd struct {
	Get    *apiV2DeviceGetCmd    `command:"get"`
	Put    *apiV2DevicePutCmd    `command:"put"`
	Delete *apiV2DeviceDeleteCmd `command:"delete"`
}

type apiV2DeviceGetCmd struct{}

func (c *apiV2DeviceGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.GetDevices()
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2DevicePutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the device"`
	} `positional-args:"true" required:"true" positional-arg-name:"device-ID"`
}

func (c *apiV2DevicePutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutDevice(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2DeviceDeleteCmd struct {
	Arguments struct {
		ID string `description:"ID of the device"`
	} `positional-args:"true" required:"true" positional-arg-name:"device-ID"`
}

func (c *apiV2DeviceDeleteCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.DeleteDevice(c.Arguments.ID)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
