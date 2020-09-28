package hue

import (
	"encoding/json"
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
	Success *Success `json:,omitempty`
	Error   *Error   `json:,omitempty`
}

type StatusResponse []Status

func (s StatusResponse) SuccessMessages() (successes SuccessMessages) {
	for _, status := range s {
		if status.Success != nil {
			successes = append(successes, *status.Success)
		}
	}
	return successes
}

// Errors returns a MultiError, which is a slice of all errors in the status
// response. If there are no errors, the MultiError slice will be nil.
func (s StatusResponse) Errors() (errs MultiError) {
	for _, status := range s {
		if status.Error != nil {
			errs = append(errs, status.Error)
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
