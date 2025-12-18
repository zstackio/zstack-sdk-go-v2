// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VolumeSnapshotGroupRefInventoryView VolumeSnapshotGroupRef
type VolumeSnapshotGroupRefInventoryView struct {
	rest string `json:"volumeSnapshotUuid,omitempty"`
	rest string `json:"volumeSnapshotGroupUuid,omitempty"`
	rest int `json:"deviceId,omitempty"`
	rest bool `json:"snapshotDeleted,omitempty"`
	rest string `json:"volumeUuid,omitempty"`
	rest string `json:"volumeName,omitempty"`
	rest string `json:"volumeType,omitempty"`
	rest string `json:"volumeSnapshotInstallPath,omitempty"`
	rest string `json:"volumeSnapshotName,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest time.Time `json:"volumeLastAttachDate,omitempty"`
}

