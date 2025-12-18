// Copyright (c) ZStack.io, Inc.

package param

// SyncDatabaseBackupFromImageStoreBackupStorageDetailParam SyncDatabaseBackupFromImageStoreBackupStorage详细参数
type SyncDatabaseBackupFromImageStoreBackupStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"srcBackupStorageUuid" validate:"required"` // 必填
	rest string `json:"dstBackupStorageUuid" validate:"required"` // 必填
}

// SyncDatabaseBackupFromImageStoreBackupStorageParam SyncDatabaseBackupFromImageStoreBackupStorage请求参数
type SyncDatabaseBackupFromImageStoreBackupStorageParam struct {
	BaseParam
	Params SyncDatabaseBackupFromImageStoreBackupStorageDetailParam `json:"params"` // 详细参数
}

