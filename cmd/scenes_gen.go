package cmd

//go:generate ./gen_list.sh OBJS_NAME=scenes OBJS_TYPE=[]hue.Scene GET_OBJ_FUNC=GetScenes()

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
}
