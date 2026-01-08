// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VolumeSnapshotBackupStorageRefInventoryView VolumeSnapshotBackupStorageRef
type VolumeSnapshotBackupStorageRefInventoryView struct {
	VolumeSnapshotUuid string `json:"volumeSnapshotUuid,omitempty"`
	BackupStorageUuid  string `json:"backupStorageUuid,omitempty"`
	InstallPath        string `json:"installPath,omitempty"`
}
