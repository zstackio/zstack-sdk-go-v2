// Copyright (c) ZStack.io, Inc.

package param

// SetVmUserDefinedXmlHookScriptDetailParam SetVmUserDefinedXmlHookScript详细参数
type SetVmUserDefinedXmlHookScriptDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"xmlHookScriptBase64" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// SetVmUserDefinedXmlHookScriptParam SetVmUserDefinedXmlHookScript请求参数
type SetVmUserDefinedXmlHookScriptParam struct {
	BaseParam
	Params SetVmUserDefinedXmlHookScriptDetailParam `json:"params"` // 详细参数
}

