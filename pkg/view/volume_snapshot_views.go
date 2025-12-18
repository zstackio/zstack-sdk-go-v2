// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VolumeSnapshotInventoryView VolumeSnapshot
type VolumeSnapshotInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"volumeUuid,omitempty"`
	rest string `json:"treeUuid,omitempty"`
	rest string `json:"parentUuid,omitempty"`
	rest string `json:"primaryStorageUuid,omitempty"`
	rest string `json:"primaryStorageInstallPath,omitempty"`
	rest string `json:"volumeType,omitempty"`
	rest string `json:"format,omitempty"`
	rest bool `json:"latest,omitempty"`
	rest int64 `json:"size,omitempty"`
	rest int `json:"distance,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []VolumeSnapshotBackupStorageRefInventoryView `json:"backupStorageRefs,omitempty"`
	rest string `json:"groupUuid,omitempty"`
}

