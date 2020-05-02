package cmd

import (
	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// huecfg api capabilities
type apiCapabilitiesCmd struct{}

func (c *apiCapabilitiesCmd) Execute(args []string) error {
	bridge := cmd.getHueAPI()

	respBytes, err := bridge.GetCapabilities()
	if err != nil {
		return err
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}
