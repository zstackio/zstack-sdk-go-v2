// Copyright (c) ZStack.io, Inc.

package param

// UpdateJitSecurityMachineDetailParam UpdateJitSecurityMachine detail param
type UpdateJitSecurityMachineDetailParam struct {
	Port int `json:"port" validate:"required"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	Model string `json:"model,omitempty"`
}

// UpdateJitSecurityMachineParam UpdateJitSecurityMachine request param
type UpdateJitSecurityMachineParam struct {
	BaseParam
	Params UpdateJitSecurityMachineDetailParam `json:"params"`
}
