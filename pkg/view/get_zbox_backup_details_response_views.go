// Copyright (c) ZStack.io, Inc.

package view

// GetZBoxBackupDetailsView GetZBoxBackupDetails
type GetZBoxBackupDetailsView struct {
	VmBackupInfos []VmExternalBackupInfoView `json:"vmBackupInfos,omitempty"`
	VolumeBackupInfos []VolumeExternalBackupInfoView `json:"volumeBackupInfos,omitempty"`
	BackupStorageBackupInfos []BackupStorageExternalBackupInfoView `json:"backupStorageBackupInfos,omitempty"`
	Version string `json:"version,omitempty"`
	Success bool `json:"success,omitempty"`
}

