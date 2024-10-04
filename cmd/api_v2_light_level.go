package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2LightLevelCmd struct {
	Get *apiV2LightLevelGetCmd `command:"get"`
	Put *apiV2LightLevelPutCmd `command:"put"`
}

type apiV2LightLevelGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the light level" optional:"true"`
	} `positional-args:"true" positional-arg-name:"light-level-ID"`
}

func (c *apiV2LightLevelGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID != "" {
		respBytes, err = huev2.GetLightLevel(c.Arguments.ID)
	} else {
		respBytes, err = huev2.GetLightLevels()
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2LightLevelPutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the light level"`
	} `positional-args:"true" required:"true" positional-arg-name:"light-level-ID"`
}

func (c *apiV2LightLevelPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutLightLevel(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
