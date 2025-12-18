// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VolumeBackupStorageRefInventoryView VolumeBackupStorageRef
type VolumeBackupStorageRefInventoryView struct {
	rest string `json:"volumeBackupUuid,omitempty"`
	rest string `json:"backupStorageUuid,omitempty"`
	rest string `json:"installPath,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

