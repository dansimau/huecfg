package cmd

//go:generate ./gen_list.sh OBJS_NAME=sensors OBJS_TYPE=[]hue.Sensor GET_OBJ_FUNC=GetSensors()

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
	SensorsList *sensorsListCmd `command:"list" alias:"ls" description:"List sensors"`
}
