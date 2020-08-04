package hue

import (
	"fmt"
)

// GetCapabilities returns information about resources remaining for different
// object types, as well as capabilities of features.
func (api *API) GetCapabilities() ([]byte, error) {
	return api.get(fmt.Sprintf("/api/%s/capabilities", api.username()))
}
