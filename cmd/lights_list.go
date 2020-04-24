package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/dansimau/huecfg/pkg/hue"
	"github.com/iancoleman/strcase"
)

func lightsToGenericSlice(s []hue.Light) []interface{} {
	var res = make([]interface{}, len(s))
	for i, obj := range s {
		res[i] = obj
	}
	return res
}

// huecfg lights ls ...
type lightsListCmd struct {
	Fields string `long:"fields" description:"List of fields to include"`
	Sort   string `long:"sort" description:"Field to sort by"`
}

func (c *lightsListCmd) Execute(args []string) error {
	bridge := cmd.getHue()

	lights, err := bridge.GetLights()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fields := lightsDefaultFields
	if c.Fields != "" {
		fields = []string{}
		for _, fieldName := range strings.Split(c.Fields, ",") {
			fields = append(fields, lightsHeaderTransform.TransformInput(fieldName))
		}
	}

	sortField := lightsDefaultSortField
	if c.Sort != "" {
		sortField = strcase.ToCamel(c.Sort)
	}

	sortedLights, err := sortByField(lightsToGenericSlice(lights), sortField)
	if err != nil {
		return err
	}

	rows := [][]string{}

	headerRow := []string{}
	for _, fieldName := range fields {
		headerRow = append(headerRow, lightsHeaderTransform.TransformOutput(fieldName))
	}
	rows = append(rows, headerRow)

	for _, light := range sortedLights {
		row := []string{}

		for _, field := range fields {
			v, err := lookupField(light, field)
			if err != nil {
				return err
			}

			row = append(row, lightsFieldTransform.TransformOutput(field, reflectValueToString(v)))
		}

		rows = append(rows, row)
	}

	printTable(rows)

	return nil
}
