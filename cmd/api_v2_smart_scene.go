package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2SmartSceneCmd struct{}

func (c *apiV2SmartSceneCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.GetSmartScenes()
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
