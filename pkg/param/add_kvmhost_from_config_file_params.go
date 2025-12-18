// Copyright (c) ZStack.io, Inc.

package param

// AddKVMHostFromConfigFileDetailParam AddKVMHostFromConfigFile detail param
type AddKVMHostFromConfigFileDetailParam struct {
	HostInfo string `json:"hostInfo" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddKVMHostFromConfigFileParam AddKVMHostFromConfigFile request param
type AddKVMHostFromConfigFileParam struct {
	BaseParam
	Params AddKVMHostFromConfigFileDetailParam `json:"params"`
}
