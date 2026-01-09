// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VolumeSnapshotGroupRefInventoryView VolumeSnapshotGroupRef
type VolumeSnapshotGroupRefInventoryView struct {
	VolumeSnapshotUuid *string `json:"volumeSnapshotUuid,omitempty"`
	VolumeSnapshotGroupUuid *string `json:"volumeSnapshotGroupUuid,omitempty"`
	DeviceId int `json:"deviceId,omitempty"`
	SnapshotDeleted bool `json:"snapshotDeleted,omitempty"`
	VolumeUuid *string `json:"volumeUuid,omitempty"`
	VolumeName *string `json:"volumeName,omitempty"`
	VolumeType *string `json:"volumeType,omitempty"`
	VolumeSnapshotInstallPath *string `json:"volumeSnapshotInstallPath,omitempty"`
	VolumeSnapshotName *string `json:"volumeSnapshotName,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	VolumeLastAttachDate *time.Time `json:"volumeLastAttachDate,omitempty"`
}

