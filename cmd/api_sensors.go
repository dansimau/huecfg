package cmd

import (
	"fmt"
	"os"

	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// huecfg api sensors ...
type apiSensorsCmd struct {
	All *apiSensorsAllCmd `command:"all" description:"Gets a list of all sensors that have been added to the bridge."`
	Get *apiSensorsGetCmd `command:"get" description:"Gets the sensor from the bridge with the given ID."`
}

type apiSensorsAllCmd struct{}

func (c *apiSensorsAllCmd) Execute(args []string) error {
	bridge := cmd.getHueAPI()

	respBytes, err := bridge.GetSensors()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}

// huecfg api sensors get ...
type apiSensorsGetCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"sensor-ID"`
}

func (c *apiSensorsGetCmd) Execute(args []string) error {
	bridge := cmd.getHueAPI()

	respBytes, err := bridge.GetSensor(c.Arguments.ID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}
