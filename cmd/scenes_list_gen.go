// Code generated by go generate; DO NOT EDIT.
package cmd

import (
	"strings"

	"github.com/dansimau/huecfg/pkg/hue"
	"github.com/iancoleman/strcase"
)

func scenesToGenericSlice(s []hue.Scene) []interface{} {
	var res = make([]interface{}, len(s))
	for i, obj := range s {
		res[i] = obj
	}
	return res
}

type scenesListCmd struct {
	Fields  string `long:"fields" description:"List of fields to include"`
	Reverse bool   `long:"reverse" description:"Reverse sort order"`
	Sort    string `long:"sort" description:"Field to sort by"`
}

func (c *scenesListCmd) Execute(args []string) error {
	if err := errorOnUnknownArgs(args); err != nil {
		return err
	}

	bridge := cmd.getHue()

	scenes, err := bridge.GetScenes()
	if err != nil {
		return err
	}

	fields := scenesDefaultFields
	if c.Fields != "" {
		fields = []string{}
		for _, fieldName := range strings.Split(c.Fields, ",") {
			fields = append(fields, scenesHeaderTransform.TransformInput(fieldName))
		}
	}

	sortField := scenesDefaultSortField
	if c.Sort != "" {
		sortField = strcase.ToCamel(scenesHeaderTransform.TransformInput(c.Sort))
	}

	sortedscenes, err := sortByField(scenesToGenericSlice(scenes), sortField, c.Reverse)
	if err != nil {
		return err
	}

	rows := [][]string{}

	headerRow := []string{}
	for _, fieldName := range fields {
		headerRow = append(headerRow, scenesHeaderTransform.TransformOutput(fieldName))
	}
	rows = append(rows, headerRow)

	for _, light := range sortedscenes {
		row := []string{}

		for _, field := range fields {
			v, err := lookupField(light, field)
			if err != nil {
				return err
			}

			row = append(row, scenesFieldTransform.TransformOutput(field, reflectValueToString(v)))
		}

		rows = append(rows, row)
	}

	printTable(rows)

	return nil
}
