// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VolumeSnapshotReferenceInventoryView VolumeSnapshotReference
type VolumeSnapshotReferenceInventoryView struct {
	rest int64 `json:"id,omitempty"`
	rest int64 `json:"parentId,omitempty"`
	rest string `json:"volumeUuid,omitempty"`
	rest string `json:"volumeSnapshotUuid,omitempty"`
	rest string `json:"volumeSnapshotInstallUrl,omitempty"`
	rest string `json:"directSnapshotUuid,omitempty"`
	rest string `json:"directSnapshotInstallUrl,omitempty"`
	rest string `json:"referenceUuid,omitempty"`
	rest string `json:"referenceType,omitempty"`
	rest string `json:"referenceInstallUrl,omitempty"`
	rest string `json:"referenceVolumeUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

