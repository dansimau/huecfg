package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2BridgeHomeCmd struct {
	Get *apiV2BridgeHomeGetCmd `command:"get"`
}

type apiV2BridgeHomeGetCmd struct{}

func (c *apiV2BridgeHomeGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.GetBridgeHomes()
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
