package cmd

type apiV1RulesCmd struct {
	Create *apiRulesCreateCmd `command:"create" description:"Create a new rule"`
	Delete *apiRulesDeleteCmd `command:"delete" description:"Delete a device from the bridge"`
	Get    *apiRulesGetCmd    `command:"get" description:"Fetch the specified rule by ID"`
	GetAll *apiRulesGetAllCmd `command:"get-all" description:"Fetch all rule data at once"`
}

//go:generate ./gen_api_write.sh ID=rules_create TYPE=apiV1RulesCreateCmd DATA=c.Data FUNC_CALL=bridge.CreateRule(data)
type apiRulesCreateCmd struct {
	Data string `long:"data" description:"JSON data to send" default:"-"`
}

//go:generate ./gen_api_read.sh ID=rules_delete TYPE=apiV1RulesDeleteCmd FUNC_CALL=bridge.DeleteRule(c.Arguments.ID)
type apiRulesDeleteCmd struct {
	Arguments struct {
		ID string `description:"ID of the rule to delete."`
	} `positional-args:"true" required:"true" positional-arg-name:"rule-ID"`
}

//go:generate ./gen_api_read.sh ID=rules_get TYPE=apiV1RulesGetCmd FUNC_CALL=bridge.GetRule(c.Arguments.ID)
type apiRulesGetCmd struct {
	Arguments struct {
		ID string
	} `positional-args:"true" required:"true" positional-arg-name:"rule-ID"`
}

//go:generate ./gen_api_read.sh ID=rules_get_all TYPE=apiV1RulesGetAllCmd FUNC_CALL=bridge.GetRules()
type apiRulesGetAllCmd struct{}
