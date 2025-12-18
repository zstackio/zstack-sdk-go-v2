// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunEbsBackupStorageDetailParam UpdateAliyunEbsBackupStorage detail param
type UpdateAliyunEbsBackupStorageDetailParam struct {
	OssBucketUuid string `json:"ossBucketUuid,omitempty"`
	Url string `json:"url,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateAliyunEbsBackupStorageParam UpdateAliyunEbsBackupStorage request param
type UpdateAliyunEbsBackupStorageParam struct {
	BaseParam
	Params UpdateAliyunEbsBackupStorageDetailParam `json:"params"`
}
