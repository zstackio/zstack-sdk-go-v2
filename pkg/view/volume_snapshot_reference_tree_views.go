// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VolumeSnapshotReferenceTreeInventoryView VolumeSnapshotReferenceTree
type VolumeSnapshotReferenceTreeInventoryView struct {
	rest string `json:"primaryStorageUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"rootImageUuid,omitempty"`
	rest string `json:"rootVolumeUuid,omitempty"`
	rest string `json:"rootVolumeSnapshotUuid,omitempty"`
	rest string `json:"rootVolumeSnapshotTreeUuid,omitempty"`
	rest string `json:"rootInstallUrl,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"resourceName,omitempty"`
	rest string `json:"resourceType,omitempty"`
	rest string `json:"concreteResourceType,omitempty"`
}

