// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteDatabaseBackupParamDetail DeleteDatabaseBackup detail param
type DeleteDatabaseBackupParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	BackupStorageUuids []string `json:"backupStorageUuids,omitempty"`
}

// DeleteDatabaseBackupParam DeleteDatabaseBackup request param
type DeleteDatabaseBackupParam struct {
	BaseParam
	Params DeleteDatabaseBackupParamDetail `json:"deleteDatabaseBackup"`
}
// CreateDatabaseBackupParamDetail CreateDatabaseBackup detail param
type CreateDatabaseBackupParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	BackupStorageUuid string `json:"backupStorageUuid" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateDatabaseBackupParam CreateDatabaseBackup request param
type CreateDatabaseBackupParam struct {
	BaseParam
	Params CreateDatabaseBackupParamDetail `json:"createDatabaseBackup"`
}
// SyncDatabaseBackupParamDetail SyncDatabaseBackup detail param
type SyncDatabaseBackupParamDetail struct {
	ImageStoreUuid string `json:"imageStoreUuid" validate:"required"`
}

// SyncDatabaseBackupParam SyncDatabaseBackup request param
type SyncDatabaseBackupParam struct {
	BaseParam
	Params SyncDatabaseBackupParamDetail `json:"syncDatabaseBackup"`
}
