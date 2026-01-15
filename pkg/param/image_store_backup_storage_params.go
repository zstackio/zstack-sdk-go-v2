// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// AddImageStoreBackupStorageParamDetail AddImageStoreBackupStorage detail param
type AddImageStoreBackupStorageParamDetail struct {
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
	AddImageStoreBackupStorage AddImageStoreBackupStorageParamDetail `json:"addImageStoreBackupStorage"`
}
// UpdateImageStoreBackupStorageParamDetail UpdateImageStoreBackupStorage detail param
type UpdateImageStoreBackupStorageParamDetail struct {
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
	UpdateImageStoreBackupStorage UpdateImageStoreBackupStorageParamDetail `json:"updateImageStoreBackupStorage"`
}
// ReconnectImageStoreBackupStorageParamDetail ReconnectImageStoreBackupStorage detail param
type ReconnectImageStoreBackupStorageParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReconnectImageStoreBackupStorageParam ReconnectImageStoreBackupStorage request param
type ReconnectImageStoreBackupStorageParam struct {
	BaseParam
	ReconnectImageStoreBackupStorage ReconnectImageStoreBackupStorageParamDetail `json:"reconnectImageStoreBackupStorage"`
}
