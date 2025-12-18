// Copyright (c) ZStack.io, Inc.

package param

// CreateVmUserDefinedXmlHookScriptDetailParam CreateVmUserDefinedXmlHookScript detail param
type CreateVmUserDefinedXmlHookScriptDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	HookScript string `json:"hookScript" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmUserDefinedXmlHookScriptParam CreateVmUserDefinedXmlHookScript request param
type CreateVmUserDefinedXmlHookScriptParam struct {
	BaseParam
	Params CreateVmUserDefinedXmlHookScriptDetailParam `json:"params"`
}
