// Copyright (c) ZStack.io, Inc.

package param

// UpdateSftpBackupStorageDetailParam UpdateSftpBackupStorage detail param
type UpdateSftpBackupStorageDetailParam struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	SshPort int `json:"sshPort,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateSftpBackupStorageParam UpdateSftpBackupStorage request param
type UpdateSftpBackupStorageParam struct {
	BaseParam
	Params UpdateSftpBackupStorageDetailParam `json:"params"`
}
