// Copyright (c) ZStack.io, Inc.

package param

// CheckKVMHostConfigFileDetailParam CheckKVMHostConfigFile detail param
type CheckKVMHostConfigFileDetailParam struct {
	HostInfo string `json:"hostInfo" validate:"required"`
}

// CheckKVMHostConfigFileParam CheckKVMHostConfigFile request param
type CheckKVMHostConfigFileParam struct {
	BaseParam
	Params CheckKVMHostConfigFileDetailParam `json:"params"`
}
