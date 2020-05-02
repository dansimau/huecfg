package cmd

import (
	"fmt"
	"os"

	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// huecfg api schedules ...
type apiSchedulesCmd struct {
	All *apiSchedulesAllCmd `command:"all" description:"Gets a list of all schedules that have been added to the bridge."`
	Get *apiSchedulesGetCmd `command:"get" description:"Gets all attributes for a schedule."`
}

type apiSchedulesAllCmd struct{}

func (c *apiSchedulesAllCmd) Execute(args []string) error {
	bridge := cmd.getHueAPI()

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
type apiSchedulesGetCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"schedule-ID"`
}

func (c *apiSchedulesGetCmd) Execute(args []string) error {
	bridge := cmd.getHueAPI()

	respBytes, err := bridge.GetSchedule(c.Arguments.ID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}
