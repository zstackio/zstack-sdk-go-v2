// Copyright (c) ZStack.io, Inc.

package param

// UpdateSecurityMachineDetailParam UpdateSecurityMachine detail param
type UpdateSecurityMachineDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	Model string `json:"model,omitempty"`
}

// UpdateSecurityMachineParam UpdateSecurityMachine request param
type UpdateSecurityMachineParam struct {
	BaseParam
	Params UpdateSecurityMachineDetailParam `json:"params"`
}
