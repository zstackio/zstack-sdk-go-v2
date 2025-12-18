// Copyright (c) ZStack.io, Inc.

package param

// ApplyTemplateConfigDetailParam ApplyTemplateConfig详细参数
type ApplyTemplateConfigDetailParam struct {
	rest string `json:"templateUuid" validate:"required"` // 必填
}

// ApplyTemplateConfigParam ApplyTemplateConfig请求参数
type ApplyTemplateConfigParam struct {
	BaseParam
	Params ApplyTemplateConfigDetailParam `json:"params"` // 详细参数
}

