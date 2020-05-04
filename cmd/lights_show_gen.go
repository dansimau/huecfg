// Code generated by go generate; DO NOT EDIT.
package cmd

import (
	"fmt"
	"os"

	"github.com/mikefarah/yq/v3/pkg/yqlib"
	"gopkg.in/yaml.v2"
)

type lightsShowCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"light-ID"`
}

func (c *lightsShowCmd) Execute(args []string) error {
	bridge := cmd.getHue()

	resp, err := bridge.GetLight(c.Arguments.ID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bytes, err := yaml.Marshal(resp)
	if err != nil {
		return err
	}

	return yqlib.ColorizeAndPrint(bytes, os.Stdout)
}