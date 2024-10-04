package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2MatterFabricCmd struct {
	Get    *apiV2MatterFabricGetCmd    `command:"get"`
	Delete *apiV2MatterFabricDeleteCmd `command:"delete"`
}

type apiV2MatterFabricGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the matter fabric" optional:"true"`
	} `positional-args:"true" positional-arg-name:"matter-fabric-ID"`
}

func (c *apiV2MatterFabricGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID != "" {
		respBytes, err = huev2.GetMatterFabric(c.Arguments.ID)
	} else {
		respBytes, err = huev2.GetMatterFabrics()
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2MatterFabricDeleteCmd struct {
	Arguments struct {
		ID string `description:"ID of the matter fabric"`
	} `positional-args:"true" required:"true" positional-arg-name:"matter-fabric-ID"`
}

func (c *apiV2MatterFabricDeleteCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.DeleteMatterFabric(c.Arguments.ID)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
