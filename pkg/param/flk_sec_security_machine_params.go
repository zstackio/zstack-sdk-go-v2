// Copyright (c) ZStack.io, Inc.

package param

// UpdateFlkSecSecurityMachineDetailParam UpdateFlkSecSecurityMachine详细参数
type UpdateFlkSecSecurityMachineDetailParam struct {
	rest int `json:"port,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"managementIp,omitempty"`
	rest string `json:"model,omitempty"`
}

// UpdateFlkSecSecurityMachineParam UpdateFlkSecSecurityMachine请求参数
type UpdateFlkSecSecurityMachineParam struct {
	BaseParam
	Params UpdateFlkSecSecurityMachineDetailParam `json:"params"` // 详细参数
}

