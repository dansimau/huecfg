package cmd

import "gopkg.in/yaml.v3"

//go:generate ./gen_list.sh OBJS_NAME=groups OBJS_TYPE=[]hue.Group GET_OBJ_FUNC=GetGroups()
//go:generate ./gen_show.sh OBJ_NAME=group GET_OBJ_FUNC=GetGroup

const groupsDefaultSortField = "ID"

var (
	groupsDefaultFields = []string{"ID", "Name", "Type"}

	groupsFieldTransform  fieldTransform
	groupsHeaderTransform headerTransform
)

func init() {
	_, err := parser.AddCommand("groups", "Manage groups", "", &groupsCmd{})
	if err != nil {
		panic(err)
	}
}

type groupsCmd struct {
	GroupsList *groupsListCmd `command:"list" alias:"ls" description:"List groups"`
	GroupsShow *groupsShowCmd `command:"show" description:"Gets the group attributes for a given group"`
}

func (c *groupsShowCmd) PostProcessShowCmd(bytes []byte) ([]byte, error) {
	bridge := cmd.getHue()

	var data yaml.Node
	err := yaml.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	group := yamlMap(data.Content[0])

	// Add light names as YAML comments
	for _, lightID := range group["lights"].Content {
		// TODO: add lights cache and collapse into single API call?
		light, err := bridge.GetLight(lightID.Value)
		if err != nil {
			return nil, err
		}

		lightID.LineComment = light.Name
	}

	return yaml.Marshal(&data)
}
