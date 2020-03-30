package cmd

import (
	"fmt"
	"os"

	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// huecfg api scenes ...
type apiScenesCmd struct{}

func (c *apiScenesCmd) Execute(args []string) error {
	bridge := api.getHueAPI()

	respBytes, err := bridge.GetScenes()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}

// huecfg api scenes get ...
type apiScenesCmdGet struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"scene-ID"`
}

func (c *apiScenesCmdGet) Execute(args []string) error {
	bridge := api.getHueAPI()

	respBytes, err := bridge.GetScene(c.Arguments.ID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}
