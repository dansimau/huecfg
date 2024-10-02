package cmd

//go:generate ./gen_list.sh OBJS_NAME=lights OBJS_TYPE=[]huev1.Light GET_OBJ_FUNC=GetLights()
//go:generate ./gen_show.sh OBJ_NAME=light GET_OBJ_FUNC=GetLight

var lightsDefaultFields = []string{
	"ID",
	"Name",
	"Type",
	"ManufacturerName",
	"State.Reachable",
	"State.On",
}

const lightsDefaultSortField = "ID"

var lightsHeaderTransform = newHeaderTransform(map[string]string{
	"State.On":        "State",
	"State.Reachable": "Reachable",
})

var lightsFieldTransform = newFieldTransform(map[string]fieldTransformFunc{
	"State.On": func(v string) string {
		return boolToOnOff(mustStrToBool(v))
	},
	"State.Reachable": func(v string) string {
		return boolToYesNo(mustStrToBool(v))
	},
})

func init() {
	_, err := parser.AddCommand("lights", "Manage lights", "", &lightsCmd{})
	if err != nil {
		panic(err)
	}
}

type lightsCmd struct {
	LightsList     *lightsListCmd     `command:"list" alias:"ls" description:"List lights"`
	LightsSetAttr  *lightsSetAttrCmd  `command:"set-attr" description:"Set attributes of a light"`
	LightsSetState *lightsSetStateCmd `command:"set-state" description:"Set state of a light"`
	LightsShow     *lightsShowCmd     `command:"show" description:"Gets the attributes and state of a given light"`
}
