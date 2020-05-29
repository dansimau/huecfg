package cmd

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/mcuadros/go-lookup"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v3"
)

var stringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

// annotateResourcePathWithYAMLComment takes a YAML node with a string path to
// a resource on the bridge (e.g. "/groups/4/action") and adds a comment to
// the YAML node with the name of the resource it is pointing to. This allows
// us to output human-readable names of resources in the YAML output to make
// it easier to read an object's config.
func annotateResourcePathWithYAMLComment(node *yaml.Node) error {
	bridge := cmd.getHue()

	urlParts := strings.Split(node.Value, "/")
	if len(urlParts) < 3 {
		return fmt.Errorf("invalid resource path: %s", node.Value)
	}

	var resourceType string
	var resourceID string

	if strings.HasPrefix(node.Value, "/api") {
		// Deal with URL paths prefixed with the api/username, e.g.:
		// "/api/8p-fLSPhGBkpPWSp8NXt67LGiGSXavlS/sensors/99/state"
		resourceType = urlParts[3]
		resourceID = urlParts[4]
	} else {
		// URL paths without the /api prefix
		resourceType = urlParts[1]
		resourceID = urlParts[2]
	}

	var err error
	var obj interface{}

	switch resourceType {
	case "groups":
		obj, err = bridge.GetGroup(resourceID)
	case "lights":
		obj, err = bridge.GetLight(resourceID)
	case "resourcelinks":
		obj, err = bridge.GetResourceLink(resourceID)
	case "rules":
		obj, err = bridge.GetRule(resourceID)
	case "scenes":
		obj, err = bridge.GetScene(resourceID)
	case "schedules":
		obj, err = bridge.GetSchedule(resourceID)
	case "sensors":
		obj, err = bridge.GetSensor(resourceID)
	default:
		// Unknown or unsupported type
		return nil
	}

	// TODO: we should NOT error out if the resource is not found
	if err != nil {
		return err
	}

	name, err := lookupField(obj, "Name")
	if err != nil {
		return err
	}

	node.LineComment = name.String()

	return nil
}

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
