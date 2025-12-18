// Copyright (c) ZStack.io, Inc.

package param

// SetVmUserDefinedXmlHookScriptDetailParam SetVmUserDefinedXmlHookScript detail param
type SetVmUserDefinedXmlHookScriptDetailParam struct {
	VmInstanceUuid string `json:"vmInstanceUuid" validate:"required"`
	XmlHookScriptBase64 string `json:"xmlHookScriptBase64" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// SetVmUserDefinedXmlHookScriptParam SetVmUserDefinedXmlHookScript request param
type SetVmUserDefinedXmlHookScriptParam struct {
	BaseParam
	Params SetVmUserDefinedXmlHookScriptDetailParam `json:"params"`
}
