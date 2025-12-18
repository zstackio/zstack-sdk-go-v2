// Copyright (c) ZStack.io, Inc.

package param

// DetachUserDefinedXmlHookScriptFromVmDetailParam DetachUserDefinedXmlHookScriptFromVm detail param
type DetachUserDefinedXmlHookScriptFromVmDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	StartupStrategy string `json:"startupStrategy,omitempty"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DetachUserDefinedXmlHookScriptFromVmParam DetachUserDefinedXmlHookScriptFromVm request param
type DetachUserDefinedXmlHookScriptFromVmParam struct {
	BaseParam
	Params DetachUserDefinedXmlHookScriptFromVmDetailParam `json:"params"`
}
