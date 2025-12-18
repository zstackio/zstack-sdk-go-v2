// Copyright (c) ZStack.io, Inc.

package view

import "time"

// DatabaseBackupStorageRefInventoryView DatabaseBackupStorageRef
type DatabaseBackupStorageRefInventoryView struct {
	rest string `json:"databaseBackupUuid,omitempty"`
	rest string `json:"backupStorageUuid,omitempty"`
	rest string `json:"installPath,omitempty"`
	rest string `json:"exportUrl,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

