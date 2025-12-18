// Copyright (c) ZStack.io, Inc.

package param

// UpdateTemplateConfigDetailParam UpdateTemplateConfig详细参数
type UpdateTemplateConfigDetailParam struct {
	rest string `json:"templateUuid" validate:"required"` // 必填
	rest string `json:"category" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"value" validate:"required"` // 必填
}

// UpdateTemplateConfigParam UpdateTemplateConfig请求参数
type UpdateTemplateConfigParam struct {
	BaseParam
	Params UpdateTemplateConfigDetailParam `json:"params"` // 详细参数
}

