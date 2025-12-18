// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmUserDefinedXmlHookScriptDetailParam DeleteVmUserDefinedXmlHookScript detail param
type DeleteVmUserDefinedXmlHookScriptDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVmUserDefinedXmlHookScriptParam DeleteVmUserDefinedXmlHookScript request param
type DeleteVmUserDefinedXmlHookScriptParam struct {
	BaseParam
	Params DeleteVmUserDefinedXmlHookScriptDetailParam `json:"params"`
}
