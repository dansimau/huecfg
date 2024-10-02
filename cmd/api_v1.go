package cmd

type apiV1Cmd struct {
	APICapabilities  *apiV1CapabilitiesCmd  `command:"capabilities" description:"Show capabilities and resource usage"`
	APIConfig        *apiV1ConfigCmd        `command:"config" description:"Manage config"`
	APIGroups        *apiV1GroupsCmd        `command:"groups" description:"Manage groups"`
	APILights        *apiV1LightsCmd        `command:"lights" description:"Manage lights"`
	APIResourceLinks *apiV1ResourceLinksCmd `command:"resourcelinks" description:"Manage resource links"`
	APIRules         *apiV1RulesCmd         `command:"rules" description:"Manage rules"`
	APIScenes        *apiV1ScenesCmd        `command:"scenes" description:"Manage scenes"`
	APISchedules     *apiV1SchedulesCmd     `command:"schedules" description:"Manage schedules"`
	APISensors       *apiV1SensorsCmd       `command:"sensors" description:"Manage sensors"`
}
