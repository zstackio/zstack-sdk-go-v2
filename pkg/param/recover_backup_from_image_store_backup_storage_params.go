// Copyright (c) ZStack.io, Inc.

package param

// RecoverBackupFromImageStoreBackupStorageDetailParam RecoverBackupFromImageStoreBackupStorage detail param
type RecoverBackupFromImageStoreBackupStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" validate:"required"`
	DstBackupStorageUuid string `json:"dstBackupStorageUuid" validate:"required"`
}

// RecoverBackupFromImageStoreBackupStorageParam RecoverBackupFromImageStoreBackupStorage request param
type RecoverBackupFromImageStoreBackupStorageParam struct {
	BaseParam
	Params RecoverBackupFromImageStoreBackupStorageDetailParam `json:"params"`
}
