// Copyright (c) ZStack.io, Inc.

package param

// GetVmXmlHookScriptDetailParam GetVmXmlHookScript detail param
type GetVmXmlHookScriptDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
}

// GetVmXmlHookScriptParam GetVmXmlHookScript request param
type GetVmXmlHookScriptParam struct {
	BaseParam
	Params GetVmXmlHookScriptDetailParam `json:"params"`
}
