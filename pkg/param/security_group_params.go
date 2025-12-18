// Copyright (c) ZStack.io, Inc.

package param

// DeleteSecurityGroupDetailParam DeleteSecurityGroup详细参数
type DeleteSecurityGroupDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteSecurityGroupParam DeleteSecurityGroup请求参数
type DeleteSecurityGroupParam struct {
	BaseParam
	Params DeleteSecurityGroupDetailParam `json:"params"` // 详细参数
}

