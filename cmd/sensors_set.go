package cmd

import "github.com/dansimau/huecfg/pkg/huev1"

type sensorsSetAttrCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"sensor-ID"`

	Name *string `long:"name" description:"New name for the sensor"`
}

func (c *sensorsSetAttrCmd) Execute(args []string) error {
	if err := errorOnUnknownArgs(args); err != nil {
		return err
	}

	bridge := cmd.getHue()

	attrs := huev1.SetSensorAttributeParams{
		Name: c.Name,
	}

	_, err := bridge.SetSensorAttributes(c.Arguments.ID, attrs)
	if err != nil {
		return err
	}

	return nil
}

type sensorsSetConfigCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"sensor-ID"`

	On *bool `long:"on" description:"Enable/disable the sensor"`
}

func (c *sensorsSetConfigCmd) Execute(args []string) error {
	if err := errorOnUnknownArgs(args); err != nil {
		return err
	}

	bridge := cmd.getHue()

	attrs := huev1.SetSensorConfigDefaultParams{
		On: c.On,
	}

	_, err := bridge.SetSensorConfig(c.Arguments.ID, attrs)
	if err != nil {
		return err
	}

	return nil
}

type sensorsSetStateCmd struct{}

func (c *sensorsSetStateCmd) Execute(args []string) error {
	params, args, err := extractParams(args)
	if err != nil {
		return err
	}

	if len(args) > 1 {
		return errorOnUnknownArgs(args[1:])
	}

	id := args[0]

	data := map[string]interface{}{}
	for key, value := range params {
		if err := setFieldValue(data, key, value); err != nil {
			return err
		}
	}

	bridge := cmd.getHue()

	_, err = bridge.SetSensorState(id, data)
	if err != nil {
		return err
	}

	return nil
}
