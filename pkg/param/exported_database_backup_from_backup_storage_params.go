// Copyright (c) ZStack.io, Inc.

package param

// DeleteExportedDatabaseBackupFromBackupStorageDetailParam DeleteExportedDatabaseBackupFromBackupStorage详细参数
type DeleteExportedDatabaseBackupFromBackupStorageDetailParam struct {
	rest string `json:"backupStorageUuid" validate:"required"` // 必填
	rest string `json:"databaseBackupUuid" validate:"required"` // 必填
}

// DeleteExportedDatabaseBackupFromBackupStorageParam DeleteExportedDatabaseBackupFromBackupStorage请求参数
type DeleteExportedDatabaseBackupFromBackupStorageParam struct {
	BaseParam
	Params DeleteExportedDatabaseBackupFromBackupStorageDetailParam `json:"params"` // 详细参数
}

