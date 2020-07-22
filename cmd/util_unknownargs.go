package cmd

import (
	"errors"
	"fmt"
	"strings"
)

var errUnexpectedArguments = errors.New("unexpected arguments")

func errorOnUnknownArgs(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("%w: %s", errUnexpectedArguments, strings.Join(args, ", "))
	}

	return nil
}
