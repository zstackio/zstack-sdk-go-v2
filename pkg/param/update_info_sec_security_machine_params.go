// Copyright (c) ZStack.io, Inc.

package param

// UpdateInfoSecSecurityMachineDetailParam UpdateInfoSecSecurityMachine detail param
type UpdateInfoSecSecurityMachineDetailParam struct {
	Password string `json:"password,omitempty"`
	Port int `json:"port,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	Model string `json:"model,omitempty"`
}

// UpdateInfoSecSecurityMachineParam UpdateInfoSecSecurityMachine request param
type UpdateInfoSecSecurityMachineParam struct {
	BaseParam
	Params UpdateInfoSecSecurityMachineDetailParam `json:"params"`
}
