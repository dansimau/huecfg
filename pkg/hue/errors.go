package hue

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errEmptyID = errors.New("ID cannot be empty")
)

type MultiError []error

func (m MultiError) errorStrings() (errStrings []string) {
	for _, e := range m {
		errStrings = append(errStrings, e.Error())
	}
	return errStrings
}

func (m MultiError) Error() string {
	if len(m) > 1 {
		return fmt.Sprintf("* %s", strings.Join(m.errorStrings(), "\n* "))
	}
	return m[0].Error()
}
