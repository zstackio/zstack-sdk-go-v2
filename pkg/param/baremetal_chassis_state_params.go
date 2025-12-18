// Copyright (c) ZStack.io, Inc.

package param

// ChangeBaremetalChassisStateDetailParam ChangeBaremetalChassisState详细参数
type ChangeBaremetalChassisStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeBaremetalChassisStateParam ChangeBaremetalChassisState请求参数
type ChangeBaremetalChassisStateParam struct {
	BaseParam
	Params ChangeBaremetalChassisStateDetailParam `json:"params"` // 详细参数
}

