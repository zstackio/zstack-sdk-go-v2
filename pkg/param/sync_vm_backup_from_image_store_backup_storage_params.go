// Copyright (c) ZStack.io, Inc.

package param

// SyncVmBackupFromImageStoreBackupStorageDetailParam SyncVmBackupFromImageStoreBackupStorage detail param
type SyncVmBackupFromImageStoreBackupStorageDetailParam struct {
	GroupUuid string `json:"groupUuid" validate:"required"`
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" validate:"required"`
	DstBackupStorageUuid string `json:"dstBackupStorageUuid" validate:"required"`
}

// SyncVmBackupFromImageStoreBackupStorageParam SyncVmBackupFromImageStoreBackupStorage request param
type SyncVmBackupFromImageStoreBackupStorageParam struct {
	BaseParam
	Params SyncVmBackupFromImageStoreBackupStorageDetailParam `json:"params"`
}
