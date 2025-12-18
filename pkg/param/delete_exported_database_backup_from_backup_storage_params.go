// Copyright (c) ZStack.io, Inc.

package param

// DeleteExportedDatabaseBackupFromBackupStorageDetailParam DeleteExportedDatabaseBackupFromBackupStorage detail param
type DeleteExportedDatabaseBackupFromBackupStorageDetailParam struct {
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	DatabaseBackupUuid string `json:"databaseBackupUuid" validate:"required"`
}

// DeleteExportedDatabaseBackupFromBackupStorageParam DeleteExportedDatabaseBackupFromBackupStorage request param
type DeleteExportedDatabaseBackupFromBackupStorageParam struct {
	BaseParam
	Params DeleteExportedDatabaseBackupFromBackupStorageDetailParam `json:"params"`
}
