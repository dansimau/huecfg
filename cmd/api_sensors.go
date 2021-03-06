package cmd

// huecfg api sensors ...
type apiSensorsCmd struct {
	Create *apiSensorsCreateCmd `command:"create" description:"Create a new CLIP software sensor"`
	Delete *apiSensorsDeleteCmd `command:"delete" description:"Delete a device from the bridge"`
	Get    *apiSensorsGetCmd    `command:"get" description:"Fetch the specified sensor by ID"`
	GetAll *apiSensorsGetAllCmd `command:"get-all" description:"Fetch all sensors at once"`
}

// huecfg api sensors create
//go:generate ./gen_api_write.sh ID=sensors_create TYPE=apiSensorsCreateCmd DATA=c.Data FUNC_CALL=bridge.CreateSensor(data)
type apiSensorsCreateCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`
}

// huecfg api sensors delete ...
//go:generate ./gen_api_read.sh ID=sensors_delete TYPE=apiSensorsDeleteCmd FUNC_CALL=bridge.DeleteSensor(c.Arguments.ID)
type apiSensorsDeleteCmd struct {
	Arguments struct {
		ID string `description:"ID of the sensor to delete."`
	} `positional-args:"true" required:"true" positional-arg-name:"sensor-ID"`
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
