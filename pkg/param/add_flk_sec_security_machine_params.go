// Copyright (c) ZStack.io, Inc.

package param

// AddFlkSecSecurityMachineDetailParam AddFlkSecSecurityMachine detail param
type AddFlkSecSecurityMachineDetailParam struct {
	Port int `json:"port" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp" validate:"required"`
	Model string `json:"model" validate:"required"`
	Type string `json:"type" validate:"required"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	SecretResourcePoolUuid string `json:"secretResourcePoolUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddFlkSecSecurityMachineParam AddFlkSecSecurityMachine request param
type AddFlkSecSecurityMachineParam struct {
	BaseParam
	Params AddFlkSecSecurityMachineDetailParam `json:"params"`
}
