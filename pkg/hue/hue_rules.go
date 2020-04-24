package hue

import (
	"encoding/json"
)

type Rule struct {
	ID int

	Created        AbsoluteTime
	Name           string
	LastTriggered  *AbsoluteTime
	Owner          string
	TimesTriggered int
	Status         string

	Actions    []Action
	Conditions []Condition
}

type Action struct {
	Address string
	Method  string
	Body    map[string]interface{}
}

type Condition struct {
	Address  string
	Operator string
	Value    string
}

// GetRules gets a list of all rules that are in the bridge.
func (h *Hue) GetRules() ([]Rule, error) {
	respBytes, err := h.API.GetRules()
	if err != nil {
		return nil, err
	}

	var objs map[int]Rule
	if err := json.Unmarshal(respBytes, &objs); err != nil {
		return nil, err
	}

	var res = []Rule{}
	for ID, obj := range objs {
		obj.ID = ID
		res = append(res, obj)
	}

	return res, nil
}

// GetRule returns a rule matching ID
func (h *Hue) GetRule(ID int) (Rule, error) {
	respBytes, err := h.API.GetRule(ID)
	if err != nil {
		return Rule{}, err
	}

	var obj Rule
	if err := json.Unmarshal(respBytes, &obj); err != nil {
		return Rule{}, err
	}

	obj.ID = ID

	return obj, nil
}
