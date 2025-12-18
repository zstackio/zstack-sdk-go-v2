// Copyright (c) ZStack.io, Inc.

package param

// RevertTemplateConfigDetailParam RevertTemplateConfig详细参数
type RevertTemplateConfigDetailParam struct {
	rest string `json:"templateUuid" validate:"required"` // 必填
}

// RevertTemplateConfigParam RevertTemplateConfig请求参数
type RevertTemplateConfigParam struct {
	BaseParam
	Params RevertTemplateConfigDetailParam `json:"params"` // 详细参数
}

