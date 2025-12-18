// Copyright (c) ZStack.io, Inc.

package param

// ChangeRoleStateDetailParam ChangeRoleState详细参数
type ChangeRoleStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeRoleStateParam ChangeRoleState请求参数
type ChangeRoleStateParam struct {
	BaseParam
	Params ChangeRoleStateDetailParam `json:"params"` // 详细参数
}

