// Copyright (c) ZStack.io, Inc.

package param

// AddCephBackupStorageDetailParam AddCephBackupStorage detail param
type AddCephBackupStorageDetailParam struct {
	MonUrls []string `json:"monUrls" validate:"required"`
	PoolName string `json:"poolName,omitempty"`
	Url string `json:"url,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	ImportImages bool `json:"importImages,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddCephBackupStorageParam AddCephBackupStorage request param
type AddCephBackupStorageParam struct {
	BaseParam
	Params AddCephBackupStorageDetailParam `json:"params"`
}
