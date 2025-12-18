// Copyright (c) ZStack.io, Inc.

package param

// CreateDatabaseBackupDetailParam CreateDatabaseBackup detail param
type CreateDatabaseBackupDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDatabaseBackupParam CreateDatabaseBackup request param
type CreateDatabaseBackupParam struct {
	BaseParam
	Params CreateDatabaseBackupDetailParam `json:"params"`
}
