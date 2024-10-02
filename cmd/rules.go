package cmd

import (
	"gopkg.in/yaml.v3"
)

//go:generate ./gen_list.sh OBJS_NAME=rules OBJS_TYPE=[]huev1.Rule GET_OBJ_FUNC=GetRules()
//go:generate ./gen_show.sh OBJ_NAME=rule GET_OBJ_FUNC=GetRule

const rulesDefaultSortField = "ID"

var rulesDefaultFields = []string{
	"ID",
	"Name",
	"Status",
	"Created",
	"LastTriggered",
	"TimesTriggered",
}

var rulesHeaderTransform = newHeaderTransform(map[string]string{
	"TimesTriggered": "Times",
})

var rulesFieldTransform = newFieldTransform(map[string]fieldTransformFunc{
	"LastTriggered": func(v string) string {
		if v == "" {
			return "never"
		}
		return v
	},
})

func init() {
	_, err := parser.AddCommand("rules", "Manage rules", "", &rulesCmd{})
	if err != nil {
		panic(err)
	}
}

type rulesCmd struct {
	RulesList *rulesListCmd `command:"list" alias:"ls" description:"List rules"`
	RulesShow *rulesShowCmd `command:"show" description:"Display the specified rule"`
}

func (c *rulesShowCmd) PostProcessShowCmd(bytes []byte) ([]byte, error) {
	var data yaml.Node
	err := yaml.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	rule := yamlMap(data.Content[0])

	for _, action := range rule["actions"].Content {
		actionMap := yamlMap(action)
		annotateResourcePathWithYAMLComment(actionMap["address"])
	}

	for _, condition := range rule["conditions"].Content {
		conditionMap := yamlMap(condition)
		annotateResourcePathWithYAMLComment(conditionMap["address"])
	}

	return yaml.Marshal(&data)
}
