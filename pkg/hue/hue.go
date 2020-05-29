package hue

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Hue represents a Hue Bridge
type Hue struct {
	API *API
}

// NewConn creates a connection to a Hue Bridge.
func NewConn(host string, username string) *Hue {
	return &Hue{
		API: &API{
			Host:     host,
			Username: username,
		},
	}
}

// Error represents an error from the Hue Bridge API
type Error struct {
	Address     string
	Description string
	Type        int
}

// Error is the description of the error return from the Hue Bridge API.
func (e Error) Error() string {
	return e.Description
}

type Success map[string]string

type SuccessMessages []Success

func (s SuccessMessages) successStrings() (successStrings []string) {
	for _, success := range s {
		for key, message := range success {
			successStrings = append(successStrings, fmt.Sprintf("%s: %s", key, message))
		}
	}
	return successStrings
}

func (s SuccessMessages) String() string {
	messages := s.successStrings()

	if len(messages) > 1 {
		return fmt.Sprintf("* %s", strings.Join(messages, "\n* "))
	}

	if len(messages) == 1 {
		return messages[0]
	}

	return ""
}

type Status struct {
	Success json.RawMessage
	Error   json.RawMessage
}

func (s Status) ToSuccess() (*Success, error) {
	v := &Success{}
	if err := json.Unmarshal(s.Success, &v); err != nil {
		return nil, err
	}

	return v, nil
}

func (s Status) ToError() (*Error, error) {
	v := &Error{}
	if err := json.Unmarshal(s.Error, &v); err != nil {
		return nil, err
	}

	return v, nil
}

func (s Status) ToInterface() (interface{}, error) {
	if v, err := s.ToError(); err == nil {
		return v, nil
	}
	if v, err := s.ToSuccess(); err == nil {
		return v, nil
	}
	return nil, errors.New("unknown status type or not a valid status")
}

type StatusResponse []Status

func (s StatusResponse) SuccessMessages() (successes SuccessMessages) {
	for _, status := range s {
		if statusSuccess, _ := status.ToSuccess(); statusSuccess != nil {
			successes = append(successes, *statusSuccess)
		}
	}
	return successes
}

// Errors returns a MultiError, which is a slide of all errors in the status
// response. If there are no errors, the MultiError slice will be nil.
func (s StatusResponse) Errors() (errs MultiError) {
	for _, status := range s {
		if statusErr, _ := status.ToError(); statusErr != nil {
			errs = append(errs, statusErr)
		}
	}
	return errs
}

// parseAsHueError tries to parse errors returned by the Hue Bridge in the
// response. If there are errors, it returns them as a MultiError. Otherwise,
// it returns nil.
func parseAsHueError(response []byte) MultiError {
	var statusResp StatusResponse
	if err := json.Unmarshal(response, &statusResp); err != nil {
		return nil
	}

	return statusResp.Errors()
}
