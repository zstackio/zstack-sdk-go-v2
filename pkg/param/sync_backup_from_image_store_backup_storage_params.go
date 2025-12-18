// Copyright (c) ZStack.io, Inc.

package param

// SyncBackupFromImageStoreBackupStorageDetailParam SyncBackupFromImageStoreBackupStorage detail param
type SyncBackupFromImageStoreBackupStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" validate:"required"`
	DstBackupStorageUuid string `json:"dstBackupStorageUuid" validate:"required"`
}

// SyncBackupFromImageStoreBackupStorageParam SyncBackupFromImageStoreBackupStorage request param
type SyncBackupFromImageStoreBackupStorageParam struct {
	BaseParam
	Params SyncBackupFromImageStoreBackupStorageDetailParam `json:"params"`
}
