package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/mcuadros/go-lookup"
	"github.com/olekukonko/tablewriter"
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

func setFieldValue(obj map[string]interface{}, path string, value string) error {
	pathComponents := strings.Split(path, ".")

	node := obj

	for i := 0; i < len(pathComponents)-1; i++ {
		k := pathComponents[i]

		val, ok := node[k]
		if !ok {
			val = map[string]interface{}{}
			node[k] = val
		}

		node, ok = val.(map[string]interface{})
		if !ok {
			return fmt.Errorf("cannot set %v, would override value", strings.Join(pathComponents[0:i], "."))
		}
	}

	// Set the value
	node[pathComponents[len(pathComponents)-1]] = value

	return nil
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

func mustStrToBool(s string) bool {
	switch strings.ToLower(s) {
	case "true", "yes", "1":
		return true
	case "false", "no", "0":
		return false
	}
	panic(fmt.Sprintf("cannot convert \"%s\" to bool", s))
}

func boolToOnOff(b bool) string {
	if b {
		return "On"
	} else {
		return "Off"
	}
}

func boolToYesNo(b bool) string {
	if b {
		return "Yes"
	} else {
		return "No"
	}
}

// userInputToJSON takes user input from the command line, marshalls it to JSON
// and then returns it. If "-" is specified then it reads the data from stdin.
func userInputToJSON(data string) (map[string]interface{}, error) {
	if data == "-" {
		buf := &bytes.Buffer{}
		_, err := io.Copy(buf, os.Stdin)
		if err != nil {
			return nil, err
		}
		data = buf.String()
	}

	v := map[string]interface{}{}
	if err := json.Unmarshal([]byte(data), &v); err != nil {
		return nil, err
	}

	return v, nil
}
