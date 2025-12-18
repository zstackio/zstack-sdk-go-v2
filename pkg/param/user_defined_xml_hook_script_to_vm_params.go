// Copyright (c) ZStack.io, Inc.

package param

// AttachUserDefinedXmlHookScriptToVmDetailParam AttachUserDefinedXmlHookScriptToVm详细参数
type AttachUserDefinedXmlHookScriptToVmDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"xmlHookUuid" validate:"required"` // 必填
	rest string `json:"startupStrategy,omitempty"`
}

// AttachUserDefinedXmlHookScriptToVmParam AttachUserDefinedXmlHookScriptToVm请求参数
type AttachUserDefinedXmlHookScriptToVmParam struct {
	BaseParam
	Params AttachUserDefinedXmlHookScriptToVmDetailParam `json:"params"` // 详细参数
}

