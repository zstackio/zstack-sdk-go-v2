// Copyright (c) ZStack.io, Inc.

package param

// SyncImageFromImageStoreBackupStorageDetailParam SyncImageFromImageStoreBackupStorage详细参数
type SyncImageFromImageStoreBackupStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"srcBackupStorageUuid" validate:"required"` // 必填
	rest string `json:"dstBackupStorageUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
}

// SyncImageFromImageStoreBackupStorageParam SyncImageFromImageStoreBackupStorage请求参数
type SyncImageFromImageStoreBackupStorageParam struct {
	BaseParam
	Params SyncImageFromImageStoreBackupStorageDetailParam `json:"params"` // 详细参数
}

