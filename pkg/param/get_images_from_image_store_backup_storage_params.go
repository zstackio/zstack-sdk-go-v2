// Copyright (c) ZStack.io, Inc.

package param

// GetImagesFromImageStoreBackupStorageDetailParam GetImagesFromImageStoreBackupStorage detail param
type GetImagesFromImageStoreBackupStorageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetImagesFromImageStoreBackupStorageParam GetImagesFromImageStoreBackupStorage request param
type GetImagesFromImageStoreBackupStorageParam struct {
	BaseParam
	Params GetImagesFromImageStoreBackupStorageDetailParam `json:"params"`
}
