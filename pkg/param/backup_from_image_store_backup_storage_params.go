// Copyright (c) ZStack.io, Inc.

package param

// RecoverBackupFromImageStoreBackupStorageDetailParam RecoverBackupFromImageStoreBackupStorage详细参数
type RecoverBackupFromImageStoreBackupStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"srcBackupStorageUuid" validate:"required"` // 必填
	rest string `json:"dstBackupStorageUuid" validate:"required"` // 必填
}

// RecoverBackupFromImageStoreBackupStorageParam RecoverBackupFromImageStoreBackupStorage请求参数
type RecoverBackupFromImageStoreBackupStorageParam struct {
	BaseParam
	Params RecoverBackupFromImageStoreBackupStorageDetailParam `json:"params"` // 详细参数
}

