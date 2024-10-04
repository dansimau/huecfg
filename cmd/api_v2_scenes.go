package cmd

import "github.com/dansimau/huecfg/pkg/jsonutil"

type apiV2SceneCmd struct {
	Get    *apiV2SceneGetCmd    `command:"get"`
	Post   *apiV2ScenePostCmd   `command:"post"`
	Put    *apiV2ScenePutCmd    `command:"put"`
	Delete *apiV2SceneDeleteCmd `command:"delete"`
}

type apiV2SceneGetCmd struct{}

func (c *apiV2SceneGetCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.GetScenes()
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2ScenePostCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`
}

func (c *apiV2ScenePostCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PostScene(c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2ScenePutCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`

	Arguments struct {
		ID string `description:"ID of the scene"`
	} `positional-args:"true" required:"true" positional-arg-name:"scene-ID"`
}

func (c *apiV2ScenePutCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.PutScene(c.Arguments.ID, c.Data)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}

type apiV2SceneDeleteCmd struct {
	Arguments struct {
		ID string `description:"ID of the scene"`
	} `positional-args:"true" required:"true" positional-arg-name:"scene-ID"`
}

func (c *apiV2SceneDeleteCmd) Execute(args []string) error {
	huev2 := cmd.getHueAPIV2()

	respBytes, err := huev2.DeleteScene(c.Arguments.ID)
	if err != nil {
		return err
	}

	return jsonutil.PrintBytes(respBytes)
}
