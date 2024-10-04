package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2ResourceCmd struct {
	Get *apiV2ResourceGetCmd `command:"get"`
}

type apiV2ResourceGetCmd struct{}

func (c *apiV2ResourceGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.GetResources()
	if err != nil {
		return err
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}
