// Copyright (c) ZStack.io, Inc.

package param

// ListVMsFromKVMHostDetailParam ListVMsFromKVMHost detail param
type ListVMsFromKVMHostDetailParam struct {
	LibvirtURI string `json:"libvirtURI" validate:"required"`
	ConversionHostUuid string `json:"conversionHostUuid" validate:"required"`
	SshPrivKey string `json:"sshPrivKey,omitempty"`
	V2vType string `json:"v2vType,omitempty"`
}

// ListVMsFromKVMHostParam ListVMsFromKVMHost request param
type ListVMsFromKVMHostParam struct {
	BaseParam
	Params ListVMsFromKVMHostDetailParam `json:"params"`
}
