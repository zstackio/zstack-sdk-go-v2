// Copyright (c) ZStack.io, Inc.

package param

// UpdateEcsSecurityGroupDetailParam UpdateEcsSecurityGroup详细参数
type UpdateEcsSecurityGroupDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"name,omitempty"`
}

// UpdateEcsSecurityGroupParam UpdateEcsSecurityGroup请求参数
type UpdateEcsSecurityGroupParam struct {
	BaseParam
	Params UpdateEcsSecurityGroupDetailParam `json:"params"` // 详细参数
}

