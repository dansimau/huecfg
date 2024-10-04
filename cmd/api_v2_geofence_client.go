package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2GeofenceClientCmd struct{}

func (c *apiV2GeofenceClientCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.GetGeofenceClients()
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
