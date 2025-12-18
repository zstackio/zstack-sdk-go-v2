// Copyright (c) ZStack.io, Inc.

package param

// UpdateVmUserDefinedXmlHookScriptDetailParam UpdateVmUserDefinedXmlHookScript detail param
type UpdateVmUserDefinedXmlHookScriptDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	HookScript string `json:"hookScript,omitempty"`
	StartupStrategy string `json:"startupStrategy,omitempty"`
}

// UpdateVmUserDefinedXmlHookScriptParam UpdateVmUserDefinedXmlHookScript request param
type UpdateVmUserDefinedXmlHookScriptParam struct {
	BaseParam
	Params UpdateVmUserDefinedXmlHookScriptDetailParam `json:"params"`
}
