package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2RoomCmd struct {
	Get    *apiV2RoomGetCmd    `command:"get"`
	Post   *apiV2RoomPostCmd   `command:"post"`
	Put    *apiV2RoomPutCmd    `command:"put"`
	Delete *apiV2RoomDeleteCmd `command:"delete"`
}

type apiV2RoomGetCmd struct{}

func (c *apiV2RoomGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.GetRooms()
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2RoomPostCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`
}

func (c *apiV2RoomPostCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PostRoom(c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2RoomPutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the room"`
	} `positional-args:"true" required:"true" positional-arg-name:"room-ID"`
}

func (c *apiV2RoomPutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutRoom(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2RoomDeleteCmd struct {
	Arguments struct {
		ID string `description:"ID of the room"`
	} `positional-args:"true" required:"true" positional-arg-name:"room-ID"`
}

func (c *apiV2RoomDeleteCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.DeleteRoom(c.Arguments.ID)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
