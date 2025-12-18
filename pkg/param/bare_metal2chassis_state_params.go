// Copyright (c) ZStack.io, Inc.

package param

// ChangeBareMetal2ChassisStateDetailParam ChangeBareMetal2ChassisState详细参数
type ChangeBareMetal2ChassisStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeBareMetal2ChassisStateParam ChangeBareMetal2ChassisState请求参数
type ChangeBareMetal2ChassisStateParam struct {
	BaseParam
	Params ChangeBareMetal2ChassisStateDetailParam `json:"params"` // 详细参数
}

