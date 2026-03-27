// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateJitSecurityMachineParamDetail UpdateJitSecurityMachine detail param
type UpdateJitSecurityMachineParamDetail struct {
	Port int `json:"port" validate:"required"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ManagementIp *string `json:"managementIp,omitempty"`
	Model *string `json:"model,omitempty"`
}

// UpdateJitSecurityMachineParam UpdateJitSecurityMachine request param
type UpdateJitSecurityMachineParam struct {
	BaseParam
	Params UpdateJitSecurityMachineParamDetail `json:"updateJitSecurityMachine"`
}
// AddJitSecurityMachineParamDetail AddJitSecurityMachine detail param
type AddJitSecurityMachineParamDetail struct {
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

// AddJitSecurityMachineParam AddJitSecurityMachine request param
type AddJitSecurityMachineParam struct {
	BaseParam
	Params AddJitSecurityMachineParamDetail `json:"params"`
}
