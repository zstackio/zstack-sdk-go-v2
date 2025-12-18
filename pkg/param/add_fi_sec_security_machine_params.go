// Copyright (c) ZStack.io, Inc.

package param

// AddFiSecSecurityMachineDetailParam AddFiSecSecurityMachine详细参数
type AddFiSecSecurityMachineDetailParam struct {
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

// AddFiSecSecurityMachineParam AddFiSecSecurityMachine请求参数
type AddFiSecSecurityMachineParam struct {
	BaseParam
	Params AddFiSecSecurityMachineDetailParam `json:"params"` // 详细参数
}

