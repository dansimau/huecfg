package cmd

//go:generate ./gen_list.sh OBJS_NAME=schedules OBJS_TYPE=[]hue.Schedule GET_OBJ_FUNC=GetSchedules()

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
}
