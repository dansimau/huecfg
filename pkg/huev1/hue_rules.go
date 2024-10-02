package huev1

import (
	"encoding/json"
	"fmt"
)

type Rule struct {
	ID string

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

func (v Condition) String() string {
	return fmt.Sprintf("%s %s \"%s\"", v.Address, v.Operator, v.Value)
}

// GetRules gets a list of all rules that are in the bridge.
func (h *Hue) GetRules() ([]Rule, error) {
	respBytes, err := h.API.GetRules()
	if err != nil {
		return nil, err
	}

	if hueErr := parseAsHueError(respBytes); hueErr != nil {
		return nil, hueErr
	}

	var objs map[string]Rule
	if err := json.Unmarshal(respBytes, &objs); err != nil {
		return nil, err
	}

	res := []Rule{}
	for id, obj := range objs {
		obj.ID = id
		res = append(res, obj)
	}

	return res, nil
}

// GetRule returns a rule matching ID
func (h *Hue) GetRule(id string) (Rule, error) {
	respBytes, err := h.API.GetRule(id)
	if err != nil {
		return Rule{}, err
	}

	if hueErr := parseAsHueError(respBytes); hueErr != nil {
		return Rule{}, hueErr
	}

	var obj Rule
	if err := json.Unmarshal(respBytes, &obj); err != nil {
		return Rule{}, err
	}

	obj.ID = id

	return obj, nil
}
