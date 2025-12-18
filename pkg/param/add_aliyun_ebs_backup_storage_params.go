// Copyright (c) ZStack.io, Inc.

package param

// AddAliyunEbsBackupStorageDetailParam AddAliyunEbsBackupStorage详细参数
type AddAliyunEbsBackupStorageDetailParam struct {
	rest string `json:"ossBucketUuid" validate:"required"` // 必填
	rest string `json:"url,omitempty"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest bool `json:"importImages,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// AddAliyunEbsBackupStorageParam AddAliyunEbsBackupStorage请求参数
type AddAliyunEbsBackupStorageParam struct {
	BaseParam
	Params AddAliyunEbsBackupStorageDetailParam `json:"params"` // 详细参数
}

