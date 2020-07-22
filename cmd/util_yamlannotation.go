package cmd

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

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

func annotateUsername(node *yaml.Node) error {
	bridge := cmd.getHue()

	config, err := bridge.GetConfig()
	if err != nil {
		return err
	}

	users := configToUsersMap(config)

	user, ok := users[node.Value]
	if !ok {
		return errors.New("username not found")
	}

	node.LineComment = user.DeviceType

	return nil

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
