// Copyright (c) ZStack.io, Inc.

package param

// AddStackTemplateDetailParam AddStackTemplate详细参数
type AddStackTemplateDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"templateContent,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddStackTemplateParam AddStackTemplate请求参数
type AddStackTemplateParam struct {
	BaseParam
	Params AddStackTemplateDetailParam `json:"params"` // 详细参数
}

