package cmd

//go:generate ./gen_list.sh OBJS_NAME=scenes OBJS_TYPE=[]hue.Scene GET_OBJ_FUNC=GetScenes()
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
