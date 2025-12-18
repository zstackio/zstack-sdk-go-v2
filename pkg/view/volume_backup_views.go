// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VolumeBackupInventoryView VolumeBackup
type VolumeBackupInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"volumeUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest string `json:"metadata,omitempty"`
	rest string `json:"groupUuid,omitempty"`
	rest string `json:"mode,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []VolumeBackupStorageRefInventoryView `json:"backupStorageRefs,omitempty"`
}

