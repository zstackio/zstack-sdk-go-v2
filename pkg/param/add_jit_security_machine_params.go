// Copyright (c) ZStack.io, Inc.

package param

// AddJitSecurityMachineDetailParam AddJitSecurityMachine详细参数
type AddJitSecurityMachineDetailParam struct {
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

// AddJitSecurityMachineParam AddJitSecurityMachine请求参数
type AddJitSecurityMachineParam struct {
	BaseParam
	Params AddJitSecurityMachineDetailParam `json:"params"` // 详细参数
}

