// Copyright (c) ZStack.io, Inc.

package param

// GetImagesFromImageStoreBackupStorageDetailParam GetImagesFromImageStoreBackupStorage详细参数
type GetImagesFromImageStoreBackupStorageDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetImagesFromImageStoreBackupStorageParam GetImagesFromImageStoreBackupStorage请求参数
type GetImagesFromImageStoreBackupStorageParam struct {
	BaseParam
	Params GetImagesFromImageStoreBackupStorageDetailParam `json:"params"` // 详细参数
}

