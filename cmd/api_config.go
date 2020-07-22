package cmd

import (
	"fmt"
	"os"

	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// huecfg api config ...
type apiConfigCmd struct {
	CreateUser *apiConfigCreateUserCmd `command:"create-user" description:"Create user to interact with Hue Bridge"`
	Dump       *apiConfigDumpCmd       `command:"dump" description:"Fetch the full state of the device in a single JSON document"`
	Get        *apiConfigGetCmd        `command:"get" description:"Show Hue Bridge configuration"`
}

// huecfg api config create-user
type apiConfigCreateUserCmd struct {
	DeviceType        string `long:"device-type" description:"A string in the format '<application_name>#<devicename>'" default:"huecfg#cli"`
	GenerateClientKey bool   `long:"generate-client-key" description:"Generate a random username"`
}

func (c *apiConfigCreateUserCmd) Execute(args []string) error {
	if err := errorOnUnknownArgs(args); err != nil {
		return err
	}

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

// huecfg api config dump
//go:generate ./gen_api_read.sh ID=config_dump TYPE=apiConfigDumpCmd FUNC_CALL=bridge.GetFullState()
type apiConfigDumpCmd struct{}

// huecfg api config get
//go:generate ./gen_api_read.sh ID=config_get TYPE=apiConfigGetCmd FUNC_CALL=bridge.GetConfig()
type apiConfigGetCmd struct{}
