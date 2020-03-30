package cmd

import (
	"fmt"
	"os"

	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// huecfg api schedules ...
type apiSchedulesCmd struct{}

func (c *apiSchedulesCmd) Execute(args []string) error {
	bridge := api.getHueAPI()

	respBytes, err := bridge.GetSchedules()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}

// huecfg api schedules get ...
type apiSchedulesCmdGet struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"schedule-ID"`
}

func (c *apiSchedulesCmdGet) Execute(args []string) error {
	bridge := api.getHueAPI()

	respBytes, err := bridge.GetSchedule(c.Arguments.ID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}
