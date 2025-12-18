// Copyright (c) ZStack.io, Inc.

package param

// RecoveryImageFromImageStoreBackupStorageDetailParam RecoveryImageFromImageStoreBackupStorage detail param
type RecoveryImageFromImageStoreBackupStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" validate:"required"`
	DstBackupStorageUuid string `json:"dstBackupStorageUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
}

// RecoveryImageFromImageStoreBackupStorageParam RecoveryImageFromImageStoreBackupStorage request param
type RecoveryImageFromImageStoreBackupStorageParam struct {
	BaseParam
	Params RecoveryImageFromImageStoreBackupStorageDetailParam `json:"params"`
}
