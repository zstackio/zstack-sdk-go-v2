// Copyright (c) ZStack.io, Inc.

package param

// RecoverVmBackupFromImageStoreBackupStorageDetailParam RecoverVmBackupFromImageStoreBackupStorage详细参数
type RecoverVmBackupFromImageStoreBackupStorageDetailParam struct {
	rest string `json:"groupUuid" validate:"required"` // 必填
	rest string `json:"srcBackupStorageUuid" validate:"required"` // 必填
	rest string `json:"dstBackupStorageUuid" validate:"required"` // 必填
}

// RecoverVmBackupFromImageStoreBackupStorageParam RecoverVmBackupFromImageStoreBackupStorage请求参数
type RecoverVmBackupFromImageStoreBackupStorageParam struct {
	BaseParam
	Params RecoverVmBackupFromImageStoreBackupStorageDetailParam `json:"params"` // 详细参数
}

