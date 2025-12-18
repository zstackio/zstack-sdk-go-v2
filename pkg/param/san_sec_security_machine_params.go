// Copyright (c) ZStack.io, Inc.

package param

// UpdateSanSecSecurityMachineDetailParam UpdateSanSecSecurityMachine详细参数
type UpdateSanSecSecurityMachineDetailParam struct {
	rest string `json:"password,omitempty"`
	rest int `json:"port,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"managementIp,omitempty"`
	rest string `json:"model,omitempty"`
}

// UpdateSanSecSecurityMachineParam UpdateSanSecSecurityMachine请求参数
type UpdateSanSecSecurityMachineParam struct {
	BaseParam
	Params UpdateSanSecSecurityMachineDetailParam `json:"params"` // 详细参数
}

