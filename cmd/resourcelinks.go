package cmd

//go:generate ./gen_list.sh OBJS_NAME=resourcelinks OBJS_TYPE=[]hue.ResourceLink GET_OBJ_FUNC=GetResourceLinks()

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
}
