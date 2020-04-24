package cmd

//go:generate ./gen_list.sh OBJS_NAME=groups OBJS_TYPE=[]hue.Group GET_OBJ_FUNC=GetGroups()

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
}
