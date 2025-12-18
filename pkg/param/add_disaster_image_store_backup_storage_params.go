// Copyright (c) ZStack.io, Inc.

package param

// AddDisasterImageStoreBackupStorageDetailParam AddDisasterImageStoreBackupStorage detail param
type AddDisasterImageStoreBackupStorageDetailParam struct {
	AttachPoint string `json:"attachPoint,omitempty"`
	EndPoint string `json:"endPoint,omitempty"`
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

// AddDisasterImageStoreBackupStorageParam AddDisasterImageStoreBackupStorage request param
type AddDisasterImageStoreBackupStorageParam struct {
	BaseParam
	Params AddDisasterImageStoreBackupStorageDetailParam `json:"params"`
}
