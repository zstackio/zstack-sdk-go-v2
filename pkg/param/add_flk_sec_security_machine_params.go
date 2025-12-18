// Copyright (c) ZStack.io, Inc.

package param

// AddFlkSecSecurityMachineDetailParam AddFlkSecSecurityMachine详细参数
type AddFlkSecSecurityMachineDetailParam struct {
	rest int `json:"port" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"managementIp" validate:"required"` // 必填
	rest string `json:"model" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"secretResourcePoolUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddFlkSecSecurityMachineParam AddFlkSecSecurityMachine请求参数
type AddFlkSecSecurityMachineParam struct {
	BaseParam
	Params AddFlkSecSecurityMachineDetailParam `json:"params"` // 详细参数
}

