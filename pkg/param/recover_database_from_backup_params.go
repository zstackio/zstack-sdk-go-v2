// Copyright (c) ZStack.io, Inc.

package param

// RecoverDatabaseFromBackupDetailParam RecoverDatabaseFromBackup detail param
type RecoverDatabaseFromBackupDetailParam struct {
	Uuid string `json:"uuid,omitempty"`
	BackupStorageUrl string `json:"backupStorageUrl,omitempty"`
	BackupInstallPath string `json:"backupInstallPath,omitempty"`
	MysqlRootPassword string `json:"mysqlRootPassword" validate:"required"`
}

// RecoverDatabaseFromBackupParam RecoverDatabaseFromBackup request param
type RecoverDatabaseFromBackupParam struct {
	BaseParam
	Params RecoverDatabaseFromBackupDetailParam `json:"params"`
}
