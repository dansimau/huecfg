package cmd

func init() {
	_, err := parser.AddCommand("api", "Interact with Hue Bridge APIs", "", &apiCmd{})
	if err != nil {
		panic(err)
	}
}

type apiCmd struct {
	V1 *apiV1Cmd `command:"v1" description:"Interact with v1 of the Hue Bridge API"`
	V2 *apiV2Cmd `command:"v2" description:"Interact with v2 (CLIP) of the Hue Bridge API"`
}
