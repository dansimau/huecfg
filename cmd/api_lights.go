package cmd

import (
	"fmt"
	"os"

	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// huecfg api lights ...
type apiLightsCmd struct {
	All    *apiLightsAllCmd    `command:"all" description:"Gets a list of all lights that have been discovered by the bridge."`
	Delete *apiLightsDeleteCmd `command:"delete" description:"Delete a device from the bridge."`
	Get    *apiLightsGetCmd    `command:"get" description:"Gets the attributes and state of a given light."`
	Search *apiLightsSearchCmd `command:"search" description:"Search for new devices."`
}

// huecfg api lights all
type apiLightsAllCmd struct{}

func (c *apiLightsAllCmd) Execute(args []string) error {
	bridge := cmd.getHueAPI()

	respBytes, err := bridge.GetLights()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}

// huecfg api lights delete ...
type apiLightsDeleteCmd struct {
	Arguments struct {
		ID string `description:"ID of the device to delete."`
	} `positional-args:"true" required:"true" positional-arg-name:"device-ID"`
}

func (c *apiLightsDeleteCmd) Execute(args []string) error {
	bridge := cmd.getHueAPI()

	respBytes, err := bridge.DeleteLight(c.Arguments.ID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}

// huecfg api lights get ...
type apiLightsGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the light to get attributes of."`
	} `positional-args:"true" required:"true" positional-arg-name:"light-ID"`
}

func (c *apiLightsGetCmd) Execute(args []string) error {
	bridge := cmd.getHueAPI()

	respBytes, err := bridge.GetLight(c.Arguments.ID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}

// huecfg api lights search
type apiLightsSearchCmd struct {
	Arguments struct {
		IDs []string `description:"ID of the light to get attributes of."`
	} `positional-args:"true" required:"false" positional-arg-name:"device-ID"`
}

func (c *apiLightsSearchCmd) Execute(args []string) error {
	bridge := cmd.getHueAPI()

	respBytes, err := bridge.SearchForNewLights(c.Arguments.IDs...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}
