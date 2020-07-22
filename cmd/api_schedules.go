package cmd

// huecfg api schedules ...
type apiSchedulesCmd struct {
	Get    *apiSchedulesGetCmd    `command:"get" description:"Fetch the specified schedule by ID"`
	GetAll *apiSchedulesGetAllCmd `command:"get-all" description:"Fetch all schedule data at once"`
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
