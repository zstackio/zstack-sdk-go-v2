// Copyright (c) ZStack.io, Inc.

package param

// ListVMsFromKVMHostDetailParam ListVMsFromKVMHost详细参数
type ListVMsFromKVMHostDetailParam struct {
	rest string `json:"libvirtURI" validate:"required"` // 必填
	rest string `json:"conversionHostUuid" validate:"required"` // 必填
	rest string `json:"sshPrivKey,omitempty"`
	rest string `json:"v2vType,omitempty"`
}

// ListVMsFromKVMHostParam ListVMsFromKVMHost请求参数
type ListVMsFromKVMHostParam struct {
	BaseParam
	Params ListVMsFromKVMHostDetailParam `json:"params"` // 详细参数
}

