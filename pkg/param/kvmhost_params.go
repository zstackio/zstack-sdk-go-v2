// Copyright (c) ZStack.io, Inc.

package param

// UpdateKVMHostDetailParam UpdateKVMHost详细参数
type UpdateKVMHostDetailParam struct {
	rest string `json:"username,omitempty"`
	rest string `json:"password,omitempty"`
	rest int `json:"sshPort,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"managementIp,omitempty"`
}

// UpdateKVMHostParam UpdateKVMHost请求参数
type UpdateKVMHostParam struct {
	BaseParam
	Params UpdateKVMHostDetailParam `json:"params"` // 详细参数
}

