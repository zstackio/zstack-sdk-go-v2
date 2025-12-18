// Copyright (c) ZStack.io, Inc.

package param

// GetVmXmlHookScriptDetailParam GetVmXmlHookScript详细参数
type GetVmXmlHookScriptDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// GetVmXmlHookScriptParam GetVmXmlHookScript请求参数
type GetVmXmlHookScriptParam struct {
	BaseParam
	Params GetVmXmlHookScriptDetailParam `json:"params"` // 详细参数
}

