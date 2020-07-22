// Code generated by go generate; DO NOT EDIT.
package cmd

import (
	"strings"

	"github.com/iancoleman/strcase"
)

type freeListCmd struct {
	Fields  string `long:"fields" description:"List of fields to include"`
	Reverse bool   `long:"reverse" description:"Reverse sort order"`
	Sort    string `long:"sort" description:"Field to sort by"`
}

func (c *freeListCmd) Execute(args []string) error {
	if err := errorOnUnknownArgs(args); err != nil {
		return err
	}

	bridge := cmd.getHue()

	free, err := bridge.GetCapabilities()
	if err != nil {
		return err
	}

	fields := freeDefaultFields
	if c.Fields != "" {
		fields = []string{}
		for _, fieldName := range strings.Split(c.Fields, ",") {
			fields = append(fields, freeHeaderTransform.TransformInput(fieldName))
		}
	}

	sortField := freeDefaultSortField
	if c.Sort != "" {
		sortField = strcase.ToCamel(freeHeaderTransform.TransformInput(c.Sort))
	}

	sortedfree, err := sortByField(capabilitiesToResourceUsageGenericSlice(free), sortField, c.Reverse)
	if err != nil {
		return err
	}

	rows := [][]string{}

	headerRow := []string{}
	for _, fieldName := range fields {
		headerRow = append(headerRow, freeHeaderTransform.TransformOutput(fieldName))
	}
	rows = append(rows, headerRow)

	for _, light := range sortedfree {
		row := []string{}

		for _, field := range fields {
			v, err := lookupField(light, field)
			if err != nil {
				return err
			}

			row = append(row, freeFieldTransform.TransformOutput(field, reflectValueToString(v)))
		}

		rows = append(rows, row)
	}

	printTable(rows)

	return nil
}
