package cmd

import "github.com/dansimau/huecfg/pkg/hue"

//go:generate ./gen_list.sh OBJS_NAME=users GET_OBJ_FUNC=GetConfig() OBJ_TRANSFORM_FUNC=configToUsersGenericSlice

var usersDefaultFields = []string{
	"APIKey",
	"Name",
	"Created",
	"LastUsed",
}

const usersDefaultSortField = "Name"

var (
	usersHeaderTransform headerTransform
	usersFieldTransform  fieldTransform
)

func init() {
	_, err := parser.AddCommand("users", "Manage users (API keys)", "", &usersListCmd{})
	if err != nil {
		panic(err)
	}
}

type user struct {
	APIKey   string
	Created  hue.AbsoluteTime
	LastUsed hue.AbsoluteTime
	Name     string
}

// capabilitiesToResourceUsageGenericSlice is customised for this particular
// cmd. We take a hue.Capabilities object and turn it into a slice of objects
// so we can reuse the existing list command codegen.
func configToUsersGenericSlice(c hue.Config) []interface{} {
	s := configToUsersSlice(c)
	var res = make([]interface{}, len(s))
	for i, obj := range s {
		res[i] = obj
	}
	return res
}

func configToUsersSlice(c hue.Config) []user {
	users := []user{}
	for key, v := range c.Whitelist {
		users = append(users, user{
			APIKey:   key,
			Name:     v.Name,
			Created:  v.CreateDate,
			LastUsed: v.LastUseDate,
		})
	}
	return users
}
