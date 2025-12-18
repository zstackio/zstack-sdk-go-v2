// Copyright (c) ZStack.io, Inc.

package param

// CreateVmUserDefinedXmlHookScriptDetailParam CreateVmUserDefinedXmlHookScript详细参数
type CreateVmUserDefinedXmlHookScriptDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"hookScript" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVmUserDefinedXmlHookScriptParam CreateVmUserDefinedXmlHookScript请求参数
type CreateVmUserDefinedXmlHookScriptParam struct {
	BaseParam
	Params CreateVmUserDefinedXmlHookScriptDetailParam `json:"params"` // 详细参数
}

