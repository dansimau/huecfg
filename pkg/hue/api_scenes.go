package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// CreateScene creates a new scene
func (api *API) CreateScene(data interface{}) ([]byte, error) {
	postJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpPost(fmt.Sprintf("/api/%s/scenes", api.username()), bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// DeleteScene deletes a scene from the bridge.
func (api *API) DeleteScene(id string) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	resp, err := api.httpDelete(fmt.Sprintf("/api/%s/scenes/%s", api.username(), id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
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
	resp, err := api.httpGet(fmt.Sprintf("/api/%s/scenes", api.username()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// GetScene get the attributes of a given scene. Please note that lightstates
// are displayed when an individual scene is retrieved (but not for all
// scenes).
func (api *API) GetScene(id string) ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}

	resp, err := api.httpGet(fmt.Sprintf("/api/%s/scenes/%s", api.username(), id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
