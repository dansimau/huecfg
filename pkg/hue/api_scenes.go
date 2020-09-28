package hue

import (
	"fmt"
)

// CreateScene creates a new scene
func (api *API) CreateScene(data interface{}) ([]byte, error) {
	return api.post(fmt.Sprintf("/api/%s/scenes", api.username()), data)
}

// DeleteScene deletes a scene from the bridge.
func (api *API) DeleteScene(id string) ([]byte, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	return api.delete(fmt.Sprintf("/api/%s/scenes/%s", api.username(), id))
}

// GetScenes gets a list of all scenes currently stored in the bridge. Scenes
// are represented by a scene id, a name and a list of lights which are part
// of the scene. The name resource can contain a “friendly name” or can contain
// a unique code.  Scenes are stored in the bridge.  This means that scene
// light state settings can easily be retrieved by developers (using ADD link)
// and shown in their respective UIs. Cached scenes (scenes stored with PUT)
// will be deprecated in the future.
//
// Additionally, bridge scenes should not be confused with the preset scenes
// stored in the Android and iOS Hue apps. In the apps these scenes are stored
// internally. Once activated they may then appear as a bridge scene.
func (api *API) GetScenes() ([]byte, error) {
	return api.get(fmt.Sprintf("/api/%s/scenes", api.username()))
}

// GetScene get the attributes of a given scene. Please note that lightstates
// are displayed when an individual scene is retrieved (but not for all
// scenes).
func (api *API) GetScene(id string) ([]byte, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	return api.get(fmt.Sprintf("/api/%s/scenes/%s", api.username(), id))
}
