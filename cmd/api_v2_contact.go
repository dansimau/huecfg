package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2ContactCmd struct {
	Get *apiV2ContactGetCmd `command:"get"`
	Put *apiV2ContactPutCmd `command:"put"`
}

type apiV2ContactGetCmd struct {
	Arguments struct {
		ID string `description:"ID of the contact sensor" optional:"true"`
	} `positional-args:"true" positional-arg-name:"contact-ID"`
}

func (c *apiV2ContactGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID != "" {
		respBytes, err = huev2.GetContact(c.Arguments.ID)
	} else {
		respBytes, err = huev2.GetContacts()
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2ContactPutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the contact sensor"`
	} `positional-args:"true" required:"true" positional-arg-name:"contact-ID"`
}

func (c *apiV2ContactPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutContact(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
