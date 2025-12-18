// Copyright (c) ZStack.io, Inc.

package param

// UpdateVpcHaGroupDetailParam UpdateVpcHaGroup详细参数
type UpdateVpcHaGroupDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateVpcHaGroupParam UpdateVpcHaGroup请求参数
type UpdateVpcHaGroupParam struct {
	BaseParam
	Params UpdateVpcHaGroupDetailParam `json:"params"` // 详细参数
}

