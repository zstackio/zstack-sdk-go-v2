// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// DatabaseBackupStorageRefInventoryView DatabaseBackupStorageRef
type DatabaseBackupStorageRefInventoryView struct {
	DatabaseBackupUuid *string `json:"databaseBackupUuid,omitempty"`
	BackupStorageUuid *string `json:"backupStorageUuid,omitempty"`
	InstallPath *string `json:"installPath,omitempty"`
	ExportUrl *string `json:"exportUrl,omitempty"`
	Status *string `json:"status,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

