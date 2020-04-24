package hue

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// GetResourceLinks gets a list of all resourcelinks that have been added to the
// bridge.
func (api *API) GetResourceLinks() ([]byte, error) {
	resp, err := api.httpGet(fmt.Sprintf("/api/%s/resourcelinks", api.username()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// GetResourceLink gets all attributes for a resourcelink.
func (api *API) GetResourceLink(ID int) ([]byte, error) {
	// if ID == "" {
	// 	return nil, errors.New("ID cannot be empty")
	// }

	strID := strconv.Itoa(ID)

	resp, err := api.httpGet(fmt.Sprintf("/api/%s/resourcelinks/%s", api.username(), strID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
