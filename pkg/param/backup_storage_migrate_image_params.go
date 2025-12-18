// Copyright (c) ZStack.io, Inc.

package param

// BackupStorageMigrateImageDetailParam BackupStorageMigrateImage详细参数
type BackupStorageMigrateImageDetailParam struct {
	rest string `json:"imageUuid" validate:"required"` // 必填
	rest string `json:"srcBackupStorageUuid" validate:"required"` // 必填
	rest string `json:"dstBackupStorageUuid" validate:"required"` // 必填
}

// BackupStorageMigrateImageParam BackupStorageMigrateImage请求参数
type BackupStorageMigrateImageParam struct {
	BaseParam
	Params BackupStorageMigrateImageDetailParam `json:"params"` // 详细参数
}

