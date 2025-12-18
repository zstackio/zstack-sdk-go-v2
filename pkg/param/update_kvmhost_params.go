// Copyright (c) ZStack.io, Inc.

package param

// UpdateKVMHostDetailParam UpdateKVMHost detail param
type UpdateKVMHostDetailParam struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	SshPort int `json:"sshPort,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
}

// UpdateKVMHostParam UpdateKVMHost request param
type UpdateKVMHostParam struct {
	BaseParam
	Params UpdateKVMHostDetailParam `json:"params"`
}
