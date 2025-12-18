// Copyright (c) ZStack.io, Inc.

package param

// SyncVmBackupFromImageStoreBackupStorageDetailParam SyncVmBackupFromImageStoreBackupStorage详细参数
type SyncVmBackupFromImageStoreBackupStorageDetailParam struct {
	rest string `json:"groupUuid" validate:"required"` // 必填
	rest string `json:"srcBackupStorageUuid" validate:"required"` // 必填
	rest string `json:"dstBackupStorageUuid" validate:"required"` // 必填
}

// SyncVmBackupFromImageStoreBackupStorageParam SyncVmBackupFromImageStoreBackupStorage请求参数
type SyncVmBackupFromImageStoreBackupStorageParam struct {
	BaseParam
	Params SyncVmBackupFromImageStoreBackupStorageDetailParam `json:"params"` // 详细参数
}

