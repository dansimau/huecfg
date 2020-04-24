package cmd

import (
	"fmt"
	"os"

	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// huecfg api resourcelinks ...
type apiResourceLinksCmd struct{}

func (c *apiResourceLinksCmd) Execute(args []string) error {
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
type apiResourceLinksCmdGet struct {
	Arguments struct {
		ID int
	} `positional-args:"true" required:"true" positional-arg-name:"resourcelink-ID"`
}

func (c *apiResourceLinksCmdGet) Execute(args []string) error {
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
