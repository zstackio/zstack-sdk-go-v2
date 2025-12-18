// Copyright (c) ZStack.io, Inc.

package param

// ExportDatabaseBackupFromBackupStorageDetailParam ExportDatabaseBackupFromBackupStorage detail param
type ExportDatabaseBackupFromBackupStorageDetailParam struct {
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	DatabaseBackupUuid string `json:"databaseBackupUuid" validate:"required"`
}

// ExportDatabaseBackupFromBackupStorageParam ExportDatabaseBackupFromBackupStorage request param
type ExportDatabaseBackupFromBackupStorageParam struct {
	BaseParam
	Params ExportDatabaseBackupFromBackupStorageDetailParam `json:"params"`
}
