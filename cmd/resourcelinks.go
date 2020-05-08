package cmd

import (
	"gopkg.in/yaml.v3"
)

//go:generate ./gen_list.sh OBJS_NAME=resourcelinks OBJS_TYPE=[]hue.ResourceLink GET_OBJ_FUNC=GetResourceLinks()
//go:generate ./gen_show.sh OBJ_NAME=resourcelink GET_OBJ_FUNC=GetResourceLink

const resourcelinksDefaultSortField = "ID"

var (
	resourcelinksDefaultFields = []string{"ID", "Name", "Description", "ClassID"}

	resourcelinksFieldTransform  fieldTransform
	resourcelinksHeaderTransform headerTransform
)

func init() {
	_, err := parser.AddCommand("resourcelinks", "Manage resourcelinks", "", &resourcelinksCmd{})
	if err != nil {
		panic(err)
	}
}

type resourcelinksCmd struct {
	ResourcelinksList *resourcelinksListCmd `command:"list" alias:"ls" description:"List resourcelinks"`
	ResourcelinksShow *resourcelinksShowCmd `command:"show" description:"Display the specified resourcelink object"`
}

func (c *resourcelinksShowCmd) PostProcessShowCmd(bytes []byte) ([]byte, error) {
	var data yaml.Node
	err := yaml.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	resourcelink := yamlMap(data.Content[0])

	// Annotate links with YAML comments of the resource names
	for _, resourcePath := range resourcelink["links"].Content {
		annotateResourcePathWithYAMLComment(resourcePath)
	}

	return yaml.Marshal(&data)
}
