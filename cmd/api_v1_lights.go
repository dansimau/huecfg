package cmd

type apiV1LightsCmd struct {
	Delete   *apiLightsDeleteCmd   `command:"delete" description:"Delete a device from the bridge"`
	Get      *apiLightsGetCmd      `command:"get" description:"Fetch the specified light by ID"`
	GetAll   *apiLightsGetAllCmd   `command:"get-all" description:"Fetch all lights at once"`
	Search   *apiLightsSearchCmd   `command:"search" description:"Search for new lights"`
	Set      *apiLightsSetCmd      `command:"set" description:"Set the attributes of a given light"`
	SetState *apiLightsSetStateCmd `command:"set-state" description:"Set the state of a given light"`
}

//go:generate ./gen_api_read.sh ID=lights_delete TYPE=apiV1LightsDeleteCmd FUNC_CALL=bridge.DeleteLight(c.Arguments.ID)
type apiLightsDeleteCmd struct {
	Arguments struct {
		ID string `description:"ID of the device to delete."`
	} `positional-args:"true" required:"true" positional-arg-name:"device-ID"`
}

//go:generate ./gen_api_read.sh ID=lights_get TYPE=apiV1LightsGetCmd FUNC_CALL=bridge.GetLight(c.Arguments.ID)
type apiLightsGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the light to get attributes of."`
	} `positional-args:"true" required:"true" positional-arg-name:"light-ID"`
}

//go:generate ./gen_api_read.sh ID=lights_all TYPE=apiV1LightsGetAllCmd FUNC_CALL=bridge.GetLights()
type apiLightsGetAllCmd struct{}

//go:generate ./gen_api_read.sh ID=lights_search TYPE=apiV1LightsSearchCmd FUNC_CALL=bridge.SearchForNewLights(c.Arguments.IDs...)
type apiLightsSearchCmd struct {
	Arguments struct {
		IDs []string `description:"ID of the light to get attributes of."`
	} `positional-args:"true" required:"false" positional-arg-name:"device-ID"`
}

//go:generate ./gen_api_write.sh ID=lights_set TYPE=apiV1LightsSetCmd "FUNC_CALL=bridge.SetLightAttributes(c.Arguments.ID, data)" DATA=c.Data
type apiLightsSetCmd struct {
	Data      string `long:"data" description:"JSON data to send"`
	Arguments struct {
		ID string `description:"ID of the light to set attributes for."`
	} `positional-args:"true" required:"true" positional-arg-name:"light-ID"`
}

//go:generate ./gen_api_write.sh ID=lights_set_state TYPE=apiV1LightsSetStateCmd "FUNC_CALL=bridge.SetLightState(c.Arguments.ID, data)" DATA=c.Data
type apiLightsSetStateCmd struct {
	Data      string `long:"data" description:"JSON data to send" default:"-"`
	Arguments struct {
		ID string `description:"ID of the light to set state for."`
	} `positional-args:"true" required:"true" positional-arg-name:"light-ID"`
}
