package cmd

import "github.com/dansimau/huecfg/pkg/huev1"

//go:generate ./gen_list.sh OBJS_NAME=users GET_OBJ_FUNC=GetConfig() OBJ_TRANSFORM_FUNC=configToUsersGenericSlice

var usersDefaultFields = []string{
	"Username",
	"DeviceType",
	"Created",
	"LastUsed",
}

const usersDefaultSortField = "DeviceType"

var (
	usersHeaderTransform headerTransform
	usersFieldTransform  fieldTransform
)

func init() {
	_, err := parser.AddCommand("users", "Manage users (API keys)", "", &usersCmd{})
	if err != nil {
		panic(err)
	}
}

type usersCmd struct {
	UsersList *usersListCmd `command:"list" alias:"ls" description:"List users/API keys"`
}

type user struct {
	Username   string
	DeviceType string
	Created    huev1.AbsoluteTime
	LastUsed   huev1.AbsoluteTime
}

// capabilitiesToResourceUsageGenericSlice is customised for this particular
// cmd. We take a huev1.Capabilities object and turn it into a slice of objects
// so we can reuse the existing list command codegen.
func configToUsersGenericSlice(c huev1.Config) []interface{} {
	s := configToUsersSlice(c)
	res := make([]interface{}, len(s))
	for i, obj := range s {
		res[i] = obj
	}
	return res
}

func configToUsersSlice(c huev1.Config) []user {
	users := []user{}
	for key, v := range c.Whitelist {
		users = append(users, user{
			Username:   key,
			DeviceType: v.Name,
			Created:    v.CreateDate,
			LastUsed:   v.LastUseDate,
		})
	}
	return users
}

func configToUsersMap(c huev1.Config) map[string]user {
	users := map[string]user{}
	for key, v := range c.Whitelist {
		users[key] = user{
			Username:   key,
			DeviceType: v.Name,
			Created:    v.CreateDate,
			LastUsed:   v.LastUseDate,
		}
	}
	return users
}
