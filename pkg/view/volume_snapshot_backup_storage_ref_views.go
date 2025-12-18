// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VolumeSnapshotBackupStorageRefInventoryView VolumeSnapshotBackupStorageRef
type VolumeSnapshotBackupStorageRefInventoryView struct {
	rest string `json:"volumeSnapshotUuid,omitempty"`
	rest string `json:"backupStorageUuid,omitempty"`
	rest string `json:"installPath,omitempty"`
}

