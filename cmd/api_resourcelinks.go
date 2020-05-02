package cmd

import (
	"fmt"
	"os"

	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// huecfg api resourcelinks ...
type apiResourceLinksCmd struct {
	All *apiResourceLinksAllCmd `command:"all" description:"Gets a list of all resourcelinks that are in the bridge."`
	Get *apiResourceLinksGetCmd `command:"get" description:"Returns resourcelink object with the specified ID."`
}

type apiResourceLinksAllCmd struct{}

func (c *apiResourceLinksAllCmd) Execute(args []string) error {
	bridge := cmd.getHueAPI()

	respBytes, err := bridge.GetResourceLinks()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}

// huecfg api resourcelinks get ...
type apiResourceLinksGetCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"resourcelink-ID"`
}

func (c *apiResourceLinksGetCmd) Execute(args []string) error {
	bridge := cmd.getHueAPI()

	respBytes, err := bridge.GetResourceLink(c.Arguments.ID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}
