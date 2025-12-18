// Copyright (c) ZStack.io, Inc.

package param

// UpdateImageStoreBackupStorageDetailParam UpdateImageStoreBackupStorage detail param
type UpdateImageStoreBackupStorageDetailParam struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	SshPort int `json:"sshPort,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateImageStoreBackupStorageParam UpdateImageStoreBackupStorage request param
type UpdateImageStoreBackupStorageParam struct {
	BaseParam
	Params UpdateImageStoreBackupStorageDetailParam `json:"params"`
}
