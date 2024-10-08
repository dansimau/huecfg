#!/usr/bin/env bash
#
# Generates Go source file for a show command.
#

set -eu

declare -ar expected_args=(
    ID
    TYPE
	DATA
    FUNC_CALL
)

for arg in "$@"; do
	varname="$(echo "$arg" | cut -f1 -d '=')"
	varval="$(echo "$arg" | cut -f2 -d '=')"

	eval "$varname='$varval'"
done

for arg in "${expected_args[@]}"; do
    if [ -z "$(eval echo \$${arg})" ]; then
        echo "ERROR: Missing arg: ${arg}=" >&2
        exit 1
    fi
done

cat <<EOF>api_v1_${ID}_gen.go
// Code generated by go generate; DO NOT EDIT.
package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

func (c *${TYPE}) Execute(args []string) error {
	if err := errorOnUnknownArgs(args); err != nil {
		return err
	}

	bridge := cmd.getHueAPI()

	data, err := userInputToJSON(${DATA})
	if err != nil {
		return err
	}

	respBytes, err := ${FUNC_CALL}
	if err != nil {
		return err
	}

	if err := jsonutil.PrintBytes(respBytes); err != nil {
		return err
	}

	return nil
}
EOF
