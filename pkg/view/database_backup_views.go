// Copyright (c) ZStack.io, Inc.

package view

import "time"

// DatabaseBackupInventoryView DatabaseBackup
type DatabaseBackupInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest string `json:"metadata,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []DatabaseBackupStorageRefInventoryView `json:"backupStorageRefs,omitempty"`
}

