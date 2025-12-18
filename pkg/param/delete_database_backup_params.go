// Copyright (c) ZStack.io, Inc.

package param

// DeleteDatabaseBackupDetailParam DeleteDatabaseBackup detail param
type DeleteDatabaseBackupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
}

// DeleteDatabaseBackupParam DeleteDatabaseBackup request param
type DeleteDatabaseBackupParam struct {
	BaseParam
	Params DeleteDatabaseBackupDetailParam `json:"params"`
}
