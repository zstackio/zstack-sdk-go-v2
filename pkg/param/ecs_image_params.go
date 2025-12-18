// Copyright (c) ZStack.io, Inc.

package param

// UpdateEcsImageDetailParam UpdateEcsImage详细参数
type UpdateEcsImageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"name,omitempty"`
}

// UpdateEcsImageParam UpdateEcsImage请求参数
type UpdateEcsImageParam struct {
	BaseParam
	Params UpdateEcsImageDetailParam `json:"params"` // 详细参数
}

