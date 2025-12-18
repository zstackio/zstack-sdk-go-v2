// Copyright (c) ZStack.io, Inc.

package param

// UpdateFiSecSecurityMachineDetailParam UpdateFiSecSecurityMachine详细参数
type UpdateFiSecSecurityMachineDetailParam struct {
	rest int `json:"port,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"managementIp,omitempty"`
	rest string `json:"model,omitempty"`
}

// UpdateFiSecSecurityMachineParam UpdateFiSecSecurityMachine请求参数
type UpdateFiSecSecurityMachineParam struct {
	BaseParam
	Params UpdateFiSecSecurityMachineDetailParam `json:"params"` // 详细参数
}

