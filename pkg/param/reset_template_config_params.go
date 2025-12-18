// Copyright (c) ZStack.io, Inc.

package param

// ResetTemplateConfigDetailParam ResetTemplateConfig详细参数
type ResetTemplateConfigDetailParam struct {
	rest string `json:"templateUuid" validate:"required"` // 必填
}

// ResetTemplateConfigParam ResetTemplateConfig请求参数
type ResetTemplateConfigParam struct {
	BaseParam
	Params ResetTemplateConfigDetailParam `json:"params"` // 详细参数
}

