// Copyright (c) ZStack.io, Inc.

package view

// CreateDataVolumeTemplateFromVolumeSnapshotEventView CreateDataVolumeTemplateFromVolumeSnapshotEvent
type CreateDataVolumeTemplateFromVolumeSnapshotEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
	Failures []FailureView `json:"failures,omitempty"`
}

