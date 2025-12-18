// Copyright (c) ZStack.io, Inc.

package param

// AddSftpBackupStorageDetailParam AddSftpBackupStorage detail param
type AddSftpBackupStorageDetailParam struct {
	Hostname string `json:"hostname" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	SshPort int `json:"sshPort,omitempty"`
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	ImportImages bool `json:"importImages,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSftpBackupStorageParam AddSftpBackupStorage request param
type AddSftpBackupStorageParam struct {
	BaseParam
	Params AddSftpBackupStorageDetailParam `json:"params"`
}
