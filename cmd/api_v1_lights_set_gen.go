// Code generated by go generate; DO NOT EDIT.
package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

func (c *apiLightsSetCmd) Execute(args []string) error {
	if err := errorOnUnknownArgs(args); err != nil {
		return err
	}

	bridge := cmd.getHueAPI()

	data, err := userInputToJSON(c.Data)
	if err != nil {
		return err
	}

	respBytes, err := bridge.SetLightAttributes(c.Arguments.ID, data)
	if err != nil {
		return err
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}