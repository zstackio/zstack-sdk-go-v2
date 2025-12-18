// Copyright (c) ZStack.io, Inc.

package view

// GetZBoxBackupDetailsView GetZBoxBackupDetails
type GetZBoxBackupDetailsView struct {
	VmBackupInfos []interface{} `json:"vmBackupInfos,omitempty"`
	VolumeBackupInfos []interface{} `json:"volumeBackupInfos,omitempty"`
	BackupStorageBackupInfos []interface{} `json:"backupStorageBackupInfos,omitempty"`
	Version string `json:"version,omitempty"`
	Success bool `json:"success,omitempty"`
}

