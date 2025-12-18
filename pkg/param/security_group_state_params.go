// Copyright (c) ZStack.io, Inc.

package param

// ChangeSecurityGroupStateDetailParam ChangeSecurityGroupState详细参数
type ChangeSecurityGroupStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeSecurityGroupStateParam ChangeSecurityGroupState请求参数
type ChangeSecurityGroupStateParam struct {
	BaseParam
	Params ChangeSecurityGroupStateDetailParam `json:"params"` // 详细参数
}

