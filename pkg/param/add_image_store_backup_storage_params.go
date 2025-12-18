// Copyright (c) ZStack.io, Inc.

package param

// AddImageStoreBackupStorageDetailParam AddImageStoreBackupStorage detail param
type AddImageStoreBackupStorageDetailParam struct {
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

// AddImageStoreBackupStorageParam AddImageStoreBackupStorage request param
type AddImageStoreBackupStorageParam struct {
	BaseParam
	Params AddImageStoreBackupStorageDetailParam `json:"params"`
}
