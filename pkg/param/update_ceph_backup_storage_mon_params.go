// Copyright (c) ZStack.io, Inc.

package param

// UpdateCephBackupStorageMonDetailParam UpdateCephBackupStorageMon detail param
type UpdateCephBackupStorageMonDetailParam struct {
	MonUuid string `json:"monUuid" validate:"required"`
	Hostname string `json:"hostname,omitempty"`
	SshUsername string `json:"sshUsername,omitempty"`
	SshPassword string `json:"sshPassword,omitempty"`
	SshPort int `json:"sshPort,omitempty"`
	MonPort int `json:"monPort,omitempty"`
}

// UpdateCephBackupStorageMonParam UpdateCephBackupStorageMon request param
type UpdateCephBackupStorageMonParam struct {
	BaseParam
	Params UpdateCephBackupStorageMonDetailParam `json:"params"`
}
