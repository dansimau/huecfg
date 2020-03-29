package cmd

import (
	"fmt"
	"os"

	"github.com/dansimau/huecfg/pkg/hue"
	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// huecfg api lights ...
type apiLightsCmd struct{}

func (c *apiLightsCmd) Execute(args []string) error {
	bridge := hue.NewConn(api.Host, api.Username)
	if len(cmd.Verbose) > 0 {
		bridge.Debug = true
	}

	lights, err := bridge.Lights.GetAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(lights.ResponseData.Bytes()); err != nil {
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
	bridge := hue.NewConn(api.Host, api.Username)
	if len(cmd.Verbose) > 0 {
		bridge.Debug = true
	}

	light, err := bridge.Lights.Get(c.Arguments.ID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(light.ResponseData.Bytes()); err != nil {
		return err
	}

	return nil
}
