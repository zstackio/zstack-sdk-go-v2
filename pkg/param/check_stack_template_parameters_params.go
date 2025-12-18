// Copyright (c) ZStack.io, Inc.

package param

// CheckStackTemplateParametersDetailParam CheckStackTemplateParameters详细参数
type CheckStackTemplateParametersDetailParam struct {
	rest string `json:"type,omitempty"`
	rest string `json:"templateContent,omitempty"`
	rest string `json:"uuid,omitempty"`
}

// CheckStackTemplateParametersParam CheckStackTemplateParameters请求参数
type CheckStackTemplateParametersParam struct {
	BaseParam
	Params CheckStackTemplateParametersDetailParam `json:"params"` // 详细参数
}

