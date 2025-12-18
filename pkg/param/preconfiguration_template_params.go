// Copyright (c) ZStack.io, Inc.

package param

// UpdatePreconfigurationTemplateDetailParam UpdatePreconfigurationTemplate详细参数
type UpdatePreconfigurationTemplateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"distribution,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"content,omitempty"`
}

// UpdatePreconfigurationTemplateParam UpdatePreconfigurationTemplate请求参数
type UpdatePreconfigurationTemplateParam struct {
	BaseParam
	Params UpdatePreconfigurationTemplateDetailParam `json:"params"` // 详细参数
}

