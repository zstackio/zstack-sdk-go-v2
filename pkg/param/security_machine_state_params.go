// Copyright (c) ZStack.io, Inc.

package param

// ChangeSecurityMachineStateDetailParam ChangeSecurityMachineState详细参数
type ChangeSecurityMachineStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeSecurityMachineStateParam ChangeSecurityMachineState请求参数
type ChangeSecurityMachineStateParam struct {
	BaseParam
	Params ChangeSecurityMachineStateDetailParam `json:"params"` // 详细参数
}

