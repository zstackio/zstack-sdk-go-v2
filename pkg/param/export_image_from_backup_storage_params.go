// Copyright (c) ZStack.io, Inc.

package param

// ExportImageFromBackupStorageDetailParam ExportImageFromBackupStorage detail param
type ExportImageFromBackupStorageDetailParam struct {
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	ImageUuid string `json:"imageUuid" validate:"required"`
	ExportFormat string `json:"exportFormat,omitempty"`
}

// ExportImageFromBackupStorageParam ExportImageFromBackupStorage request param
type ExportImageFromBackupStorageParam struct {
	BaseParam
	Params ExportImageFromBackupStorageDetailParam `json:"params"`
}
