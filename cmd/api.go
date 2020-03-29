package cmd

func init() {
	_, err := parser.AddCommand("api", "Interact with the Hue Bridge API", "", api)
	if err != nil {
		panic(err)
	}
}

var api = &apiCmd{}

type apiCmd struct {
	Host     string `short:"a" long:"host" description:"host address for Hue Bridge" required:"true"`
	Username string `short:"u" long:"username" description:"username from Hue Bridge registration"`

	APILights *apiLightsCmd `command:"lights" description:"List all lists in the system"`
	APIConfig *apiConfigCmd `command:"config" description:"Get or set Hue Bridge configuration"`
}
