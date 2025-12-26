// Copyright (c) ZStack.io, Inc.

package view

// GetDatabaseBackupFromImageStoreView GetDatabaseBackupFromImageStore
type GetDatabaseBackupFromImageStoreView struct {
	Backups []DatabaseBackupStructView `json:"backups,omitempty"`
	Success bool `json:"success,omitempty"`
}

