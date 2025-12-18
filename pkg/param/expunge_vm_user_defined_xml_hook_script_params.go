// Copyright (c) ZStack.io, Inc.

package param

// ExpungeVmUserDefinedXmlHookScriptDetailParam ExpungeVmUserDefinedXmlHookScript detail param
type ExpungeVmUserDefinedXmlHookScriptDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ExpungeVmUserDefinedXmlHookScriptParam ExpungeVmUserDefinedXmlHookScript request param
type ExpungeVmUserDefinedXmlHookScriptParam struct {
	BaseParam
	Params ExpungeVmUserDefinedXmlHookScriptDetailParam `json:"params"`
}
