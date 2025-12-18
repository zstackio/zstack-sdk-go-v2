// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VolumeSnapshotGroupInventoryView VolumeSnapshotGroup
type VolumeSnapshotGroupInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest int `json:"snapshotCount,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []VolumeSnapshotGroupRefInventoryView `json:"volumeSnapshotRefs,omitempty"`
}

