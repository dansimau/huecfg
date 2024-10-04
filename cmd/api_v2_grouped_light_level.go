package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2GroupedLightLevelCmd struct {
	Get *apiV2GroupedLightLevelGetCmd `command:"get"`
	Put *apiV2GroupedLightLevelPutCmd `command:"put"`
}

type apiV2GroupedLightLevelGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the grouped light level" optional:"true"`
	} `positional-args:"true" positional-arg-name:"grouped-light-level-ID"`
}

func (c *apiV2GroupedLightLevelGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID != "" {
		respBytes, err = huev2.GetGroupedLightLevel(c.Arguments.ID)
	} else {
		respBytes, err = huev2.GetGroupedLightLevels()
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2GroupedLightLevelPutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the grouped light level"`
	} `positional-args:"true" required:"true" positional-arg-name:"grouped-light-level-ID"`
}

func (c *apiV2GroupedLightLevelPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutGroupedLightLevel(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
