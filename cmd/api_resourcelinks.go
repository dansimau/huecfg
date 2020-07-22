package cmd

// huecfg api resourcelinks ...
type apiResourceLinksCmd struct {
	Create *apiResourceLinksCreateCmd `command:"create" description:"Create a new resourcelink"`
	Get    *apiResourceLinksGetCmd    `command:"get" description:"Fetch the specified resourcelink by ID"`
	GetAll *apiResourceLinksGetAllCmd `command:"get-all" description:"Fetch all resourcelinks at once"`
}

// huecfg api resourcelink create
//go:generate ./gen_api_write.sh ID=resourcelinks_create TYPE=apiResourceLinksCreateCmd DATA=c.Data FUNC_CALL=bridge.CreateResourceLink(data)
type apiResourceLinksCreateCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`
}

//go:generate ./gen_api_read.sh ID=resourcelinks_get_all TYPE=apiResourceLinksGetAllCmd FUNC_CALL=bridge.GetResourceLinks()
type apiResourceLinksGetAllCmd struct{}

// huecfg api resourcelinks get ...
//go:generate ./gen_api_read.sh ID=resourcelinks_get TYPE=apiResourceLinksGetCmd FUNC_CALL=bridge.GetResourceLink(c.Arguments.ID)
type apiResourceLinksGetCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"resourcelink-ID"`
}
