// Code generated by go generate; DO NOT EDIT.
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
)

type usersListCmd struct {
	Fields  string `long:"fields" description:"List of fields to include"`
	Reverse bool   `long:"reverse" description:"Reverse sort order"`
	Sort    string `long:"sort" description:"Field to sort by"`
}

func (c *usersListCmd) Execute(args []string) error {
	bridge := cmd.getHue()

	users, err := bridge.GetConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	fields := usersDefaultFields
	if c.Fields != "" {
		fields = []string{}
		for _, fieldName := range strings.Split(c.Fields, ",") {
			fields = append(fields, usersHeaderTransform.TransformInput(fieldName))
		}
	}

	sortField := usersDefaultSortField
	if c.Sort != "" {
		sortField = strcase.ToCamel(usersHeaderTransform.TransformInput(c.Sort))
	}

	sortedusers, err := sortByField(configToUsersGenericSlice(users), sortField, c.Reverse)
	if err != nil {
		return err
	}

	rows := [][]string{}

	headerRow := []string{}
	for _, fieldName := range fields {
		headerRow = append(headerRow, usersHeaderTransform.TransformOutput(fieldName))
	}
	rows = append(rows, headerRow)

	for _, light := range sortedusers {
		row := []string{}

		for _, field := range fields {
			v, err := lookupField(light, field)
			if err != nil {
				return err
			}

			row = append(row, usersFieldTransform.TransformOutput(field, reflectValueToString(v)))
		}

		rows = append(rows, row)
	}

	printTable(rows)

	return nil
}
