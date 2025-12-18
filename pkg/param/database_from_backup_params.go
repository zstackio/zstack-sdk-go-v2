// Copyright (c) ZStack.io, Inc.

package param

// RecoverDatabaseFromBackupDetailParam RecoverDatabaseFromBackup详细参数
type RecoverDatabaseFromBackupDetailParam struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"backupStorageUrl,omitempty"`
	rest string `json:"backupInstallPath,omitempty"`
	rest string `json:"mysqlRootPassword" validate:"required"` // 必填
}

// RecoverDatabaseFromBackupParam RecoverDatabaseFromBackup请求参数
type RecoverDatabaseFromBackupParam struct {
	BaseParam
	Params RecoverDatabaseFromBackupDetailParam `json:"params"` // 详细参数
}

