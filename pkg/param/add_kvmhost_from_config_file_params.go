// Copyright (c) ZStack.io, Inc.

package param

// AddKVMHostFromConfigFileDetailParam AddKVMHostFromConfigFile详细参数
type AddKVMHostFromConfigFileDetailParam struct {
	rest string `json:"hostInfo" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddKVMHostFromConfigFileParam AddKVMHostFromConfigFile请求参数
type AddKVMHostFromConfigFileParam struct {
	BaseParam
	Params AddKVMHostFromConfigFileDetailParam `json:"params"` // 详细参数
}

