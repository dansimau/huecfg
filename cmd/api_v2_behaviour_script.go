package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2BehaviorScriptCmd struct{}

func (c *apiV2BehaviorScriptCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.GetBehaviorScripts()
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
