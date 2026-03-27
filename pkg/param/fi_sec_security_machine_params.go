// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateFiSecSecurityMachineParamDetail UpdateFiSecSecurityMachine detail param
type UpdateFiSecSecurityMachineParamDetail struct {
	Port *int `json:"port,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ManagementIp *string `json:"managementIp,omitempty"`
	Model *string `json:"model,omitempty"`
}

// UpdateFiSecSecurityMachineParam UpdateFiSecSecurityMachine request param
type UpdateFiSecSecurityMachineParam struct {
	BaseParam
	Params UpdateFiSecSecurityMachineParamDetail `json:"updateFiSecSecurityMachine"`
}
// AddFiSecSecurityMachineParamDetail AddFiSecSecurityMachine detail param
type AddFiSecSecurityMachineParamDetail struct {
	Port int `json:"port" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp" validate:"required"`
	Model string `json:"model" validate:"required"`
	Type string `json:"type" validate:"required"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	SecretResourcePoolUuid string `json:"secretResourcePoolUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddFiSecSecurityMachineParam AddFiSecSecurityMachine request param
type AddFiSecSecurityMachineParam struct {
	BaseParam
	Params AddFiSecSecurityMachineParamDetail `json:"params"`
}
