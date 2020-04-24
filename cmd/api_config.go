package cmd

import (
	"fmt"
	"os"

	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// huecfg api config ...
type apiConfigCmd struct {
	CreateUser *apiConfigCreateUserCmd `command:"create-user" description:"Create user to interact with Hue Bridge"`
	Get        *apiConfigGetCmd        `command:"get" description:"Show Hue Bridge configuration"`
}

// huecfg api config get ...
type apiConfigGetCmd struct{}

func (c *apiConfigGetCmd) Execute(args []string) error {
	bridge := cmd.getHueAPI()

	respBytes, err := bridge.GetConfig()
	if err != nil {
		return err
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}

// huecfg api config create-user ...
type apiConfigCreateUserCmd struct {
	DeviceType        string `long:"device-type" description:"A string in the format '<application_name>#<devicename>'" default:"huecfg#cli"`
	GenerateClientKey bool   `long:"generate-client-key" description:"Generate a random username"`
}

func (c *apiConfigCreateUserCmd) Execute(args []string) error {
	bridge := cmd.getHueAPI()

	respBytes, err := bridge.CreateUser(c.DeviceType, c.GenerateClientKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}
