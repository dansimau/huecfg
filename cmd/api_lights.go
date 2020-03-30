package cmd

import (
	"fmt"
	"os"

	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// huecfg api lights ...
type apiLightsCmd struct{}

func (c *apiLightsCmd) Execute(args []string) error {
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
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"light-ID"`
}

func (c *apiLightsCmdGet) Execute(args []string) error {
	bridge := api.getHueAPI()

	respBytes, err := bridge.GetLight(c.Arguments.ID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}
