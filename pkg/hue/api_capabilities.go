package hue

import (
	"fmt"
	"io/ioutil"
)

// GetCapabilities returns information about resources remaining for different
// object types, as well as capabilities of features.
func (api *API) GetCapabilities() ([]byte, error) {
	resp, err := api.httpGet(fmt.Sprintf("/api/%s/capabilities", api.username()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
