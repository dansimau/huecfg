package cmd

import (
	"fmt"
	"os"

	"github.com/iancoleman/strcase"
)

var defaultFields = []string{
	"ID",
	"Name",
	"Type",
	"State.On",
}

const defaultSortField = "ID"

// huecfg lights ls ...
type lightsListCmd struct {
	Fields []string `long:"fields" description:"List of fields to include"`
	Sort   string   `long:"sort" description:"Field to sort by"`
}

func (c *lightsListCmd) Execute(args []string) error {
	bridge := lights.getHue()

	lights, err := bridge.GetLights()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fields := defaultFields
	if c.Fields != nil {
		fields = c.Fields
	}

	sortField := defaultSortField
	if c.Sort != "" {
		sortField = strcase.ToCamel(c.Sort)
	}

	sortedLights, err := sortLightsByField(lights, sortField)
	if err != nil {
		return err
	}

	values := [][]string{}
	for _, light := range sortedLights {
		row := []string{}

		for _, field := range fields {
			v, err := lookupField(light, field)
			if err != nil {
				return err
			}

			row = append(row, reflectValueToString(v))
		}

		values = append(values, row)
	}

	printTable(values)

	return nil
}
