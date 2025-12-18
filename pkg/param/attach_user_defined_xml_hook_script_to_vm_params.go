// Copyright (c) ZStack.io, Inc.

package param

// AttachUserDefinedXmlHookScriptToVmDetailParam AttachUserDefinedXmlHookScriptToVm detail param
type AttachUserDefinedXmlHookScriptToVmDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	XmlHookUuid string `json:"xmlHookUuid" validate:"required"`
	StartupStrategy string `json:"startupStrategy,omitempty"`
}

// AttachUserDefinedXmlHookScriptToVmParam AttachUserDefinedXmlHookScriptToVm request param
type AttachUserDefinedXmlHookScriptToVmParam struct {
	BaseParam
	Params AttachUserDefinedXmlHookScriptToVmDetailParam `json:"params"`
}
