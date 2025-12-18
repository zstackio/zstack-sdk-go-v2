// Copyright (c) ZStack.io, Inc.

package param

// UpdateSanSecSecurityMachineDetailParam UpdateSanSecSecurityMachine detail param
type UpdateSanSecSecurityMachineDetailParam struct {
	Password string `json:"password,omitempty"`
	Port int `json:"port,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	Model string `json:"model,omitempty"`
}

// UpdateSanSecSecurityMachineParam UpdateSanSecSecurityMachine request param
type UpdateSanSecSecurityMachineParam struct {
	BaseParam
	Params UpdateSanSecSecurityMachineDetailParam `json:"params"`
}
