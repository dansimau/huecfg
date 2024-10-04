package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2ZGPConnectivityCmd struct {
	Get *apiV2ZGPConnectivityGetCmd `command:"get" description:"Get all ZGP connectivity resources"`
	Put *apiV2ZGPConnectivityPutCmd `command:"put" description:"Update a ZGP connectivity resource"`
}

type apiV2ZGPConnectivityGetCmd struct {
	Arguments struct {
		ID string `positional-arg-name:"id" required:"false" description:"ID of the ZGP connectivity resource"`
	} `positional-args:"yes"`
}

func (c *apiV2ZGPConnectivityGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	var respBytes []byte
	var err error

	if c.Arguments.ID == "" {
		respBytes, err = huev2.GetZGPConnectivities()
	} else {
		respBytes, err = huev2.GetZGPConnectivity(c.Arguments.ID)
	}

	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2ZGPConnectivityPutCmd struct {
	Arguments struct {
		ID   string `positional-arg-name:"id" required:"true" description:"ID of the ZGP connectivity resource"`
		Data string `positional-arg-name:"data" required:"true" description:"JSON data to update the ZGP connectivity resource"`
	} `positional-args:"yes"`
}

func (c *apiV2ZGPConnectivityPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	data := []byte(c.Arguments.Data)
	respBytes, err := huev2.PutZGPConnectivity(c.Arguments.ID, data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
