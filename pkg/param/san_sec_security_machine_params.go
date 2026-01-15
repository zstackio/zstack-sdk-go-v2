// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// AddSanSecSecurityMachineParamDetail AddSanSecSecurityMachine detail param
type AddSanSecSecurityMachineParamDetail struct {
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
	AddSanSecSecurityMachine AddSanSecSecurityMachineParamDetail `json:"addSanSecSecurityMachine"`
}
// UpdateSanSecSecurityMachineParamDetail UpdateSanSecSecurityMachine detail param
type UpdateSanSecSecurityMachineParamDetail struct {
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
	UpdateSanSecSecurityMachine UpdateSanSecSecurityMachineParamDetail `json:"updateSanSecSecurityMachine"`
}
