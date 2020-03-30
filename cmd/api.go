package cmd

import (
	"github.com/dansimau/huecfg/pkg/hue"
)

func init() {
	_, err := parser.AddCommand("api", "Interact with the Hue Bridge API", "", api)
	if err != nil {
		panic(err)
	}
}

var api = &apiCmd{}

type apiCmd struct {
	Host     string `short:"a" long:"host" description:"host address for Hue Bridge" required:"true"`
	Username string `short:"u" long:"username" description:"username from Hue Bridge registration"`

	APIConfig    *apiConfigCmd    `command:"config" description:"Manage config"`
	APIGroups    *apiGroupsCmd    `command:"groups" description:"Manage groups"`
	APILights    *apiLightsCmd    `command:"lights" description:"Manage lights"`
	APIRules     *apiRulesCmd     `command:"rules" description:"Manage rules"`
	APIScenes    *apiScenesCmd    `command:"scenes" description:"Manage scenes"`
	APISchedules *apiSchedulesCmd `command:"schedules" description:"Manage schedules"`
	APISensors   *apiSensorsCmd   `command:"sensors" description:"Manage sensors"`
}

func (c *apiCmd) getHueAPI() *hue.API {
	h := &hue.API{
		Host:     c.Host,
		Username: c.Username,
	}

	if len(cmd.Verbose) > 0 {
		h.Debug = true
	}

	return h
}
