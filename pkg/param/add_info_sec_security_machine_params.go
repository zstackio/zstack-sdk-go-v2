// Copyright (c) ZStack.io, Inc.

package param

// AddInfoSecSecurityMachineDetailParam AddInfoSecSecurityMachine detail param
type AddInfoSecSecurityMachineDetailParam struct {
	Password string `json:"password" validate:"required"`
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

// AddInfoSecSecurityMachineParam AddInfoSecSecurityMachine request param
type AddInfoSecSecurityMachineParam struct {
	BaseParam
	Params AddInfoSecSecurityMachineDetailParam `json:"params"`
}
