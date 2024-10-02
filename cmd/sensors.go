package cmd

//go:generate ./gen_list.sh OBJS_NAME=sensors OBJS_TYPE=[]huev1.Sensor GET_OBJ_FUNC=GetSensors()
//go:generate ./gen_show.sh OBJ_NAME=sensor GET_OBJ_FUNC=GetSensor

const sensorsDefaultSortField = "ID"

var sensorsDefaultFields = []string{
	"ID",
	"Name",
	"Type",
	"ManufacturerName",
	"State.LastUpdated",
}

var sensorsHeaderTransform = newHeaderTransform(map[string]string{
	"State.LastUpdated": "LastEvent",
})

var sensorsFieldTransform fieldTransform

func init() {
	_, err := parser.AddCommand("sensors", "Manage sensors", "", &sensorsCmd{})
	if err != nil {
		panic(err)
	}
}

type sensorsCmd struct {
	SensorsList      *sensorsListCmd      `command:"list" alias:"ls" description:"List sensors"`
	SensorsSetAttr   *sensorsSetAttrCmd   `command:"set-attr" description:"Set attributes of a sensor"`
	SensorsSetConfig *sensorsSetConfigCmd `command:"set-config" description:"Set config of a sensor"`
	SensorsSetState  *sensorsSetStateCmd  `command:"set-state" description:"Set state of a sensor"`
	SensorsShow      *sensorsShowCmd      `command:"show" description:"Gets the sensor from the bridge with the given ID"`
}
