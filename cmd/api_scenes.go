package cmd

// huecfg api scenes ...
type apiScenesCmd struct {
	Create *apiScenesCreateCmd `command:"create" description:"Create a new scene"`
	Get    *apiLightsGetCmd    `command:"get" description:"Fetch the specified scene by ID"`
	GetAll *apiLightsGetAllCmd `command:"get-all" description:"Fetch all scenes at once"`
}

// huecfg api scenes create
//go:generate ./gen_api_write.sh ID=scenes_create TYPE=apiScenesCreateCmd DATA=c.Data FUNC_CALL=bridge.CreateResourceLink(data)
type apiScenesCreateCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`
}

// huecfg api scenes get-all
//go:generate ./gen_api_read.sh ID=scenes_get_all TYPE=apiScenesGetAllCmd FUNC_CALL=bridge.GetScenes()
type apiScenesGetAllCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"scene-ID"`
}

// huecfg api scenes get
//go:generate ./gen_api_read.sh ID=scenes_get TYPE=apiScenesGetCmd FUNC_CALL=bridge.GetScene(c.Arguments.ID)
type apiScenesGetCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"scene-ID"`
}
