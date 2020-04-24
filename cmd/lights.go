package cmd

var lightsDefaultFields = []string{
	"ID",
	"Name",
	"Type",
	"ManufacturerName",
	"State.On",
}

const lightsDefaultSortField = "ID"

var lightsHeaderTransform = newHeaderTransform(map[string]string{
	"State.On": "State",
})

var lightsFieldTransform = newFieldTransform(map[string]fieldTransformFunc{
	"State.On": func(v string) string {
		r := "Off"
		if v == "true" {
			r = "On"
		}
		return r
	},
})

func init() {
	_, err := parser.AddCommand("lights", "Manage lights", "", &lightsCmd{})
	if err != nil {
		panic(err)
	}
}

type lightsCmd struct {
	LightsList *lightsListCmd `command:"list" alias:"ls" description:"List lights"`
}
