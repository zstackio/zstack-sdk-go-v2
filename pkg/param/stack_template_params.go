// Copyright (c) ZStack.io, Inc.

package param

// UpdateStackTemplateDetailParam UpdateStackTemplate详细参数
type UpdateStackTemplateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest bool `json:"state,omitempty"`
	rest string `json:"templateContent,omitempty"`
}

// UpdateStackTemplateParam UpdateStackTemplate请求参数
type UpdateStackTemplateParam struct {
	BaseParam
	Params UpdateStackTemplateDetailParam `json:"params"` // 详细参数
}

