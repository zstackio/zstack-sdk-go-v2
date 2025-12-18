// Copyright (c) ZStack.io, Inc.

package param

// UpdateCephPrimaryStorageMonDetailParam UpdateCephPrimaryStorageMon detail param
type UpdateCephPrimaryStorageMonDetailParam struct {
	MonUuid string `json:"monUuid" validate:"required"`
	Hostname string `json:"hostname,omitempty"`
	SshUsername string `json:"sshUsername,omitempty"`
	SshPassword string `json:"sshPassword,omitempty"`
	SshPort int `json:"sshPort,omitempty"`
	MonPort int `json:"monPort,omitempty"`
}

// UpdateCephPrimaryStorageMonParam UpdateCephPrimaryStorageMon request param
type UpdateCephPrimaryStorageMonParam struct {
	BaseParam
	Params UpdateCephPrimaryStorageMonDetailParam `json:"params"`
}
