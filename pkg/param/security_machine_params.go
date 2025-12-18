// Copyright (c) ZStack.io, Inc.

package param

// UpdateSecurityMachineDetailParam UpdateSecurityMachine详细参数
type UpdateSecurityMachineDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"managementIp,omitempty"`
	rest string `json:"model,omitempty"`
}

// UpdateSecurityMachineParam UpdateSecurityMachine请求参数
type UpdateSecurityMachineParam struct {
	BaseParam
	Params UpdateSecurityMachineDetailParam `json:"params"` // 详细参数
}

