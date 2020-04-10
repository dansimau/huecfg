package cmd

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/dansimau/huecfg/pkg/hue"
)

func init() {
	_, err := parser.AddCommand("lights", "Manage lights", "", lights)
	if err != nil {
		panic(err)
	}
}

var lights = &lightsCmd{}

type lightsCmd struct {
	Host     string `short:"a" long:"host" description:"host address for Hue Bridge" required:"true"`
	Username string `short:"u" long:"username" description:"username from Hue Bridge registration"`

	LightsList *lightsListCmd `command:"list" alias:"ls" description:"List lights"`
}

func (c *lightsCmd) getHue() *hue.Hue {
	h := hue.NewConn(c.Host, c.Username)

	if len(cmd.Verbose) > 0 {
		h.API.Debug = true
	}

	return h
}

func lightsMapToSlice(m map[int]*hue.Light) (s []*hue.Light) {
	for _, light := range m {
		s = append(s, light)
	}
	return s
}

type mapVal struct {
	ID    int
	Value *reflect.Value
}

func mapValCmp(mapVals []*mapVal) func(i, j int) bool {
	return func(i, j int) bool {
		if mapVals[i].Value.Kind() != mapVals[j].Value.Kind() {
			return false
		}

		switch mapVals[i].Value.Kind() {
		case reflect.Int:
			return mapVals[i].Value.Int() < mapVals[j].Value.Int()
		case reflect.String:
			return mapVals[i].Value.String() < mapVals[j].Value.String()
		}

		return false
	}
}

func sortedIDsByField(m map[int]*hue.Light, field string) ([]int, error) {
	lights := lightsMapToSlice(m)

	mapVals := []*mapVal{}
	for _, obj := range lights {
		val, err := lookupField(obj, field)
		if err != nil {
			return nil, err
		}

		switch val.Kind() {
		case reflect.Int:
		case reflect.String:
			goto Pass
		default:
			return nil, fmt.Errorf("cannot sort by unknown field type: %v", val.Kind())
		}

	Pass:
		mapVals = append(mapVals, &mapVal{
			ID:    obj.ID,
			Value: val,
		})
	}

	sort.Slice(mapVals, mapValCmp(mapVals))

	sortedIDs := []int{}
	for _, v := range mapVals {
		sortedIDs = append(sortedIDs, v.ID)
	}

	return sortedIDs, nil
}

func sortLightsByField(m map[int]*hue.Light, field string) ([]*hue.Light, error) {
	sortedIDs, err := sortedIDsByField(m, field)
	if err != nil {
		return nil, err
	}

	sortedLights := []*hue.Light{}
	for _, ID := range sortedIDs {
		sortedLights = append(sortedLights, m[ID])
	}

	return sortedLights, nil
}
