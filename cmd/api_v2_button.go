package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2ButtonCmd struct{}

func (c *apiV2ButtonCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.GetButtons()
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
