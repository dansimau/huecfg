package cmd

// huecfg api rules ...
type apiRulesCmd struct {
	Get    *apiRulesGetCmd    `command:"get" description:"Fetch the specified rule by ID"`
	GetAll *apiRulesGetAllCmd `command:"get-all" description:"Fetch all rule data at once"`
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
