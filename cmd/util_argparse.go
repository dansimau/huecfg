package cmd

import (
	"errors"
	"fmt"
	"strings"
)

var errMissingValue = errors.New("missing value")
var errUnexpectedValue = errors.New("unexpected value")

// extractParams takes a list of arguments and parses them into parameter
// key/value pairs.
// E.g.:
// --foo=bar -baz quux
func extractParams(args []string) (params map[string]string, remainder []string, err error) {
	params = map[string]string{}

	for i := 0; i < len(args); i++ {
		curArg := args[i]

		var param string
		var value string

		if !strings.HasPrefix(curArg, "-") {
			remainder = append(remainder, curArg)
			continue
		}

		// E.g.: --foo=bar
		if strings.Contains(curArg, "=") {
			components := strings.SplitN(curArg, "=", 2)
			param = components[0]
			value = components[1]
			// E.g.: --foo bar
		} else {
			param = curArg

			// Missing value, either out of bounds (i.e. last argument is a
			// param, not a value) or next argument is a param
			if i+1 == len(args) || strings.HasPrefix(args[i+1], "-") {
				return nil, nil, fmt.Errorf("%w for: %s", errMissingValue, curArg)
			}

			value = args[i+1]
			i++
		}

		params[strings.TrimLeft(param, "-")] = value
	}

	return params, remainder, nil
}
