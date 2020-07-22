package cmd

// huecfg api sensors ...
type apiSensorsCmd struct {
	Get    *apiSensorsGetCmd    `command:"get" description:"Fetch the specified sensor by ID"`
	GetAll *apiSensorsGetAllCmd `command:"get-all" description:"Fetch all sensors at once"`
}

//go:generate ./gen_api_read.sh ID=sensors_get_all TYPE=apiSensorsGetAllCmd FUNC_CALL=bridge.GetSensors()
type apiSensorsGetAllCmd struct{}

// huecfg api sensors get ...
//go:generate ./gen_api_read.sh ID=sensors_get TYPE=apiSensorsGetCmd FUNC_CALL=bridge.GetSensor(c.Arguments.ID)
type apiSensorsGetCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"sensor-ID"`
}
