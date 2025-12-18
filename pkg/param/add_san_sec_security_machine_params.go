// Copyright (c) ZStack.io, Inc.

package param

// AddSanSecSecurityMachineDetailParam AddSanSecSecurityMachine详细参数
type AddSanSecSecurityMachineDetailParam struct {
	rest int `json:"port" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
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

// AddSanSecSecurityMachineParam AddSanSecSecurityMachine请求参数
type AddSanSecSecurityMachineParam struct {
	BaseParam
	Params AddSanSecSecurityMachineDetailParam `json:"params"` // 详细参数
}

