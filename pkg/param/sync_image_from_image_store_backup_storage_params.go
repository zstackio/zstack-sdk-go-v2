// Copyright (c) ZStack.io, Inc.

package param

// SyncImageFromImageStoreBackupStorageDetailParam SyncImageFromImageStoreBackupStorage detail param
type SyncImageFromImageStoreBackupStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" validate:"required"`
	DstBackupStorageUuid string `json:"dstBackupStorageUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
}

// SyncImageFromImageStoreBackupStorageParam SyncImageFromImageStoreBackupStorage request param
type SyncImageFromImageStoreBackupStorageParam struct {
	BaseParam
	Params SyncImageFromImageStoreBackupStorageDetailParam `json:"params"`
}
