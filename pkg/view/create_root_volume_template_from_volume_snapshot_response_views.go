// Copyright (c) ZStack.io, Inc.

package view

// CreateRootVolumeTemplateFromVolumeSnapshotEventView CreateRootVolumeTemplateFromVolumeSnapshotEvent
type CreateRootVolumeTemplateFromVolumeSnapshotEventView struct {
	Inventory ImageInventoryView `json:"inventory,omitempty"`
	Failures []FailureView `json:"failures,omitempty"`
}

