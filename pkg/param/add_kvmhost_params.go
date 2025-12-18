// Copyright (c) ZStack.io, Inc.

package param

// AddKVMHostDetailParam AddKVMHost详细参数
type AddKVMHostDetailParam struct {
	rest string `json:"username" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest int `json:"sshPort,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"managementIp" validate:"required"` // 必填
	rest string `json:"clusterUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddKVMHostParam AddKVMHost请求参数
type AddKVMHostParam struct {
	BaseParam
	Params AddKVMHostDetailParam `json:"params"` // 详细参数
}

