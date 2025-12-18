// Copyright (c) ZStack.io, Inc.

package param

// RecoveryImageFromImageStoreBackupStorageDetailParam RecoveryImageFromImageStoreBackupStorage详细参数
type RecoveryImageFromImageStoreBackupStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"srcBackupStorageUuid" validate:"required"` // 必填
	rest string `json:"dstBackupStorageUuid" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
}

// RecoveryImageFromImageStoreBackupStorageParam RecoveryImageFromImageStoreBackupStorage请求参数
type RecoveryImageFromImageStoreBackupStorageParam struct {
	BaseParam
	Params RecoveryImageFromImageStoreBackupStorageDetailParam `json:"params"` // 详细参数
}

