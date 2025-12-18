// Copyright (c) ZStack.io, Inc.

package param

// SyncDatabaseBackupFromImageStoreBackupStorageDetailParam SyncDatabaseBackupFromImageStoreBackupStorage detail param
type SyncDatabaseBackupFromImageStoreBackupStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" validate:"required"`
	DstBackupStorageUuid string `json:"dstBackupStorageUuid" validate:"required"`
}

// SyncDatabaseBackupFromImageStoreBackupStorageParam SyncDatabaseBackupFromImageStoreBackupStorage request param
type SyncDatabaseBackupFromImageStoreBackupStorageParam struct {
	BaseParam
	Params SyncDatabaseBackupFromImageStoreBackupStorageDetailParam `json:"params"`
}
