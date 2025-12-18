// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// DatabaseBackupInventoryView DatabaseBackup
type DatabaseBackupInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	Size int64 `json:"size,omitempty"`
	Metadata string `json:"metadata,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	BackupStorageRefs []DatabaseBackupStorageRefInventoryView `json:"backupStorageRefs,omitempty"`
}

