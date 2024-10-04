package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2DeviceSoftwareUpdateCmd struct {
	Get *apiV2DeviceSoftwareUpdateGetCmd `command:"get"`
	Put *apiV2DeviceSoftwareUpdatePutCmd `command:"put"`
}

type apiV2DeviceSoftwareUpdateGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the device software update" optional:"true"`
	} `positional-args:"true" positional-arg-name:"device-software-update-ID"`
}

func (c *apiV2DeviceSoftwareUpdateGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID != "" {
		respBytes, err = huev2.GetDeviceSoftwareUpdate(c.Arguments.ID)
	} else {
		respBytes, err = huev2.GetDeviceSoftwareUpdates()
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2DeviceSoftwareUpdatePutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the device software update"`
	} `positional-args:"true" required:"true" positional-arg-name:"device-software-update-ID"`
}

func (c *apiV2DeviceSoftwareUpdatePutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutDeviceSoftwareUpdate(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
