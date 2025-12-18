// Copyright (c) ZStack.io, Inc.

package param

// AddPreconfigurationTemplateDetailParam AddPreconfigurationTemplate详细参数
type AddPreconfigurationTemplateDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"distribution" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"content" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddPreconfigurationTemplateParam AddPreconfigurationTemplate请求参数
type AddPreconfigurationTemplateParam struct {
	BaseParam
	Params AddPreconfigurationTemplateDetailParam `json:"params"` // 详细参数
}

