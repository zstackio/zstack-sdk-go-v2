// Copyright (c) ZStack.io, Inc.

package param

// ExportDatabaseBackupFromBackupStorageDetailParam ExportDatabaseBackupFromBackupStorage详细参数
type ExportDatabaseBackupFromBackupStorageDetailParam struct {
	rest string `json:"backupStorageUuid" validate:"required"` // 必填
	rest string `json:"databaseBackupUuid" validate:"required"` // 必填
}

// ExportDatabaseBackupFromBackupStorageParam ExportDatabaseBackupFromBackupStorage请求参数
type ExportDatabaseBackupFromBackupStorageParam struct {
	BaseParam
	Params ExportDatabaseBackupFromBackupStorageDetailParam `json:"params"` // 详细参数
}

