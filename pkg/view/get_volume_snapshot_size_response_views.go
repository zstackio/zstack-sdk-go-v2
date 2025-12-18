// Copyright (c) ZStack.io, Inc.

package view

// GetVolumeSnapshotSizeEventView GetVolumeSnapshotSizeEvent
type GetVolumeSnapshotSizeEventView struct {
	Size int64 `json:"size,omitempty"`
	ActualSize int64 `json:"actualSize,omitempty"`
	Success bool `json:"success,omitempty"`
}

