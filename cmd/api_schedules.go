package cmd

// huecfg api schedules ...
type apiSchedulesCmd struct {
	Create *apiSchedulesCreateCmd `command:"create" description:"Create a new scehdule"`
	Get    *apiSchedulesGetCmd    `command:"get" description:"Fetch the specified schedule by ID"`
	GetAll *apiSchedulesGetAllCmd `command:"get-all" description:"Fetch all schedule data at once"`
}

// huecfg api schedules create
//go:generate ./gen_api_write.sh ID=schedules_create TYPE=apiSchedulesCreateCmd DATA=c.Data FUNC_CALL=bridge.CreateResourceLink(data)
type apiSchedulesCreateCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`
}

//go:generate ./gen_api_read.sh ID=schedules_all TYPE=apiSchedulesGetAllCmd FUNC_CALL=bridge.GetSchedules()
type apiSchedulesGetAllCmd struct{}

// huecfg api schedules get ...
//go:generate ./gen_api_read.sh ID=schedules_get TYPE=apiSchedulesGetCmd FUNC_CALL=bridge.GetSchedule(c.Arguments.ID)
type apiSchedulesGetCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"schedule-ID"`
}
