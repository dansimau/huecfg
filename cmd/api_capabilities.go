package cmd

// huecfg api capabilities
//go:generate ./gen_api_read.sh ID=capabilities_get TYPE=apiCapabilitiesCmd FUNC_CALL=bridge.GetCapabilities()
type apiCapabilitiesCmd struct{}
