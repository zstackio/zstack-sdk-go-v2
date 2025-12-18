// Copyright (c) ZStack.io, Inc.

package param

// UpdateInfoSecSecurityMachineDetailParam UpdateInfoSecSecurityMachine详细参数
type UpdateInfoSecSecurityMachineDetailParam struct {
	rest string `json:"password,omitempty"`
	rest int `json:"port,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"managementIp,omitempty"`
	rest string `json:"model,omitempty"`
}

// UpdateInfoSecSecurityMachineParam UpdateInfoSecSecurityMachine请求参数
type UpdateInfoSecSecurityMachineParam struct {
	BaseParam
	Params UpdateInfoSecSecurityMachineDetailParam `json:"params"` // 详细参数
}

