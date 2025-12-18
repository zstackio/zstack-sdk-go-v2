// Copyright (c) ZStack.io, Inc.

package param

// AddKVMHostDetailParam AddKVMHost detail param
type AddKVMHostDetailParam struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	SshPort int `json:"sshPort,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddKVMHostParam AddKVMHost request param
type AddKVMHostParam struct {
	BaseParam
	Params AddKVMHostDetailParam `json:"params"`
}
