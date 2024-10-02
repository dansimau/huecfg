package cmd

type apiV1SchedulesCmd struct {
	Create *apiSchedulesCreateCmd `command:"create" description:"Create a new scehdule"`
	Delete *apiSchedulesDeleteCmd `command:"delete" description:"Delete a device from the bridge"`
	Get    *apiSchedulesGetCmd    `command:"get" description:"Fetch the specified schedule by ID"`
	GetAll *apiSchedulesGetAllCmd `command:"get-all" description:"Fetch all schedule data at once"`
}

//go:generate ./gen_api_write.sh ID=schedules_create TYPE=apiV1SchedulesCreateCmd DATA=c.Data FUNC_CALL=bridge.CreateSchedule(data)
type apiSchedulesCreateCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`
}

//go:generate ./gen_api_read.sh ID=schedules_delete TYPE=apiV1SchedulesDeleteCmd FUNC_CALL=bridge.DeleteSchedule(c.Arguments.ID)
type apiSchedulesDeleteCmd struct {
	Arguments struct {
		ID string `description:"ID of the schedule to delete."`
	} `positional-args:"true" required:"true" positional-arg-name:"schedule-ID"`
}

//go:generate ./gen_api_read.sh ID=schedules_all TYPE=apiV1SchedulesGetAllCmd FUNC_CALL=bridge.GetSchedules()
type apiSchedulesGetAllCmd struct{}

//go:generate ./gen_api_read.sh ID=schedules_get TYPE=apiV1SchedulesGetCmd FUNC_CALL=bridge.GetSchedule(c.Arguments.ID)
type apiSchedulesGetCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"schedule-ID"`
}
