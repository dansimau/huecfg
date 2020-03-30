package hue

import (
	"errors"
	"fmt"
	"io/ioutil"
)

// GetRules gets a list of all rules that are in the bridge.
func (api *API) GetRules() ([]byte, error) {
	resp, err := api.httpGet(fmt.Sprintf("/api/%s/rules", api.username()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// GetRule returns a rule object with id matching <id> or an error if <id> is
// not available.
func (api *API) GetRule(ID string) ([]byte, error) {
	if ID == "" {
		return nil, errors.New("ID cannot be empty")
	}

	resp, err := api.httpGet(fmt.Sprintf("/api/%s/rules/%s", api.username(), ID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
