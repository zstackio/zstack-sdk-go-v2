// Copyright (c) ZStack.io, Inc.

package param

// AddSanSecSecurityMachineDetailParam AddSanSecSecurityMachine detail param
type AddSanSecSecurityMachineDetailParam struct {
	Port int `json:"port" validate:"required"`
	Password string `json:"password" validate:"required"`
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

// AddSanSecSecurityMachineParam AddSanSecSecurityMachine request param
type AddSanSecSecurityMachineParam struct {
	BaseParam
	Params AddSanSecSecurityMachineDetailParam `json:"params"`
}
