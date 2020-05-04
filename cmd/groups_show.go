// Code generated by go generate; DO NOT EDIT.
package cmd

import (
	"fmt"
	"os"
	"reflect"

	"github.com/mikefarah/yq/v3/pkg/yqlib"
	"gopkg.in/yaml.v3"
)

type groupsShowCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"group-ID"`
}

func (c *groupsShowCmd) showCmdPostProcessFuncExists() bool {
	return reflect.ValueOf(c).MethodByName("PostProcessShowCmd").Kind() != reflect.Invalid
}

func (c *groupsShowCmd) showCmdPostProcessFuncCall(bytes []byte) ([]byte, error) {
	retVals := reflect.ValueOf(c).MethodByName("PostProcessShowCmd").Call([]reflect.Value{reflect.ValueOf(bytes)})

	var err error

	errVal := retVals[1].Interface()
	if errVal != nil {
		err = errVal.(error)
	}

	return retVals[0].Bytes(), err
}

func (c *groupsShowCmd) Execute(args []string) error {
	bridge := cmd.getHue()

	resp, err := bridge.GetGroup(c.Arguments.ID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
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

	return yqlib.ColorizeAndPrint(bytes, os.Stdout)
}
