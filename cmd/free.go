package cmd

import "github.com/dansimau/huecfg/pkg/huev1"

//go:generate ./gen_list.sh OBJS_NAME=free GET_OBJ_FUNC=GetCapabilities() OBJ_TRANSFORM_FUNC=capabilitiesToResourceUsageGenericSlice

var freeDefaultFields = []string{
	"Resource",
	"Total",
	"Used",
	"Free",
}

const freeDefaultSortField = "Resource"

var (
	freeHeaderTransform headerTransform
	freeFieldTransform  fieldTransform
)

func init() {
	_, err := parser.AddCommand("free", "Show free/used memory on the Hue Bridge", "", &freeListCmd{})
	if err != nil {
		panic(err)
	}
}

type resourceUsage struct {
	Resource string
	Free     int
	Total    int
	Used     int
}

// capabilitiesToResourceUsageGenericSlice is customised for this particular
// cmd. We take a huev1.Capabilities object and turn it into a slice of objects
// so we can reuse the existing list command codegen.
func capabilitiesToResourceUsageGenericSlice(c huev1.Capabilities) []interface{} {
	s := capabilitiesToResourceUsageSlice(c)
	res := make([]interface{}, len(s))
	for i, obj := range s {
		res[i] = obj
	}
	return res
}

func capabilitiesToResourceUsageSlice(c huev1.Capabilities) []resourceUsage {
	res := []resourceUsage{
		{
			Resource: "Groups",
			Total:    c.Groups.Total,
			Free:     c.Groups.Available,
			Used:     c.Groups.Total - c.Groups.Available,
		},
		{
			Resource: "Lights",
			Total:    c.Lights.Total,
			Free:     c.Lights.Available,
			Used:     c.Lights.Total - c.Lights.Available,
		},
		{
			Resource: "ResourceLinks",
			Total:    c.ResourceLinks.Total,
			Free:     c.ResourceLinks.Available,
			Used:     c.ResourceLinks.Total - c.ResourceLinks.Available,
		},
		{
			Resource: "Rules",
			Total:    c.Rules.Total,
			Free:     c.Rules.Available,
			Used:     c.Rules.Total - c.Rules.Available,
		},
		{
			Resource: "Scenes",
			Total:    c.Scenes.Total,
			Free:     c.Scenes.Available,
			Used:     c.Scenes.Total - c.Scenes.Available,
		},
		{
			Resource: "Schedules",
			Total:    c.Schedules.Total,
			Free:     c.Schedules.Available,
			Used:     c.Schedules.Total - c.Schedules.Available,
		},
		{
			Resource: "Sensors",
			Total:    c.Sensors.Total,
			Free:     c.Sensors.Available,
			Used:     c.Sensors.Total - c.Sensors.Available,
		},
		{
			Resource: "Streaming",
			Total:    c.Streaming.Total,
			Free:     c.Streaming.Available,
			Used:     c.Streaming.Total - c.Streaming.Available,
		},
	}
	return res
}
