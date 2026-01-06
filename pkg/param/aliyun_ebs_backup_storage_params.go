// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// AddAliyunEbsBackupStorageParamDetail AddAliyunEbsBackupStorage detail param
type AddAliyunEbsBackupStorageParamDetail struct {
	OssBucketUuid string `json:"ossBucketUuid" validate:"required"`
	Url string `json:"url,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	ImportImages bool `json:"importImages,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAliyunEbsBackupStorageParam AddAliyunEbsBackupStorage request param
type AddAliyunEbsBackupStorageParam struct {
	BaseParam
	Params AddAliyunEbsBackupStorageParamDetail `json:"params"`
}
// UpdateAliyunEbsBackupStorageParamDetail UpdateAliyunEbsBackupStorage detail param
type UpdateAliyunEbsBackupStorageParamDetail struct {
	OssBucketUuid string `json:"ossBucketUuid,omitempty"`
	Url string `json:"url,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateAliyunEbsBackupStorageParam UpdateAliyunEbsBackupStorage request param
type UpdateAliyunEbsBackupStorageParam struct {
	BaseParam
	Params UpdateAliyunEbsBackupStorageParamDetail `json:"params"`
}
