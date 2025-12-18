// Copyright (c) ZStack.io, Inc.

package param

// SyncBackupFromImageStoreBackupStorageDetailParam SyncBackupFromImageStoreBackupStorage详细参数
type SyncBackupFromImageStoreBackupStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"srcBackupStorageUuid" validate:"required"` // 必填
	rest string `json:"dstBackupStorageUuid" validate:"required"` // 必填
}

// SyncBackupFromImageStoreBackupStorageParam SyncBackupFromImageStoreBackupStorage请求参数
type SyncBackupFromImageStoreBackupStorageParam struct {
	BaseParam
	Params SyncBackupFromImageStoreBackupStorageDetailParam `json:"params"` // 详细参数
}

