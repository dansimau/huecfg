package hue

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// GetGroups gets a list of all groups that have been added to the bridge. A
// group is a list of lights that can be created, modified and deleted by a
// user.
func (api *API) GetGroups() ([]byte, error) {
	resp, err := api.httpGet(fmt.Sprintf("/api/%s/groups", api.username()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// GetGroup gets the group attributes, e.g. name, light membership and last
// command for a given group.
func (api *API) GetGroup(ID int) ([]byte, error) {
	// if ID == "" {
	// 	return nil, errors.New("ID cannot be empty")
	// }

	strID := strconv.Itoa(ID)

	resp, err := api.httpGet(fmt.Sprintf("/api/%s/groups/%s", api.username(), strID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
