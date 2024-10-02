package cmd

import "gopkg.in/yaml.v3"

//go:generate ./gen_list.sh OBJS_NAME=scenes OBJS_TYPE=[]huev1.Scene GET_OBJ_FUNC=GetScenes()
//go:generate ./gen_show.sh OBJ_NAME=scene GET_OBJ_FUNC=GetScene

const scenesDefaultSortField = "ID"

var (
	scenesDefaultFields = []string{"ID", "Name"}

	scenesFieldTransform  fieldTransform
	scenesHeaderTransform headerTransform
)

func init() {
	_, err := parser.AddCommand("scenes", "Manage scenes", "", &scenesCmd{})
	if err != nil {
		panic(err)
	}
}

type scenesCmd struct {
	ScenesList *scenesListCmd `command:"list" alias:"ls" description:"List scenes"`
	ScenesShow *scenesShowCmd `command:"show" description:"Gets the attributes of a given scene"`
}

func (c *scenesShowCmd) PostProcessShowCmd(bytes []byte) ([]byte, error) {
	bridge := cmd.getHue()

	var data yaml.Node
	err := yaml.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	scene := yamlMap(data.Content[0])

	// Add light names as YAML comments
	for _, lightID := range scene["lights"].Content {
		// TODO: add lights cache and collapse into single API call?
		light, err := bridge.GetLight(lightID.Value)
		if err != nil {
			// TODO: we should NOT error out if the resource is not found
			return nil, err
		}

		lightID.LineComment = light.Name
	}

	return yaml.Marshal(&data)
}
