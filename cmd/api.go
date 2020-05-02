package cmd

func init() {
	_, err := parser.AddCommand("api", "Interact directly with the Hue Bridge JSON API", "", &apiCmd{})
	if err != nil {
		panic(err)
	}
}

type apiCmd struct {
	APICapabilities  *apiCapabilitiesCmd  `command:"capabilities" description:"Show capabilities and resource usage"`
	APIConfig        *apiConfigCmd        `command:"config" description:"Manage config"`
	APIGroups        *apiGroupsCmd        `command:"groups" description:"Manage groups"`
	APILights        *apiLightsCmd        `command:"lights" description:"Manage lights"`
	APIResourceLinks *apiResourceLinksCmd `command:"resourcelinks" description:"Manage resource links"`
	APIRules         *apiRulesCmd         `command:"rules" description:"Manage rules"`
	APIScenes        *apiScenesCmd        `command:"scenes" description:"Manage scenes"`
	APISchedules     *apiSchedulesCmd     `command:"schedules" description:"Manage schedules"`
	APISensors       *apiSensorsCmd       `command:"sensors" description:"Manage sensors"`
}
