package cmd

//go:generate ./gen_api_read.sh ID=capabilities_get TYPE=apiV1CapabilitiesCmd FUNC_CALL=bridge.GetCapabilities()
type apiV1CapabilitiesCmd struct{}
