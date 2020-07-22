package cmd

import (
	"github.com/dansimau/huecfg/pkg/hue"
)

type lightsSetAttrCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"light-ID"`

	Name *string `long:"name" description:"New name for the light"`
}

func (c *lightsSetAttrCmd) Execute(args []string) error {
	if err := errorOnUnknownArgs(args); err != nil {
		return err
	}

	if err := errorOnUnknownArgs(args); err != nil {
		return err
	}

	bridge := cmd.getHue()

	attrs := hue.SetLightAttributeParams{
		Name: c.Name,
	}

	_, err := bridge.SetLightAttributes(c.Arguments.ID, attrs)
	if err != nil {
		return err
	}

	return nil
}

type lightsSetStateCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"light-ID"`

	On         *bool `long:"on" description:"Turn light on"`
	Off        *bool `long:"off" description:"Turn light off"`
	Brightness *int  `long:"bri" description:"Brightness to set the light to (1-255)"`
}

func (c *lightsSetStateCmd) ExecuteNoArgs(args []string) error {
	if err := errorOnUnknownArgs(args); err != nil {
		return err
	}

	bridge := cmd.getHue()

	params := hue.SetLightStateParams{
		Bri: c.Brightness,
	}

	if c.On != nil {
		params.On = c.On
	}
	if c.Off != nil {
		v := !(*c.Off)
		params.On = &v
	}

	_, err := bridge.SetLightState(c.Arguments.ID, params)
	if err != nil {
		return err
	}

	return nil
}
