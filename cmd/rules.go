package cmd

//go:generate ./gen_list.sh OBJS_NAME=rules OBJS_TYPE=[]hue.Rule GET_OBJ_FUNC=GetRules()

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
}
