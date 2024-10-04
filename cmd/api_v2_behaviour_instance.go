package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2BehaviorInstanceCmd struct{}

func (c *apiV2BehaviorInstanceCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.GetBehaviorInstances()
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
