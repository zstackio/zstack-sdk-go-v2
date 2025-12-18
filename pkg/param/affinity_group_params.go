// Copyright (c) ZStack.io, Inc.

package param

// UpdateAffinityGroupDetailParam UpdateAffinityGroup详细参数
type UpdateAffinityGroupDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateAffinityGroupParam UpdateAffinityGroup请求参数
type UpdateAffinityGroupParam struct {
	BaseParam
	Params UpdateAffinityGroupDetailParam `json:"params"` // 详细参数
}

