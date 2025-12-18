// Copyright (c) ZStack.io, Inc.

package param

// BackupStorageMigrateImageDetailParam BackupStorageMigrateImage detail param
type BackupStorageMigrateImageDetailParam struct {
	ImageUuid string `json:"imageUuid" validate:"required"`
	SrcBackupStorageUuid string `json:"srcBackupStorageUuid" validate:"required"`
	DstBackupStorageUuid string `json:"dstBackupStorageUuid" validate:"required"`
}

// BackupStorageMigrateImageParam BackupStorageMigrateImage request param
type BackupStorageMigrateImageParam struct {
	BaseParam
	Params BackupStorageMigrateImageDetailParam `json:"params"`
}
