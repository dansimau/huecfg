#!/usr/bin/env bash
#
# Generates Go source file for a show command.
#

set -eu

for arg in "$@"; do
	varname="$(echo "$arg" | cut -f1 -d '=')"
	varval="$(echo "$arg" | cut -f2 -d '=')"

	eval "$varname='$varval'"
done

#OBJ_NAME=rule
#GET_OBJ_FUNC="GetRule"


cat <<EOF>${OBJ_NAME}s_show_gen.go
// Code generated by go generate; DO NOT EDIT.
package cmd

import (
	"bytes"
	"io"
	"os"
	"reflect"

	"github.com/mikefarah/yq/v3/pkg/yqlib"
	"gopkg.in/yaml.v3"
)

type ${OBJ_NAME}sShowCmd struct {
	Arguments struct {
		IDs []string
	} \`positional-args:"true" required:"true" positional-arg-name:"${OBJ_NAME}-ID"\`
}

func (c *${OBJ_NAME}sShowCmd) showCmdPostProcessFuncExists() bool {
	return reflect.ValueOf(c).MethodByName("PostProcessShowCmd").Kind() != reflect.Invalid
}

func (c *${OBJ_NAME}sShowCmd) showCmdPostProcessFuncCall(bytes []byte) ([]byte, error) {
	retVals := reflect.ValueOf(c).MethodByName("PostProcessShowCmd").Call([]reflect.Value{reflect.ValueOf(bytes)})

	var err error

	errVal := retVals[1].Interface()
	if errVal != nil {
		err = errVal.(error)
	}

	return retVals[0].Bytes(), err
}

func (c *${OBJ_NAME}sShowCmd) Execute(args []string) error {
	if err := errorOnUnknownArgs(args); err != nil {
		return err
	}

	bridge := cmd.getHue()

	outputBytes := &bytes.Buffer{}

	for _, id := range c.Arguments.IDs {
		resp, err := bridge.${GET_OBJ_FUNC}(id)
		if err != nil {
			return err
		}

		bytes, err := yaml.Marshal(resp)
		if err != nil {
			return err
		}

		if c.showCmdPostProcessFuncExists() {
			bytes, err = c.showCmdPostProcessFuncCall(bytes)
			if err != nil {
				return err
			}
		}

		if len(c.Arguments.IDs) > 1 {
			if _, err := outputBytes.Write([]byte("---\n")); err != nil {
				return err
			}
		}

		if err := yqlib.ColorizeAndPrint(bytes, outputBytes); err != nil {
			return err
		}
	}

	if _, err := io.Copy(os.Stdout, outputBytes); err != nil {
		return err
	}

	return nil
}
EOF
