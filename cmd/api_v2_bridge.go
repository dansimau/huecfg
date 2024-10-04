package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2BridgeCmd struct {
	Get *apiV2BridgeGetCmd `command:"get"`
	Put *apiV2BridgePutCmd `command:"put"`
}

type apiV2BridgeGetCmd struct{}

func (c *apiV2BridgeGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.GetBridges()
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2BridgePutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the bridge"`
	} `positional-args:"true" required:"true" positional-arg-name:"bridge-ID"`
}

func (c *apiV2BridgePutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutBridge(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
