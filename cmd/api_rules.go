package cmd

import (
	"fmt"
	"os"

	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// huecfg api rules ...
type apiRulesCmd struct{}

func (c *apiRulesCmd) Execute(args []string) error {
	bridge := api.getHueAPI()

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
type apiRulesCmdGet struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"rule-ID"`
}

func (c *apiRulesCmdGet) Execute(args []string) error {
	bridge := api.getHueAPI()

	respBytes, err := bridge.GetRule(c.Arguments.ID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}
