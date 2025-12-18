// Copyright (c) ZStack.io, Inc.

package param

// RecoverVmBackupFromImageStoreBackupStorageDetailParam RecoverVmBackupFromImageStoreBackupStorage detail param
type RecoverVmBackupFromImageStoreBackupStorageDetailParam struct {
	GroupUuid string `json:"groupUuid" validate:"required"`
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" validate:"required"`
	DstBackupStorageUuid string `json:"dstBackupStorageUuid" validate:"required"`
}

// RecoverVmBackupFromImageStoreBackupStorageParam RecoverVmBackupFromImageStoreBackupStorage request param
type RecoverVmBackupFromImageStoreBackupStorageParam struct {
	BaseParam
	Params RecoverVmBackupFromImageStoreBackupStorageDetailParam `json:"params"`
}
