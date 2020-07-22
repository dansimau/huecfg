package cmd

// huecfg api rules ...
type apiRulesCmd struct {
	Create *apiRulesCreateCmd `command:"create" description:"Create a new rule"`
	Get    *apiRulesGetCmd    `command:"get" description:"Fetch the specified rule by ID"`
	GetAll *apiRulesGetAllCmd `command:"get-all" description:"Fetch all rule data at once"`
}

// huecfg api rules create
//go:generate ./gen_api_write.sh ID=rules_create TYPE=apiRulesCreateCmd DATA=c.Data FUNC_CALL=bridge.CreateResourceLink(data)
type apiRulesCreateCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`
}

// huecfg api rules get
//go:generate ./gen_api_read.sh ID=rules_get TYPE=apiRulesGetCmd FUNC_CALL=bridge.GetRule(c.Arguments.ID)
type apiRulesGetCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"rule-ID"`
}

// huecfg api rules get-all
//go:generate ./gen_api_read.sh ID=rules_get_all TYPE=apiRulesGetAllCmd FUNC_CALL=bridge.GetRules()
type apiRulesGetAllCmd struct{}
