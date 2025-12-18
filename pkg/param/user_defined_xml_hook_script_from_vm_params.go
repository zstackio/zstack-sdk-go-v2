// Copyright (c) ZStack.io, Inc.

package param

// DetachUserDefinedXmlHookScriptFromVmDetailParam DetachUserDefinedXmlHookScriptFromVm详细参数
type DetachUserDefinedXmlHookScriptFromVmDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"startupStrategy,omitempty"`
	rest string `json:"deleteMode,omitempty"`
}

// DetachUserDefinedXmlHookScriptFromVmParam DetachUserDefinedXmlHookScriptFromVm请求参数
type DetachUserDefinedXmlHookScriptFromVmParam struct {
	BaseParam
	Params DetachUserDefinedXmlHookScriptFromVmDetailParam `json:"params"` // 详细参数
}

