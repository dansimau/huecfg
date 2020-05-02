package cmd

import (
	"fmt"
	"os"

	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// huecfg api groups ...
type apiGroupsCmd struct {
	All *apiGroupsAllCmd `command:"all" description:"Gets a list of all groups that have been added to the bridge."`
	//Create *apiGroupsCreateCmd `command:"create" description:"Creates a new group containing the lights specified and optional name."`
	Get *apiGroupsGetCmd `command:"get" description:"Gets the group attributes for a given group."`
}

type apiGroupsAllCmd struct{}

func (c *apiGroupsAllCmd) Execute(args []string) error {
	bridge := cmd.getHueAPI()

	respBytes, err := bridge.GetGroups()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}

// huecfg api groups get ...
type apiGroupsGetCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"group-ID"`
}

func (c *apiGroupsGetCmd) Execute(args []string) error {
	bridge := cmd.getHueAPI()

	respBytes, err := bridge.GetGroup(c.Arguments.ID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}
