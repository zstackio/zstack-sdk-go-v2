// Copyright (c) ZStack.io, Inc.

package view

// BatchDeleteVolumeSnapshotEventView BatchDeleteVolumeSnapshotEvent
type BatchDeleteVolumeSnapshotEventView struct {
	Results []BatchDeleteVolumeSnapshotStructView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

