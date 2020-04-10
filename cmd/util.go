package cmd

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/mcuadros/go-lookup"
	"github.com/olekukonko/tablewriter"
)

func lookupField(s interface{}, path string) (*reflect.Value, error) {
	res, err := lookup.LookupStringI(s, path)
	if err != nil {
		if err == lookup.ErrKeyNotFound {
			return nil, fmt.Errorf("%w: %v", err, path)
		}
		return nil, err
	}

	return &res, nil
}

func printTable(rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(rows[0])
	table.SetAutoFormatHeaders(false)

	for _, row := range rows[1:] {
		table.Append(row)
	}

	table.Render()
}

func reflectValueToString(v *reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Bool:
		if v.Bool() {
			return "true"
		} else {
			return "false"
		}
	case reflect.Int:
		return strconv.Itoa(int(v.Int()))
	default:
		return fmt.Sprintf("<%s>", v.Type().String())
	}
}
