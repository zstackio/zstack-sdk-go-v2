// Copyright (c) ZStack.io, Inc.

package param

// DeleteExportedImageFromBackupStorageDetailParam DeleteExportedImageFromBackupStorage detail param
type DeleteExportedImageFromBackupStorageDetailParam struct {
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	ImageUuid string `json:"imageUuid" validate:"required"`
}

// DeleteExportedImageFromBackupStorageParam DeleteExportedImageFromBackupStorage request param
type DeleteExportedImageFromBackupStorageParam struct {
	BaseParam
	Params DeleteExportedImageFromBackupStorageDetailParam `json:"params"`
}
