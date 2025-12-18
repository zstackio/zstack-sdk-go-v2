// Copyright (c) ZStack.io, Inc.

package param

// AddAliyunEbsBackupStorageDetailParam AddAliyunEbsBackupStorage detail param
type AddAliyunEbsBackupStorageDetailParam struct {
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
	Params AddAliyunEbsBackupStorageDetailParam `json:"params"`
}
