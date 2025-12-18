// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VolumeSnapshotGroupInventoryView VolumeSnapshotGroup
type VolumeSnapshotGroupInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	SnapshotCount int `json:"snapshotCount,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	VolumeSnapshotRefs []VolumeSnapshotGroupRefInventoryView `json:"volumeSnapshotRefs,omitempty"`
}

