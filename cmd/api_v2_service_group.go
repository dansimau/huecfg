package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2ServiceGroupCmd struct{}

func (c *apiV2ServiceGroupCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.GetServiceGroups()
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
