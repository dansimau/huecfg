package cmd

type apiV1SensorsCmd struct {
	Create *apiSensorsCreateCmd `command:"create" description:"Create a new CLIP software sensor"`
	Delete *apiSensorsDeleteCmd `command:"delete" description:"Delete a device from the bridge"`
	Get    *apiSensorsGetCmd    `command:"get" description:"Fetch the specified sensor by ID"`
	GetAll *apiSensorsGetAllCmd `command:"get-all" description:"Fetch all sensors at once"`
}

//go:generate ./gen_api_write.sh ID=sensors_create TYPE=apiV1SensorsCreateCmd DATA=c.Data FUNC_CALL=bridge.CreateSensor(data)
type apiSensorsCreateCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`
}

//go:generate ./gen_api_read.sh ID=sensors_delete TYPE=apiV1SensorsDeleteCmd FUNC_CALL=bridge.DeleteSensor(c.Arguments.ID)
type apiSensorsDeleteCmd struct {
	Arguments struct {
		ID string `description:"ID of the sensor to delete."`
	} `positional-args:"true" required:"true" positional-arg-name:"sensor-ID"`
}

//go:generate ./gen_api_read.sh ID=sensors_get_all TYPE=apiV1SensorsGetAllCmd FUNC_CALL=bridge.GetSensors()
type apiSensorsGetAllCmd struct{}

//go:generate ./gen_api_read.sh ID=sensors_get TYPE=apiV1SensorsGetCmd FUNC_CALL=bridge.GetSensor(c.Arguments.ID)
type apiSensorsGetCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"sensor-ID"`
}
