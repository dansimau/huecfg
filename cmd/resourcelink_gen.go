package cmd

//go:generate ./gen_list.sh OBJS_NAME=resourcelink OBJS_TYPE=[]hue.ResourceLink GET_OBJ_FUNC=GetResourceLinks()

const resourcelinkDefaultSortField = "ID"

var (
	resourcelinkDefaultFields = []string{"ID", "Name", "Description", "ClassID"}

	resourcelinkFieldTransform  fieldTransform
	resourcelinkHeaderTransform headerTransform
)

func init() {
	_, err := parser.AddCommand("resourcelink", "Manage resourcelink", "", &resourcelinkCmd{})
	if err != nil {
		panic(err)
	}
}

type resourcelinkCmd struct {
	ResourcelinkList *resourcelinkListCmd `command:"list" alias:"ls" description:"List resourcelink"`
}
