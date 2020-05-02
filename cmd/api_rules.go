package cmd

import (
	"fmt"
	"os"

	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// huecfg api rules ...
type apiRulesCmd struct {
	All *apiRulesAllCmd `command:"all" description:"Gets a list of all rules that are in the bridge."`
	Get *apiRulesGetCmd `command:"get" description:"Returns the rule object with specified ID."`
}

type apiRulesAllCmd struct{}

func (c *apiRulesAllCmd) Execute(args []string) error {
	bridge := cmd.getHueAPI()

	respBytes, err := bridge.GetRules()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}

// huecfg api rules get ...
type apiRulesGetCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"rule-ID"`
}

func (c *apiRulesGetCmd) Execute(args []string) error {
	bridge := cmd.getHueAPI()

	respBytes, err := bridge.GetRule(c.Arguments.ID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}
