package huev1

import (
	"encoding/json"
)

type Scene struct {
	ID string

	// App specific data linked to the scene.  Each individual application
	// should take responsibility for the data written in this field.
	AppData struct {
		// App specific data. Free format string.
		Data string

		// App specific version of the data field. App should take versioning
		// into account when parsing the data string.
		Version int
	}

	// Group ID that a scene is linked to.
	Group string // TODO: convert to int?

	// UTC time the scene has been created or has been updated
	LastUpdated *AbsoluteTime

	// Light IDs which are in the scene. This array can empty. As of 1.11
	// it must contain at least 1 element. If an invalid lights resource is
	// given, error 7 is returned and the scene is not created.  When writing,
	// lightstate of all lights in list will be overwritten with current light
	// state. As of 1.15 when writing, lightstate of lights which are not yet
	// in list will be created with current light state. The array is
	// informational for GroupScene, it is generated automatically from the
	// lights in the linked group.
	Lights []string

	// Indicates that the scene is locked by a rule or a schedule and cannot be
	// deleted until all resources requiring or that reference the scene are
	// deleted.
	Locked bool

	// Human readable name of the scene. Is set to <id> if omitted on creation.
	Name string

	// Whitelist user that created or modified the content of the scene. Note
	// that changing name does not change the owner.
	Owner string

	// Reserved for future use.
	Picture string

	// Indicates whether the scene can be automatically deleted by the bridge.
	Recycle bool

	Type string // TODO: convert to string enum

	// Version of scene document
	// 1 – Scene created via PUT, lightstates will be empty.
	// 2 – Scene created via POST lightstates available.
	Version int
}

func (h *Hue) GetScenes() ([]Scene, error) {
	respBytes, err := h.API.GetScenes()
	if err != nil {
		return nil, err
	}

	var objs map[string]Scene
	if err := json.Unmarshal(respBytes, &objs); err != nil {
		return nil, err
	}

	if hueErr := parseAsHueError(respBytes); hueErr != nil {
		return nil, hueErr
	}

	res := []Scene{}
	for ID, obj := range objs {
		obj.ID = ID
		res = append(res, obj)
	}

	return res, nil
}

func (h *Hue) GetScene(ID string) (Scene, error) {
	respBytes, err := h.API.GetScene(ID)
	if err != nil {
		return Scene{}, err
	}

	if hueErr := parseAsHueError(respBytes); hueErr != nil {
		return Scene{}, hueErr
	}

	var obj Scene
	if err := json.Unmarshal(respBytes, &obj); err != nil {
		return Scene{}, err
	}

	obj.ID = ID

	return obj, nil
}
