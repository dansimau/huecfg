package hue

import (
	"encoding/json"
)

type ResourceLinkType string

const ResourceLinkTypeLink ResourceLinkType = "Link"

type ResourceLink struct {
	ID int

	// Class of resourcelink given by application. The resourcelink class can be
	// used to identify resourcelink with the same purpose, like classid 1 for
	// wake-up, 2 for going to sleep, etc. (best practice use range 1 – 10000)
	ClassID int

	// Human readable description of what this resourcelink does. If not
	// specified it's set to "".
	Description string

	// References to resources which are used by this resourcelink resource. In
	// case the referenced resource was created with "recycle":true and no other
	// references are present, the resourcelink resource will be automatically
	// deleted when all links are empty.
	Links []string

	// Human readable name for this resourcelink
	Name string

	// Not writeable, this respresents the owner (username) of the creator of
	// the resourcelink
	Owner string

	// When true: Resource is automatically deleted when not referenced anymore
	// in any resource link. Only on creation of resourcelink. "false" when
	// omitted.
	Recycle bool

	// Not writeable and there is only 1 type: "Link"
	Type ResourceLinkType
}

func (h *Hue) GetResourceLinks() ([]ResourceLink, error) {
	respBytes, err := h.API.GetResourceLinks()
	if err != nil {
		return nil, err
	}

	var objs map[int]ResourceLink
	if err := json.Unmarshal(respBytes, &objs); err != nil {
		return nil, err
	}

	var res = []ResourceLink{}
	for ID, obj := range objs {
		obj.ID = ID
		res = append(res, obj)
	}

	return res, nil
}

func (h *Hue) GetResourceLink(ID int) (ResourceLink, error) {
	respBytes, err := h.API.GetResourceLink(ID)
	if err != nil {
		return ResourceLink{}, err
	}

	var obj ResourceLink
	if err := json.Unmarshal(respBytes, &obj); err != nil {
		return ResourceLink{}, err
	}

	obj.ID = ID

	return obj, nil
}
