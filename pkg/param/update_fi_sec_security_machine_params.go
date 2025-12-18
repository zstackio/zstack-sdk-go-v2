// Copyright (c) ZStack.io, Inc.

package param

// UpdateFiSecSecurityMachineDetailParam UpdateFiSecSecurityMachine detail param
type UpdateFiSecSecurityMachineDetailParam struct {
	Port int `json:"port,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	Model string `json:"model,omitempty"`
}

// UpdateFiSecSecurityMachineParam UpdateFiSecSecurityMachine request param
type UpdateFiSecSecurityMachineParam struct {
	BaseParam
	Params UpdateFiSecSecurityMachineDetailParam `json:"params"`
}
