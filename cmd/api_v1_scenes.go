package cmd

type apiV1ScenesCmd struct {
	Create *apiScenesCreateCmd `command:"create" description:"Create a new scene"`
	Delete *apiScenesDeleteCmd `command:"delete" description:"Delete a device from the bridge"`
	Get    *apiLightsGetCmd    `command:"get" description:"Fetch the specified scene by ID"`
	GetAll *apiLightsGetAllCmd `command:"get-all" description:"Fetch all scenes at once"`
}

//go:generate ./gen_api_write.sh ID=scenes_create TYPE=apiV1ScenesCreateCmd DATA=c.Data FUNC_CALL=bridge.CreateScene(data)
type apiScenesCreateCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`
}

//go:generate ./gen_api_read.sh ID=scenes_delete TYPE=apiV1ScenesDeleteCmd FUNC_CALL=bridge.DeleteScene(c.Arguments.ID)
type apiScenesDeleteCmd struct {
	Arguments struct {
		ID string `description:"ID of the scene to delete."`
	} `positional-args:"true" required:"true" positional-arg-name:"scene-ID"`
}

//go:generate ./gen_api_read.sh ID=scenes_get_all TYPE=apiV1ScenesGetAllCmd FUNC_CALL=bridge.GetScenes()
type apiScenesGetAllCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"scene-ID"`
}

//go:generate ./gen_api_read.sh ID=scenes_get TYPE=apiV1ScenesGetCmd FUNC_CALL=bridge.GetScene(c.Arguments.ID)
type apiScenesGetCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"scene-ID"`
}
