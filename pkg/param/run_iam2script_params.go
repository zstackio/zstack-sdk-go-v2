// Copyright (c) ZStack.io, Inc.

package param

// RunIAM2ScriptDetailParam RunIAM2Script详细参数
type RunIAM2ScriptDetailParam struct {
	rest string `json:"scriptContent" validate:"required"` // 必填
	rest string `json:"scriptExecutor,omitempty"`
	rest []string `json:"scriptParams,omitempty"`
}

// RunIAM2ScriptParam RunIAM2Script请求参数
type RunIAM2ScriptParam struct {
	BaseParam
	Params RunIAM2ScriptDetailParam `json:"params"` // 详细参数
}

