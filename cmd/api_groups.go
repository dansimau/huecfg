package cmd

// huecfg api groups ...
type apiGroupsCmd struct {
	Create   *apiGroupsCreateCmd   `command:"create" description:"Create a new group"`
	Delete   *apiGroupsDeleteCmd   `command:"delete" description:"Delete a device from the bridge"`
	Get      *apiGroupsGetCmd      `command:"get" description:"Fetch the specified group by ID"`
	GetAll   *apiGroupsGetAllCmd   `command:"get-all" description:"Fetch all group data at once"`
	Set      *apiGroupsSetCmd      `command:"set" description:"Set attributes of a group"`
	SetState *apiGroupsSetStateCmd `command:"set-state" description:"Set the state of all lights in a group"`
}

// huecfg api groups create
//go:generate ./gen_api_write.sh ID=groups_create TYPE=apiGroupsCreateCmd DATA=c.Data FUNC_CALL=bridge.CreateGroup(data)
type apiGroupsCreateCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`
}

// huecfg api groups delete ...
//go:generate ./gen_api_read.sh ID=groups_delete TYPE=apiGroupsDeleteCmd FUNC_CALL=bridge.DeleteGroup(c.Arguments.ID)
type apiGroupsDeleteCmd struct {
	Arguments struct {
		ID string `description:"ID of the group to delete."`
	} `positional-args:"true" required:"true" positional-arg-name:"group-ID"`
}

// huecfg api groups get-all
//go:generate ./gen_api_read.sh ID=groups_get_all TYPE=apiGroupsGetAllCmd FUNC_CALL=bridge.GetGroups()
type apiGroupsGetAllCmd struct{}

// huecfg api groups get
//go:generate ./gen_api_read.sh ID=groups_get TYPE=apiGroupsGetCmd FUNC_CALL=bridge.GetGroup(c.Arguments.ID)
type apiGroupsGetCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"group-ID"`
}

// huecfg api groups set
//go:generate ./gen_api_write.sh ID=groups_set TYPE=apiGroupsSetCmd "FUNC_CALL=bridge.SetGroupAttributes(c.Arguments.ID, data)" DATA=c.Data
type apiGroupsSetCmd struct {
	Data      string `long:"data" description:"JSON data to send"`
	Arguments struct {
		ID string `description:"ID of the group to set attributes for."`
	} `positional-args:"true" required:"true" positional-arg-name:"light-ID"`
}

// huecfg api groups set-state ...
//go:generate ./gen_api_write.sh ID=groups_set_state TYPE=apiGroupsSetStateCmd "FUNC_CALL=bridge.SetGroupState(c.Arguments.ID, data)" DATA=c.Data
type apiGroupsSetStateCmd struct {
	Data      string `long:"data" description:"JSON data to send" default:"-"`
	Arguments struct {
		ID string `description:"ID of the light to set state for."`
	} `positional-args:"true" required:"true" positional-arg-name:"light-ID"`
}
