package cmd

import "gopkg.in/yaml.v3"

//go:generate ./gen_list.sh OBJS_NAME=schedules OBJS_TYPE=[]huev1.Schedule GET_OBJ_FUNC=GetSchedules()
//go:generate ./gen_show.sh OBJ_NAME=schedule GET_OBJ_FUNC=GetSchedule

const schedulesDefaultSortField = "ID"

var (
	schedulesDefaultFields = []string{"ID", "Name"}

	schedulesFieldTransform  fieldTransform
	schedulesHeaderTransform headerTransform
)

func init() {
	_, err := parser.AddCommand("schedules", "Manage schedules", "", &schedulesCmd{})
	if err != nil {
		panic(err)
	}
}

type schedulesCmd struct {
	SchedulesList *schedulesListCmd `command:"list" alias:"ls" description:"List schedules"`
	SchedulesShow *schedulesShowCmd `command:"show" description:"Get all attributes for a schedule"`
}

func (c *schedulesShowCmd) PostProcessShowCmd(bytes []byte) ([]byte, error) {
	var data yaml.Node
	err := yaml.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	schedule := yamlMap(data.Content[0])
	scheduleCommand := yamlMap(schedule["command"])
	annotateResourcePathWithYAMLComment(scheduleCommand["address"])

	return yaml.Marshal(&data)
}
