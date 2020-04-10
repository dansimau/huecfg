package cmd

import (
	"fmt"
	"os"

	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// huecfg api lights ...
type apiLightsCmd struct {
	All *apiLightsCmdAll `command:"all" description:"Gets a list of all lights that have been discovered by the bridge."`
	Get *apiLightsCmdGet `command:"get" description:"Gets the attributes and state of a given light."`
}

// huecfg api lights all
type apiLightsCmdAll struct{}

func (c *apiLightsCmdAll) Execute(args []string) error {
	bridge := api.getHueAPI()

	respBytes, err := bridge.GetLights()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}

// huecfg api lights get ...
type apiLightsCmdGet struct {
	Arguments struct {
		LightID int `description:"ID of the light to get attributes of."`
	} `positional-args:"true" required:"true" positional-arg-name:"light-ID"`
}

func (c *apiLightsCmdGet) Execute(args []string) error {
	bridge := api.getHueAPI()

	respBytes, err := bridge.GetLight(c.Arguments.LightID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}
