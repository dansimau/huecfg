// Code generated by go generate; DO NOT EDIT.
package cmd

import (
	"github.com/dansimau/huecfg/pkg/jsonutil"
)

func (c *apiLightsSearchCmd) Execute(args []string) error {
	if err := errorOnUnknownArgs(args); err != nil {
		return err
	}

	bridge := cmd.getHueAPI()

	respBytes, err := bridge.SearchForNewLights(c.Arguments.IDs...)
	if err != nil {
		return err
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}
