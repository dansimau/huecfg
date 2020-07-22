package cmd

// huecfg api resourcelinks ...
type apiResourceLinksCmd struct {
	Create *apiResourceLinksCreateCmd `command:"create" description:"Create a new resourcelink"`
	Delete *apiResourceLinksDeleteCmd `command:"delete" description:"Delete a device from the bridge"`
	Get    *apiResourceLinksGetCmd    `command:"get" description:"Fetch the specified resourcelink by ID"`
	GetAll *apiResourceLinksGetAllCmd `command:"get-all" description:"Fetch all resourcelinks at once"`
}

// huecfg api resourcelinks create
//go:generate ./gen_api_write.sh ID=resourcelinks_create TYPE=apiResourceLinksCreateCmd DATA=c.Data FUNC_CALL=bridge.CreateResourceLink(data)
type apiResourceLinksCreateCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`
}

// huecfg api resourcelinks delete ...
//go:generate ./gen_api_read.sh ID=resourcelinks_delete TYPE=apiResourceLinksDeleteCmd FUNC_CALL=bridge.DeleteResourceLink(c.Arguments.ID)
type apiResourceLinksDeleteCmd struct {
	Arguments struct {
		ID string `description:"ID of the resourcelink to delete."`
	} `positional-args:"true" required:"true" positional-arg-name:"resourcelink-ID"`
}

// huecfg api resourcelinks get ...
//go:generate ./gen_api_read.sh ID=resourcelinks_get TYPE=apiResourceLinksGetCmd FUNC_CALL=bridge.GetResourceLink(c.Arguments.ID)
type apiResourceLinksGetCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"resourcelink-ID"`
}

// huecfg api resourcelinks get-all
//go:generate ./gen_api_read.sh ID=resourcelinks_get_all TYPE=apiResourceLinksGetAllCmd FUNC_CALL=bridge.GetResourceLinks()
type apiResourceLinksGetAllCmd struct{}
