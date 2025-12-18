// Copyright (c) ZStack.io, Inc.

package param

// UpdateEcsVpcDetailParam UpdateEcsVpc详细参数
type UpdateEcsVpcDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateEcsVpcParam UpdateEcsVpc请求参数
type UpdateEcsVpcParam struct {
	BaseParam
	Params UpdateEcsVpcDetailParam `json:"params"` // 详细参数
}

