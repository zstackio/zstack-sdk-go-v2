// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VolumeSnapshotTreeInventoryView VolumeSnapshotTree
type VolumeSnapshotTreeInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"volumeUuid,omitempty"`
	rest bool `json:"current,omitempty"`
	rest string `json:"status,omitempty"`
	rest SnapshotLeafInventoryView `json:"tree,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

