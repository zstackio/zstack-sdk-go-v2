// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VolumeBackupStorageRefInventoryView VolumeBackupStorageRef
type VolumeBackupStorageRefInventoryView struct {
	VolumeBackupUuid  string    `json:"volumeBackupUuid,omitempty"`
	BackupStorageUuid string    `json:"backupStorageUuid,omitempty"`
	InstallPath       string    `json:"installPath,omitempty"`
	Status            string    `json:"status,omitempty"`
	CreateDate        time.Time `json:"createDate,omitempty"`
	LastOpDate        time.Time `json:"lastOpDate,omitempty"`
}
