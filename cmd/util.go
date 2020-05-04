package cmd

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/mcuadros/go-lookup"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v3"
)

var stringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

func lookupField(s interface{}, path string) (reflect.Value, error) {
	res, err := lookup.LookupStringI(s, path)
	if err != nil {
		if err == lookup.ErrKeyNotFound {
			return reflect.Value{}, fmt.Errorf("%w: %v", err, path)
		}
		return reflect.Value{}, err
	}

	return res, nil
}

func printTable(rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader(rows[0])

	for _, row := range rows[1:] {
		table.Append(row)
	}

	table.Render()
}

func reflectValueToString(v reflect.Value) string {
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
		if v.Type().Implements(stringerType) {
			return v.MethodByName("String").Call([]reflect.Value{})[0].String()
		}
		return fmt.Sprintf("<%s>", v.Type().String())
	}
}

// yamlMap takes a yaml.Node that is a map and returns a Go map with the values
// of type yaml.Node.
func yamlMap(node *yaml.Node) map[string]*yaml.Node {
	m := make(map[string]*yaml.Node, len(node.Content)/2)

	var key string
	for i, node := range node.Content {
		switch i % 2 {
		case 0:
			key = node.Value
		case 1:
			m[key] = node
		}
	}

	return m
}
