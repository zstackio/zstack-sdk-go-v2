// Copyright (c) ZStack.io, Inc.

package param

// CheckKVMHostConfigFileDetailParam CheckKVMHostConfigFile详细参数
type CheckKVMHostConfigFileDetailParam struct {
	rest string `json:"hostInfo" validate:"required"` // 必填
}

// CheckKVMHostConfigFileParam CheckKVMHostConfigFile请求参数
type CheckKVMHostConfigFileParam struct {
	BaseParam
	Params CheckKVMHostConfigFileDetailParam `json:"params"` // 详细参数
}

