// Copyright (c) ZStack.io, Inc.

package view

// GetMemorySnapshotGroupReferenceView GetMemorySnapshotGroupReference
type GetMemorySnapshotGroupReferenceView struct {
	Inventories []VolumeSnapshotGroupInventoryView `json:"inventories,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	Success bool `json:"success,omitempty"`
}

