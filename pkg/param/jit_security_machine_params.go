// Copyright (c) ZStack.io, Inc.

package param

// UpdateJitSecurityMachineDetailParam UpdateJitSecurityMachine详细参数
type UpdateJitSecurityMachineDetailParam struct {
	rest int `json:"port" validate:"required"` // 必填
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"managementIp,omitempty"`
	rest string `json:"model,omitempty"`
}

// UpdateJitSecurityMachineParam UpdateJitSecurityMachine请求参数
type UpdateJitSecurityMachineParam struct {
	BaseParam
	Params UpdateJitSecurityMachineDetailParam `json:"params"` // 详细参数
}

