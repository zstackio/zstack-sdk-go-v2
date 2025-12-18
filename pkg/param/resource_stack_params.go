// Copyright (c) ZStack.io, Inc.

package param

// UpdateResourceStackDetailParam UpdateResourceStack详细参数
type UpdateResourceStackDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest bool `json:"rollback,omitempty"`
	rest string `json:"templateContent,omitempty"`
	rest string `json:"parameters,omitempty"`
}

// UpdateResourceStackParam UpdateResourceStack请求参数
type UpdateResourceStackParam struct {
	BaseParam
	Params UpdateResourceStackDetailParam `json:"params"` // 详细参数
}

