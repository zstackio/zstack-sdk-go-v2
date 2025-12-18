// Copyright (c) ZStack.io, Inc.

package param

// AddInfoSecSecurityMachineDetailParam AddInfoSecSecurityMachine详细参数
type AddInfoSecSecurityMachineDetailParam struct {
	rest string `json:"password" validate:"required"` // 必填
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

// AddInfoSecSecurityMachineParam AddInfoSecSecurityMachine请求参数
type AddInfoSecSecurityMachineParam struct {
	BaseParam
	Params AddInfoSecSecurityMachineDetailParam `json:"params"` // 详细参数
}

