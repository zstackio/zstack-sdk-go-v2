// Copyright (c) ZStack.io, Inc.

package param

// UpdateFlkSecSecurityMachineDetailParam UpdateFlkSecSecurityMachine detail param
type UpdateFlkSecSecurityMachineDetailParam struct {
	Port int `json:"port,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	Model string `json:"model,omitempty"`
}

// UpdateFlkSecSecurityMachineParam UpdateFlkSecSecurityMachine request param
type UpdateFlkSecSecurityMachineParam struct {
	BaseParam
	Params UpdateFlkSecSecurityMachineDetailParam `json:"params"`
}
