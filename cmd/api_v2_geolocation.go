package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2GeolocationCmd struct {
	Get *apiV2GeolocationGetCmd `command:"get"`
	Put *apiV2GeolocationPutCmd `command:"put"`
}

type apiV2GeolocationGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the geolocation" optional:"true"`
	} `positional-args:"true" positional-arg-name:"geolocation-ID"`
}

func (c *apiV2GeolocationGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID != "" {
		respBytes, err = huev2.GetGeolocation(c.Arguments.ID)
	} else {
		respBytes, err = huev2.GetGeolocations()
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2GeolocationPutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the geolocation"`
	} `positional-args:"true" required:"true" positional-arg-name:"geolocation-ID"`
}

func (c *apiV2GeolocationPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutGeolocation(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
